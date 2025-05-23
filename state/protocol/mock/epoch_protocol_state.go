// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	protocol "github.com/onflow/flow-go/state/protocol"
)

// EpochProtocolState is an autogenerated mock type for the EpochProtocolState type
type EpochProtocolState struct {
	mock.Mock
}

// Clustering provides a mock function with no fields
func (_m *EpochProtocolState) Clustering() (flow.ClusterList, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clustering")
	}

	var r0 flow.ClusterList
	var r1 error
	if rf, ok := ret.Get(0).(func() (flow.ClusterList, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() flow.ClusterList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.ClusterList)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DKG provides a mock function with no fields
func (_m *EpochProtocolState) DKG() (protocol.DKG, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DKG")
	}

	var r0 protocol.DKG
	var r1 error
	if rf, ok := ret.Get(0).(func() (protocol.DKG, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() protocol.DKG); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.DKG)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Entry provides a mock function with no fields
func (_m *EpochProtocolState) Entry() *flow.RichEpochStateEntry {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Entry")
	}

	var r0 *flow.RichEpochStateEntry
	if rf, ok := ret.Get(0).(func() *flow.RichEpochStateEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.RichEpochStateEntry)
		}
	}

	return r0
}

// Epoch provides a mock function with no fields
func (_m *EpochProtocolState) Epoch() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Epoch")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// EpochCommit provides a mock function with no fields
func (_m *EpochProtocolState) EpochCommit() *flow.EpochCommit {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EpochCommit")
	}

	var r0 *flow.EpochCommit
	if rf, ok := ret.Get(0).(func() *flow.EpochCommit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.EpochCommit)
		}
	}

	return r0
}

// EpochExtensions provides a mock function with no fields
func (_m *EpochProtocolState) EpochExtensions() []flow.EpochExtension {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EpochExtensions")
	}

	var r0 []flow.EpochExtension
	if rf, ok := ret.Get(0).(func() []flow.EpochExtension); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.EpochExtension)
		}
	}

	return r0
}

// EpochFallbackTriggered provides a mock function with no fields
func (_m *EpochProtocolState) EpochFallbackTriggered() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EpochFallbackTriggered")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EpochPhase provides a mock function with no fields
func (_m *EpochProtocolState) EpochPhase() flow.EpochPhase {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EpochPhase")
	}

	var r0 flow.EpochPhase
	if rf, ok := ret.Get(0).(func() flow.EpochPhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flow.EpochPhase)
	}

	return r0
}

// EpochSetup provides a mock function with no fields
func (_m *EpochProtocolState) EpochSetup() *flow.EpochSetup {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EpochSetup")
	}

	var r0 *flow.EpochSetup
	if rf, ok := ret.Get(0).(func() *flow.EpochSetup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.EpochSetup)
		}
	}

	return r0
}

// GlobalParams provides a mock function with no fields
func (_m *EpochProtocolState) GlobalParams() protocol.GlobalParams {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GlobalParams")
	}

	var r0 protocol.GlobalParams
	if rf, ok := ret.Get(0).(func() protocol.GlobalParams); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.GlobalParams)
		}
	}

	return r0
}

// Identities provides a mock function with no fields
func (_m *EpochProtocolState) Identities() flow.GenericIdentityList[flow.Identity] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Identities")
	}

	var r0 flow.GenericIdentityList[flow.Identity]
	if rf, ok := ret.Get(0).(func() flow.GenericIdentityList[flow.Identity]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.GenericIdentityList[flow.Identity])
		}
	}

	return r0
}

// PreviousEpochExists provides a mock function with no fields
func (_m *EpochProtocolState) PreviousEpochExists() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PreviousEpochExists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewEpochProtocolState creates a new instance of EpochProtocolState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEpochProtocolState(t interface {
	mock.TestingT
	Cleanup(func())
}) *EpochProtocolState {
	mock := &EpochProtocolState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
