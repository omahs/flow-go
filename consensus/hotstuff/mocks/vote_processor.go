// Code generated by mockery v2.13.0. DO NOT EDIT.

package mocks

import (
	hotstuff "github.com/onflow/flow-go/consensus/hotstuff"
	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"
)

// VoteProcessor is an autogenerated mock type for the VoteProcessor type
type VoteProcessor struct {
	mock.Mock
}

// Process provides a mock function with given fields: vote
func (_m *VoteProcessor) Process(vote *model.Vote) error {
	ret := _m.Called(vote)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Vote) error); ok {
		r0 = rf(vote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *VoteProcessor) Status() hotstuff.VoteCollectorStatus {
	ret := _m.Called()

	var r0 hotstuff.VoteCollectorStatus
	if rf, ok := ret.Get(0).(func() hotstuff.VoteCollectorStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(hotstuff.VoteCollectorStatus)
	}

	return r0
}

type NewVoteProcessorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewVoteProcessor creates a new instance of VoteProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVoteProcessor(t NewVoteProcessorT) *VoteProcessor {
	mock := &VoteProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
