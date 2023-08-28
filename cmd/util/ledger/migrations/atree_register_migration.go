package migrations

import (
	"context"
	"fmt"
	"github.com/onflow/atree"
	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/rs/zerolog"
	"io"

	"github.com/onflow/flow-go/cmd/util/ledger/reporters"
	"github.com/onflow/flow-go/cmd/util/ledger/util"
	"github.com/onflow/flow-go/fvm/environment"
	"github.com/onflow/flow-go/fvm/storage/derived"
	"github.com/onflow/flow-go/fvm/storage/logical"
	"github.com/onflow/flow-go/fvm/storage/state"
	"github.com/onflow/flow-go/fvm/tracing"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
)

func MigrateAtreeRegisters(
	log zerolog.Logger,
	rwf reporters.ReportWriterFactory,
	nWorker int,
) func([]ledger.Payload) ([]ledger.Payload, error) {
	return CreateAccountBasedMigration(
		log,
		func(allRegisters []ledger.Payload, nWorker int) (AccountMigrator, error) {
			return NewMigrator(
				log.With().Str("component", "atree-register-migrator").Logger(),
				rwf,
				allRegisters,
			)
		},
		nWorker,
	)
}

// AtreeRegisterMigrator is a migrator that converts the storage of an account from the
// old atree format to the new atree format.
// Account "storage used" should be correctly updated after the migration.
type AtreeRegisterMigrator struct {
	log zerolog.Logger

	snapshot *util.PayloadSnapshot

	rw      reporters.ReportWriter
	sampler zerolog.Logger
}

var _ AccountMigrator = (*AtreeRegisterMigrator)(nil)
var _ io.Closer = (*AtreeRegisterMigrator)(nil)

func NewMigrator(
	log zerolog.Logger,
	rwf reporters.ReportWriterFactory,
	allRegisters []ledger.Payload,
) (*AtreeRegisterMigrator, error) {
	// creating a snapshot of all the registers will take a while
	snapshot, err := util.NewPayloadSnapshot(allRegisters)
	if err != nil {
		return nil, fmt.Errorf("failed to create payload snapshot: %w", err)
	}
	log.Info().Msgf("created snapshot")

	migrator := &AtreeRegisterMigrator{
		snapshot: snapshot,

		log: log,

		rw: rwf.ReportWriter("atree-register-migrator"),
	}

	return migrator, nil
}

func (m *AtreeRegisterMigrator) Close() error {
	m.rw.Close()
	return nil
}

func (m *AtreeRegisterMigrator) MigratePayloads(
	_ context.Context,
	address common.Address,
	payloads []ledger.Payload,
) ([]ledger.Payload, error) {
	// don't migrate the zero address
	// these are the non-account registers and are not atrees
	if address == common.ZeroAddress {
		return payloads, nil
	}

	err := m.checkStorageHealth(address, payloads)
	if err != nil {
		return nil, fmt.Errorf("storage health issues for address %s: %w", address.Hex(), err)
	}

	// Do the storage conversion
	changes, err := m.migrateAccountStorage(address)
	if err != nil {
		return nil, fmt.Errorf("failed to convert storage for address %s: %w", address.Hex(), err)
	}

	return m.validateChangesAndCreateNewRegisters(payloads, address, changes)
}

func (m *AtreeRegisterMigrator) migrateAccountStorage(
	address common.Address,
) (map[flow.RegisterID]flow.RegisterValue, error) {

	// create all the runtime components we need for the migration
	r, err := m.newMigratorRuntime()
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator runtime: %w", err)
	}

	// iterate through all domains and migrate them
	for _, domain := range domains {
		err := m.convertStorageDomain(address, r.Storage, r.Interpreter, domain)
		if err != nil {
			return nil, fmt.Errorf("failed to convert storage domain %s : %w", domain, err)
		}
	}

	// commit the storage changes
	err = r.Storage.Commit(r.Interpreter, true)
	if err != nil {
		return nil, fmt.Errorf("failed to commit storage: %w", err)
	}

	// finalize the transaction
	result, err := r.TransactionState.FinalizeMainTransaction()
	if err != nil {
		return nil, fmt.Errorf("failed to finalize main transaction: %w", err)
	}

	return result.WriteSet, nil
}

