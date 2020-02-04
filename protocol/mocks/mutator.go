// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dapperlabs/flow-go/protocol (interfaces: Mutator)

// Package mocks is a generated GoMock package.
package mocks

import (
	flow "github.com/dapperlabs/flow-go/model/flow"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMutator is a mock of Mutator interface
type MockMutator struct {
	ctrl     *gomock.Controller
	recorder *MockMutatorMockRecorder
}

// MockMutatorMockRecorder is the mock recorder for MockMutator
type MockMutatorMockRecorder struct {
	mock *MockMutator
}

// NewMockMutator creates a new mock instance
func NewMockMutator(ctrl *gomock.Controller) *MockMutator {
	mock := &MockMutator{ctrl: ctrl}
	mock.recorder = &MockMutatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutator) EXPECT() *MockMutatorMockRecorder {
	return m.recorder
}

// Bootstrap mocks base method
func (m *MockMutator) Bootstrap(arg0 *flow.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrap indicates an expected call of Bootstrap
func (mr *MockMutatorMockRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockMutator)(nil).Bootstrap), arg0)
}

// Extend mocks base method
func (m *MockMutator) Extend(arg0 flow.Identifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extend", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Extend indicates an expected call of Extend
func (mr *MockMutatorMockRecorder) Extend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extend", reflect.TypeOf((*MockMutator)(nil).Extend), arg0)
}

// Finalize mocks base method
func (m *MockMutator) Finalize(arg0 flow.Identifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finalize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Finalize indicates an expected call of Finalize
func (mr *MockMutatorMockRecorder) Finalize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockMutator)(nil).Finalize), arg0)
}

// StorePayload mocks base method
func (m *MockMutator) StorePayload(arg0 *flow.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorePayload", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorePayload indicates an expected call of StorePayload
func (mr *MockMutatorMockRecorder) StorePayload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorePayload", reflect.TypeOf((*MockMutator)(nil).StorePayload), arg0)
}
