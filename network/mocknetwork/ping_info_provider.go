// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocknetwork

import mock "github.com/stretchr/testify/mock"

// PingInfoProvider is an autogenerated mock type for the PingInfoProvider type
type PingInfoProvider struct {
	mock.Mock
}

// HotstuffView provides a mock function with given fields:
func (_m *PingInfoProvider) HotstuffView() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// SealedBlockHeight provides a mock function with given fields:
func (_m *PingInfoProvider) SealedBlockHeight() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// SoftwareVersion provides a mock function with given fields:
func (_m *PingInfoProvider) SoftwareVersion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewPingInfoProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewPingInfoProvider creates a new instance of PingInfoProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPingInfoProvider(t mockConstructorTestingTNewPingInfoProvider) *PingInfoProvider {
	mock := &PingInfoProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
