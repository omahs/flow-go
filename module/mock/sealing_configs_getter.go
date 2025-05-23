// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// SealingConfigsGetter is an autogenerated mock type for the SealingConfigsGetter type
type SealingConfigsGetter struct {
	mock.Mock
}

// ApprovalRequestsThresholdConst provides a mock function with no fields
func (_m *SealingConfigsGetter) ApprovalRequestsThresholdConst() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ApprovalRequestsThresholdConst")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ChunkAlphaConst provides a mock function with no fields
func (_m *SealingConfigsGetter) ChunkAlphaConst() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ChunkAlphaConst")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// EmergencySealingActiveConst provides a mock function with no fields
func (_m *SealingConfigsGetter) EmergencySealingActiveConst() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EmergencySealingActiveConst")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RequireApprovalsForSealConstructionDynamicValue provides a mock function with no fields
func (_m *SealingConfigsGetter) RequireApprovalsForSealConstructionDynamicValue() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequireApprovalsForSealConstructionDynamicValue")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// RequireApprovalsForSealVerificationConst provides a mock function with no fields
func (_m *SealingConfigsGetter) RequireApprovalsForSealVerificationConst() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequireApprovalsForSealVerificationConst")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// NewSealingConfigsGetter creates a new instance of SealingConfigsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSealingConfigsGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *SealingConfigsGetter {
	mock := &SealingConfigsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
