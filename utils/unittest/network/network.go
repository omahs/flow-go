package network

import (
	"fmt"

	"github.com/stretchr/testify/mock"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/network"
	"github.com/onflow/flow-go/network/channels"
	"github.com/onflow/flow-go/network/mocknetwork"
)

type EngineProcessFunc func(channels.Channel, flow.Identifier, interface{}) error
type PublishFunc func(channels.Channel, interface{}, ...flow.Identifier) error

// Conduit represents a mock conduit.

// Network represents a mock network. The implementation is not concurrency-safe.
type Network struct {
	mocknetwork.Network
	conduits    map[channels.Channel]*Conduit
	engines     map[channels.Channel]network.MessageProcessor
	publishFunc PublishFunc
}

var _ network.Network = (*Network)(nil)

// NewNetwork returns a new mock network.
func NewNetwork() *Network {
	return &Network{
		Network:  mocknetwork.Network{},
		conduits: make(map[channels.Channel]*Conduit),
		engines:  make(map[channels.Channel]network.MessageProcessor),
	}
}

// Register registers an engine with this mock network. If an engine is already registered on the
// given channel, this will return an error.
func (n *Network) Register(channel channels.Channel, engine network.MessageProcessor) (network.Conduit, error) {
	_, ok := n.engines[channel]
	if ok {
		return nil, fmt.Errorf("channel already registered: %s", channel)
	}

	n.engines[channel] = engine
	conduit := &Conduit{net: n, channel: channel}
	n.conduits[channel] = conduit

	return conduit, nil
}

// Send sends a message to the engine registered to the given channel on this mock network and returns
// an error if one occurs. If no engine is registered, this is a noop.
func (n *Network) Send(channel channels.Channel, originID flow.Identifier, event interface{}) error {
	if eng, ok := n.engines[channel]; ok {
		return eng.Process(channel, originID, event)
	}
	return nil
}

// OnPublish specifies the callback that should be executed when `Publish` is called on any conduits
// created by this mock network.
func (n *Network) OnPublish(publishFunc PublishFunc) *Network {
	n.publishFunc = publishFunc
	return n
}

// Engine represents a mock engine. The implementation is not concurrency-safe.
type Engine struct {
	mocknetwork.Engine
}

// NewEngine returns a new mock engine.
func NewEngine() *Engine {
	return &Engine{
		mocknetwork.Engine{},
	}
}

// OnProcess specifies the callback that should be executed when `Process` is called on this mock engine.
func (e *Engine) OnProcess(processFunc EngineProcessFunc) *Engine {
	e.On("Process", mock.AnythingOfType("channels.Channel"), mock.AnythingOfType("flow.Identifier"), mock.Anything).
		Return((func(channels.Channel, flow.Identifier, interface{}) error)(processFunc))

	return e
}
