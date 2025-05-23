// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// MinimumCadenceRequiredVersion is an autogenerated mock type for the MinimumCadenceRequiredVersion type
type MinimumCadenceRequiredVersion struct {
	mock.Mock
}

// MinimumRequiredVersion provides a mock function with no fields
func (_m *MinimumCadenceRequiredVersion) MinimumRequiredVersion() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MinimumRequiredVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMinimumCadenceRequiredVersion creates a new instance of MinimumCadenceRequiredVersion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMinimumCadenceRequiredVersion(t interface {
	mock.TestingT
	Cleanup(func())
}) *MinimumCadenceRequiredVersion {
	mock := &MinimumCadenceRequiredVersion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
