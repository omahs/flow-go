// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package execute

import (
	"github.com/dapperlabs/bamboo-node/internal/nodes/execute/config"
	"github.com/dapperlabs/bamboo-node/internal/nodes/execute/controllers"
	"github.com/dapperlabs/bamboo-node/internal/nodes/execute/data"
)

// Injectors from wire.go:

func InitializeServer() (*Server, error) {
	configConfig := config.New()
	dal := data.New(configConfig)
	controller := controllers.NewController(dal)
	server, err := NewServer(dal, configConfig, controller)
	if err != nil {
		return nil, err
	}
	return server, nil
}
