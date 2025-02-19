// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// GossipSubRpcValidationInspectorMetrics is an autogenerated mock type for the GossipSubRpcValidationInspectorMetrics type
type GossipSubRpcValidationInspectorMetrics struct {
	mock.Mock
}

// AsyncProcessingFinished provides a mock function with given fields: msgType, duration
func (_m *GossipSubRpcValidationInspectorMetrics) AsyncProcessingFinished(msgType string, duration time.Duration) {
	_m.Called(msgType, duration)
}

// AsyncProcessingStarted provides a mock function with given fields: msgType
func (_m *GossipSubRpcValidationInspectorMetrics) AsyncProcessingStarted(msgType string) {
	_m.Called(msgType)
}

// BlockingPreProcessingFinished provides a mock function with given fields: msgType, sampleSize, duration
func (_m *GossipSubRpcValidationInspectorMetrics) BlockingPreProcessingFinished(msgType string, sampleSize uint, duration time.Duration) {
	_m.Called(msgType, sampleSize, duration)
}

// BlockingPreProcessingStarted provides a mock function with given fields: msgType, sampleSize
func (_m *GossipSubRpcValidationInspectorMetrics) BlockingPreProcessingStarted(msgType string, sampleSize uint) {
	_m.Called(msgType, sampleSize)
}

type mockConstructorTestingTNewGossipSubRpcValidationInspectorMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipSubRpcValidationInspectorMetrics creates a new instance of GossipSubRpcValidationInspectorMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipSubRpcValidationInspectorMetrics(t mockConstructorTestingTNewGossipSubRpcValidationInspectorMetrics) *GossipSubRpcValidationInspectorMetrics {
	mock := &GossipSubRpcValidationInspectorMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
