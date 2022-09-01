// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// QCCreatedConsumer is an autogenerated mock type for the QCCreatedConsumer type
type QCCreatedConsumer struct {
	mock.Mock
}

// OnQcConstructedFromVotes provides a mock function with given fields: _a0
func (_m *QCCreatedConsumer) OnQcConstructedFromVotes(_a0 *flow.QuorumCertificate) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewQCCreatedConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewQCCreatedConsumer creates a new instance of QCCreatedConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQCCreatedConsumer(t mockConstructorTestingTNewQCCreatedConsumer) *QCCreatedConsumer {
	mock := &QCCreatedConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
