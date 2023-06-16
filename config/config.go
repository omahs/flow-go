package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/onflow/flow-go/config/network"
)

var (
	conf = viper.New()

	//go:embed default-config.yml
	configFile string
)

// FlowConfig Flow configuration.
type FlowConfig struct {
	// ConfigFile used to set a path to a config.yml file used to override the default-config.yml file.
	ConfigFile    string          `mapstructure:"config-file"`
	NetworkConfig *network.Config `mapstructure:"network-config"`
}

// Validate checks validity of the Flow config. Errors indicate that either the configuration is broken, 
// incompatible with the node's internal state, or that the node's internal state is corrupted. In all 
// cases, continuation is impossible. 
func (fc *FlowConfig) Validate() error {
	err := fc.NetworkConfig.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate flow network configuration values: %w", err)
	}
	return nil
}

// DefaultConfig initializes the flow configuration. All default values for the Flow
// configuration are stored in the default-config.yml file. These values can be overriden
// by node operators by setting the corresponding cli flag. DefaultConfig should be called
// before any pflags are parsed, this will allow the configuration to initialize with defaults
// from default-config.yml.
// Returns:
//
//	*FlowConfig: an instance of the network configuration fully initialized to the default values set in the config file
//	error: if there is any error encountered while initializing the configuration, all errors are considered irrecoverable.
func DefaultConfig() (*FlowConfig, error) {
	var flowConfig FlowConfig
	err := Unmarshall(&flowConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall the Flow config: %w", err)
	}
	return &flowConfig, nil
}

// BindPFlags binds the configuration to the cli pflag set. This should be called
// after all pflags have been parsed. If the --config-file flag has been set the config will
// be loaded from the specified config file.
// Args:
//
//	c: The Flow configuration that will be used to unmarshall the configuration values into after binding pflags.
//	This needs to be done because pflags may override a configuration value.
//
// Returns:
//
//	error: if there is any error encountered binding pflags or unmarshalling the config struct, all errors are considered irrecoverable.
//	bool: true if --config-file flag was set and config file was loaded, false otherwise.
//
// Note: As configuration management is improved this func should accept the entire Flow config as the arg to unmarshall new config values into.
func BindPFlags(c *FlowConfig, flags *pflag.FlagSet) (error, bool) {
	if !flags.Parsed() {
		return fmt.Errorf("failed to bind flags to configuration values, pflags must be parsed before binding"), false
	}

	// update the config store values from config file if --config-file flag is set
	// if config file provided we will use values from the file and skip binding pflags
	err, overridden := overrideConfigFile(flags)
	if err != nil {
		return err, false
	}

	if !overridden {
		err = conf.BindPFlags(flags)
		if err != nil {
			return fmt.Errorf("failed to bind pflag set: %w", err), false
		}
		setAliases()
	}

	err = Unmarshall(c)
	if err != nil {
		return fmt.Errorf("failed to unmarshall the Flow config: %w", err), false
	}

	return nil, overridden
}

// Unmarshall unmarshalls the Flow configuration into the provided FlowConfig struct.
// Args:
//
//	flowConfig: the flow config struct used for unmarshalling.
//
// Returns:
//
//	error: if there is any error encountered unmarshalling the configuration, all errors are considered irrecoverable.
func Unmarshall(flowConfig *FlowConfig) error {
	err := conf.Unmarshal(flowConfig, func(decoderConfig *mapstructure.DecoderConfig) {
		// enforce all fields are set on the FlowConfig struct
		decoderConfig.ErrorUnset = true
		// currently the entire flow configuration has not been moved to this package
		// for now we all key's in the config which are unused.
		decoderConfig.ErrorUnused = false
	})
	if err != nil {
		return fmt.Errorf("failed to unmarshal network config: %w", err)
	}
	return nil
}

// Print prints current configuration keys and values.
// Returns:
// map[string]struct{}: map of keys to avoid printing if they were set by an config file.
// This is required because we still have other config values not migrated to the config package. When a
// config file is used currently only network-configs are set, we want to avoid printing the config file
// value and also the flag value.
func Print(info *zerolog.Event, flags *pflag.FlagSet) map[string]struct{} {
	// only print config values if they were overridden with a config file
	m := make(map[string]struct{})
	if flags.Lookup(configFileFlagName).Changed {
		for _, key := range conf.AllKeys() {
			info.Str(key, fmt.Sprintf("%v", conf.Get(key)))
			s := strings.Split(key, ".")
			if len(s) == 2 {
				m[s[1]] = struct{}{}
			} else {
				m[key] = struct{}{}
			}
		}
	}

	return m
}

// setAliases sets aliases for config sub packages. This should be done directly after pflags are bound to the configuration store.
// Upon initialization the conf will be loaded with the default config values, those values are then used as the default values for
// all the CLI flags, the CLI flags are then bound to the configuration store and at this point all aliases should be set if configuration
// keys do not match the CLI flags 1:1. ie: networking-connection-pruning -> network-config.networking-connection-pruning. After aliases
// are set the conf store will override values with any CLI flag values that are set as expected.
func setAliases() {
	err := network.SetAliases(conf)
	if err != nil {
		panic(fmt.Errorf("failed to set network aliases: %w", err))
	}
}

// overrideConfigFile overrides the default config file by reading in the config file at the path set
// by the --config-file and --config-file-name flags in our viper config store.
//
// Returns:
//
//	error: if there is any error encountered while reading new config file, all errors are considered irrecoverable.
//	bool: true if the config was overridden by the new config file, false otherwise or if an error is encountered reading the new config file.
func overrideConfigFile(flags *pflag.FlagSet) (error, bool) {
	configFileFlag := flags.Lookup(configFileFlagName)
	if configFileFlag.Changed {
		p := configFileFlag.Value.String()
		dirPath, fileName := splitConfigPath(p)
		conf.AddConfigPath(dirPath)
		conf.SetConfigName(fileName)
		err := conf.ReadInConfig()
		if err != nil {
			return fmt.Errorf("failed to read config file %s: %w", p, err), false
		}
		return nil, true
	}
	return nil, false
}

// getConfigNameFromPath returns the directory and name of the config file from the provided path string.
func splitConfigPath(path string) (string, string) {
	dir, name := filepath.Split(path)
	return dir, strings.Split(name, ".")[0]
}

func init() {
	buf := bytes.NewBufferString(configFile)
	conf.SetConfigType("yaml")
	if err := conf.ReadConfig(buf); err != nil {
		panic(fmt.Errorf("failed to initialize flow config failed to read in config file: %w", err))
	}
}
