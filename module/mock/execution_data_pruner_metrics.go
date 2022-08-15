// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ExecutionDataPrunerMetrics is an autogenerated mock type for the ExecutionDataPrunerMetrics type
type ExecutionDataPrunerMetrics struct {
	mock.Mock
}

// Pruned provides a mock function with given fields: height, duration
func (_m *ExecutionDataPrunerMetrics) Pruned(height uint64, duration time.Duration) {
	_m.Called(height, duration)
}

type mockConstructorTestingTNewExecutionDataPrunerMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewExecutionDataPrunerMetrics creates a new instance of ExecutionDataPrunerMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExecutionDataPrunerMetrics(t mockConstructorTestingTNewExecutionDataPrunerMetrics) *ExecutionDataPrunerMetrics {
	mock := &ExecutionDataPrunerMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
