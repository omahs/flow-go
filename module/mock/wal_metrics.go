// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// WALMetrics is an autogenerated mock type for the WALMetrics type
type WALMetrics struct {
	mock.Mock
}

// ExecutionCheckpointSize provides a mock function with given fields: bytes
func (_m *WALMetrics) ExecutionCheckpointSize(bytes uint64) {
	_m.Called(bytes)
}

// NewWALMetrics creates a new instance of WALMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWALMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *WALMetrics {
	mock := &WALMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