func (m *AtreeRegisterMigrator) convertStorageDomain(
	address common.Address,
	storage *runtime.Storage,
	inter *interpreter.Interpreter,
	domain string,
) error {
	storageMap := storage.GetStorageMap(address, domain, false)
	if storageMap == nil {
		// no storage for this domain
		return nil
	}

	iterator := storageMap.Iterator(util.NopMemoryGauge{})
	keys := make([]interpreter.StringStorageMapKey, 0)
	// to be safe avoid modifying the map while iterating
	for {
		key := iterator.NextKey()
		if key == nil {
			break
		}

		stringKey, ok := key.(interpreter.StringAtreeValue)
		if !ok {
			return fmt.Errorf("invalid key type %T, expected interpreter.StringAtreeValue", key)
		}

		keys = append(keys, interpreter.StringStorageMapKey(stringKey))
	}

	for _, key := range keys {
		err := func() error {
			var value interpreter.Value

			err := capturePanic(func() {
				value = storageMap.ReadValue(util.NopMemoryGauge{}, key)
			})
			if err != nil {
				return fmt.Errorf("failed to read value for key %s: %w", key, err)
			}

			err = capturePanic(func() {
				// force the value to be read entirely
				value = value.Clone(inter)
			})
			if err != nil {
				return fmt.Errorf("failed to clone value for key %s: %w", key, err)
			}

			err = capturePanic(func() {
				// set value will first purge the old value
				storageMap.SetValue(inter, key, value)
			})
			if err != nil {
				return fmt.Errorf("failed to set value for key %s: %w", key, err)
			}

			return nil
		}()
		if err != nil {
			m.log.Debug().Err(err).Msgf("failed to migrate key %s", key)

			m.rw.Write(migrationError{
				Address: address.Hex(),
				Key:     string(key),
				Kind:    "migration_failure",
				Msg:     err.Error(),
			})
		}
	}

	return nil
}

func (m *AtreeRegisterMigrator) newMigratorRuntime() (
	*migratorRuntime,
	error,
) {
	// the registers will be read concurrently by multiple workers, but won't be modified
	transactionState := state.NewTransactionState(m.snapshot, state.DefaultParameters())
	accounts := environment.NewAccounts(transactionState)

	derivedBlockData := derived.NewEmptyDerivedBlockData(logical.EndOfBlockExecutionTime)

	programs :=
		environment.NewPrograms(
			tracing.NewMockTracerSpan(),
			util.NopMeter{},
			environment.NoopMetricsReporter{},
			struct {
				state.NestedTransactionPreparer
				derived.DerivedTransactionPreparer
			}{
				NestedTransactionPreparer: transactionState,
				// TODO: reuse this
				DerivedTransactionPreparer: derivedBlockData.NewSnapshotReadDerivedTransactionData(),
			},
			accounts,
		)

	accountsAtreeLedger := util.NewAccountsAtreeLedger(accounts)
	storage := runtime.NewStorage(accountsAtreeLedger, util.NopMemoryGauge{})

	env := runtime.NewBaseInterpreterEnvironment(runtime.Config{
		AccountLinkingEnabled: true,
		// Attachments are enabled everywhere except for Mainnet
		AttachmentsEnabled: true,
		// Capability Controllers are enabled everywhere except for Mainnet
		CapabilityControllersEnabled: true,
	})

	ri := util.MigrationRuntimeInterface{
		Accounts: accounts,
		Programs: programs,
	}

	env.Configure(
		ri,
		runtime.NewCodesAndPrograms(),
		storage,
		runtime.NewCoverageReport(),
	)

	inter, err := interpreter.NewInterpreter(
		nil,
		nil,
		env.InterpreterConfig)
	if err != nil {
		return nil, err
	}

	return &migratorRuntime{
		TransactionState: transactionState,
		Interpreter:      inter,
		Storage:          storage,
	}, nil
}

type migratorRuntime struct {
	TransactionState state.NestedTransactionPreparer
	Interpreter      *interpreter.Interpreter
	Storage          *runtime.Storage
}

