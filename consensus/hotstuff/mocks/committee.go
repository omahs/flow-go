// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	hotstuff "github.com/onflow/flow-go/consensus/hotstuff"
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Committee is an autogenerated mock type for the Committee type
type Committee struct {
	mock.Mock
}

// DKG provides a mock function with given fields: blockID
func (_m *Committee) DKG(blockID flow.Identifier) (hotstuff.DKG, error) {
	ret := _m.Called(blockID)

	var r0 hotstuff.DKG
	if rf, ok := ret.Get(0).(func(flow.Identifier) hotstuff.DKG); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hotstuff.DKG)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Identities provides a mock function with given fields: blockID
func (_m *Committee) Identities(blockID flow.Identifier) (flow.IdentityList, error) {
	ret := _m.Called(blockID)

	var r0 flow.IdentityList
	if rf, ok := ret.Get(0).(func(flow.Identifier) flow.IdentityList); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Identity provides a mock function with given fields: blockID, participantID
func (_m *Committee) Identity(blockID flow.Identifier, participantID flow.Identifier) (*flow.Identity, error) {
	ret := _m.Called(blockID, participantID)

	var r0 *flow.Identity
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier) *flow.Identity); ok {
		r0 = rf(blockID, participantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier, flow.Identifier) error); ok {
		r1 = rf(blockID, participantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaderForView provides a mock function with given fields: view
func (_m *Committee) LeaderForView(view uint64) (flow.Identifier, error) {
	ret := _m.Called(view)

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func(uint64) flow.Identifier); ok {
		r0 = rf(view)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(view)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Self provides a mock function with given fields:
func (_m *Committee) Self() flow.Identifier {
	ret := _m.Called()

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	return r0
}

// NewCommittee creates a new instance of Committee. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommittee(t testing.TB) *Committee {
	mock := &Committee{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
