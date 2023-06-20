// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	context "context"

	connmgr "github.com/libp2p/go-libp2p/core/connmgr"

	host "github.com/libp2p/go-libp2p/core/host"

	madns "github.com/multiformats/go-multiaddr-dns"

	mock "github.com/stretchr/testify/mock"

	module "github.com/onflow/flow-go/module"

	network "github.com/libp2p/go-libp2p/core/network"

	p2p "github.com/onflow/flow-go/network/p2p"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	routing "github.com/libp2p/go-libp2p/core/routing"

	time "time"
)

// NodeBuilder is an autogenerated mock type for the NodeBuilder type
type NodeBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *NodeBuilder) Build() (p2p.LibP2PNode, error) {
	ret := _m.Called()

	var r0 p2p.LibP2PNode
	var r1 error
	if rf, ok := ret.Get(0).(func() (p2p.LibP2PNode, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() p2p.LibP2PNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.LibP2PNode)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableGossipSubPeerScoring provides a mock function with given fields: _a0, _a1
func (_m *NodeBuilder) EnableGossipSubPeerScoring(_a0 module.IdentityProvider, _a1 *p2p.PeerScoringConfig) p2p.NodeBuilder {
	ret := _m.Called(_a0, _a1)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(module.IdentityProvider, *p2p.PeerScoringConfig) p2p.NodeBuilder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetBasicResolver provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetBasicResolver(_a0 madns.BasicResolver) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(madns.BasicResolver) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetConnectionGater provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetConnectionGater(_a0 p2p.ConnectionGater) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(p2p.ConnectionGater) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetConnectionManager provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetConnectionManager(_a0 connmgr.ConnManager) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(connmgr.ConnManager) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetCreateNode provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetCreateNode(_a0 p2p.CreateNodeFunc) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(p2p.CreateNodeFunc) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetGossipSubFactory provides a mock function with given fields: _a0, _a1
func (_m *NodeBuilder) SetGossipSubFactory(_a0 p2p.GossipSubFactoryFunc, _a1 p2p.GossipSubAdapterConfigFunc) p2p.NodeBuilder {
	ret := _m.Called(_a0, _a1)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(p2p.GossipSubFactoryFunc, p2p.GossipSubAdapterConfigFunc) p2p.NodeBuilder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetGossipSubRpcInspectorSuite provides a mock function with given fields: _a0
func (_m *NodeBuilder) OverrideDefaultInspectorSuite(_a0 p2p.GossipSubInspectorSuite) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(p2p.GossipSubInspectorSuite) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetGossipSubScoreTracerInterval provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetGossipSubScoreTracerInterval(_a0 time.Duration) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(time.Duration) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetGossipSubTracer provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetGossipSubTracer(_a0 p2p.PubSubTracer) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(p2p.PubSubTracer) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetPeerManagerOptions provides a mock function with given fields: _a0, _a1
func (_m *NodeBuilder) SetPeerManagerOptions(_a0 bool, _a1 time.Duration) p2p.NodeBuilder {
	ret := _m.Called(_a0, _a1)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(bool, time.Duration) p2p.NodeBuilder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetRateLimiterDistributor provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetRateLimiterDistributor(_a0 p2p.UnicastRateLimiterDistributor) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(p2p.UnicastRateLimiterDistributor) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetResourceManager provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetResourceManager(_a0 network.ResourceManager) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(network.ResourceManager) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetRoutingSystem provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetRoutingSystem(_a0 func(context.Context, host.Host) (routing.Routing, error)) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(func(context.Context, host.Host) (routing.Routing, error)) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetStreamCreationRetryInterval provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetStreamCreationRetryInterval(_a0 time.Duration) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(time.Duration) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

// SetSubscriptionFilter provides a mock function with given fields: _a0
func (_m *NodeBuilder) SetSubscriptionFilter(_a0 pubsub.SubscriptionFilter) p2p.NodeBuilder {
	ret := _m.Called(_a0)

	var r0 p2p.NodeBuilder
	if rf, ok := ret.Get(0).(func(pubsub.SubscriptionFilter) p2p.NodeBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeBuilder)
		}
	}

	return r0
}

type mockConstructorTestingTNewNodeBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewNodeBuilder creates a new instance of NodeBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNodeBuilder(t mockConstructorTestingTNewNodeBuilder) *NodeBuilder {
	mock := &NodeBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