func (m *AtreeRegisterMigrator) validateChangesAndCreateNewRegisters(
	payloads []ledger.Payload,
	address common.Address,
	changes map[flow.RegisterID]flow.RegisterValue,
) ([]ledger.Payload, error) {
	originalPayloadsSnapshot, err := util.NewPayloadSnapshot(payloads)
	if err != nil {
		return nil, fmt.Errorf("failed to create payload snapshot: %w", err)
	}
	originalPayloads := originalPayloadsSnapshot.Payloads
	newPayloads := make([]ledger.Payload, 0, len(originalPayloads))

	for id, value := range changes {
		// delete all values that were changed from the original payloads
		delete(originalPayloads, id)

		if len(value) == 0 {
			// value was deleted
			continue
		}
		ownerAddress, err := common.BytesToAddress([]byte(id.Owner))
		if err != nil {
			return nil, fmt.Errorf("failed to convert owner address: %w", err)
		}

		if ownerAddress.Hex() != address.Hex() {
			// something was changed that does not belong to this account. Log it.
			m.log.Error().
				Str("key", id.String()).
				Str("owner_address", ownerAddress.Hex()).
				Str("account", address.Hex()).
				Msg("key is part of the change set, but is for a different account")

			return nil, fmt.Errorf("register for a different account was produced during migration")
		}

		newPayloads = append(newPayloads, *ledger.NewPayload(util.RegisterIDToKey(id), value))
	}

	// add all values that were not changed
	for id, value := range originalPayloads {
		if len(value) == 0 {
			// this is strange, but we don't want to add empty values. Log it.
			m.log.Warn().Msgf("empty value for key %s", id)
			continue
		}
		if id.IsInternalState() {
			// this is expected. Move it to the new payload
			newPayloads = append(newPayloads, *ledger.NewPayload(util.RegisterIDToKey(id), value))
			continue
		}

		isADomainKey := false
		for _, domain := range domains {
			if id.Key == domain {
				isADomainKey = true
				break
			}
		}
		if isADomainKey {
			// TODO: check if this is really expected
			// this is expected. Move it to the new payload
			newPayloads = append(newPayloads, *ledger.NewPayload(util.RegisterIDToKey(id), value))
			continue
		}

		// something was not moved. Log it.
		m.log.Debug().
			Str("key", id.String()).
			Str("account", address.Hex()).
			Str("value", fmt.Sprintf("%x", value)).
			Msg("Key was not migrated")
		m.rw.Write(migrationError{
			Address: address.Hex(),
			Key:     id.String(),
			Kind:    "not_migrated",
			Msg:     fmt.Sprintf("%x", value),
		})
	}
	return newPayloads, nil
}

func (m *AtreeRegisterMigrator) checkStorageHealth(
	address common.Address,
	payloads []ledger.Payload,
) error {
	snapshot, err := util.NewPayloadSnapshot(payloads)
	if err != nil {
		return fmt.Errorf("failed to create payload snapshot: %w", err)
	}

	transactionState := state.NewTransactionState(snapshot, state.DefaultParameters())
	accounts := environment.NewAccounts(transactionState)

	accountsAtreeLedger := util.NewAccountsAtreeLedger(accounts)
	storage := runtime.NewStorage(accountsAtreeLedger, util.NopMemoryGauge{})

	rootSlabs, err := atree.CheckStorageHealth(storage, -1)
	if err != nil {
		m.log.Info().
			Err(err).
			Str("account", address.Hex()).
			Msg("Account storage health issue")
		m.rw.Write(migrationError{
			Address: address.Hex(),
			Key:     "",
			Kind:    "storage_health_problem_1",
			Msg:     err.Error(),
		})
	}

	if len(rootSlabs) > 4 {
		m.log.Info().
			Err(err).
			Str("account", address.Hex()).
			Msg("To many root slabs")
		m.rw.Write(migrationError{
			Address: address.Hex(),
			Key:     "",
			Kind:    "storage_health_problem_2",
			Msg:     err.Error(),
		})
	}

	err = storage.CheckHealth()
	if err != nil {
		m.log.Info().
			Err(err).
			Str("account", address.Hex()).
			Msg("Account storage health issue")
		m.rw.Write(migrationError{
			Address: address.Hex(),
			Key:     "",
			Kind:    "storage_health_problem_3",
			Msg:     err.Error(),
		})
	}
	return nil

}

// capturePanic captures panics and converts them to errors
// this is needed for some cadence functions that panic on error
func capturePanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	f()

	return
}

// convert all domains
var domains = []string{
	common.PathDomainStorage.Identifier(),
	common.PathDomainPrivate.Identifier(),
	common.PathDomainPublic.Identifier(),
	runtime.StorageDomainContract,
}

// migrationError is a struct for reporting errors
type migrationError struct {
	Address string
	Key     string
	Kind    string
	Msg     string
}
