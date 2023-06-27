// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	common "github.com/onflow/cadence/runtime/common"

	mock "github.com/stretchr/testify/mock"

	sema "github.com/onflow/cadence/runtime/sema"

	stdlib "github.com/onflow/cadence/runtime/stdlib"
)

// AccountKeyUpdater is an autogenerated mock type for the AccountKeyUpdater type
type AccountKeyUpdater struct {
	mock.Mock
}

// AddAccountKey provides a mock function with given fields: runtimeAddress, publicKey, hashAlgo, weight
func (_m *AccountKeyUpdater) AddAccountKey(runtimeAddress common.Address, publicKey *stdlib.PublicKey, hashAlgo sema.HashAlgorithm, weight int) (*stdlib.AccountKey, error) {
	ret := _m.Called(runtimeAddress, publicKey, hashAlgo, weight)

	var r0 *stdlib.AccountKey
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, *stdlib.PublicKey, sema.HashAlgorithm, int) (*stdlib.AccountKey, error)); ok {
		return rf(runtimeAddress, publicKey, hashAlgo, weight)
	}
	if rf, ok := ret.Get(0).(func(common.Address, *stdlib.PublicKey, sema.HashAlgorithm, int) *stdlib.AccountKey); ok {
		r0 = rf(runtimeAddress, publicKey, hashAlgo, weight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stdlib.AccountKey)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, *stdlib.PublicKey, sema.HashAlgorithm, int) error); ok {
		r1 = rf(runtimeAddress, publicKey, hashAlgo, weight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeAccountKey provides a mock function with given fields: address, keyIndex
func (_m *AccountKeyUpdater) RevokeAccountKey(address common.Address, keyIndex int) (*stdlib.AccountKey, error) {
	ret := _m.Called(address, keyIndex)

	var r0 *stdlib.AccountKey
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, int) (*stdlib.AccountKey, error)); ok {
		return rf(address, keyIndex)
	}
	if rf, ok := ret.Get(0).(func(common.Address, int) *stdlib.AccountKey); ok {
		r0 = rf(address, keyIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stdlib.AccountKey)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, int) error); ok {
		r1 = rf(address, keyIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAccountKeyUpdater interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountKeyUpdater creates a new instance of AccountKeyUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountKeyUpdater(t mockConstructorTestingTNewAccountKeyUpdater) *AccountKeyUpdater {
	mock := &AccountKeyUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
