// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TransactionMetrics is an autogenerated mock type for the TransactionMetrics type
type TransactionMetrics struct {
	mock.Mock
}

// TransactionExecuted provides a mock function with given fields: txID, when
func (_m *TransactionMetrics) TransactionExecuted(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionExpired provides a mock function with given fields: txID
func (_m *TransactionMetrics) TransactionExpired(txID flow.Identifier) {
	_m.Called(txID)
}

// TransactionFinalized provides a mock function with given fields: txID, when
func (_m *TransactionMetrics) TransactionFinalized(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionReceived provides a mock function with given fields: txID, when
func (_m *TransactionMetrics) TransactionReceived(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionResultFetched provides a mock function with given fields: dur, size
func (_m *TransactionMetrics) TransactionResultFetched(dur time.Duration, size int) {
	_m.Called(dur, size)
}

// TransactionSealed provides a mock function with given fields: txID, when
func (_m *TransactionMetrics) TransactionSealed(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionSubmissionFailed provides a mock function with no fields
func (_m *TransactionMetrics) TransactionSubmissionFailed() {
	_m.Called()
}

// NewTransactionMetrics creates a new instance of TransactionMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionMetrics {
	mock := &TransactionMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
