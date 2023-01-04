// Code generated by MockGen. DO NOT EDIT.
// Source: address.go

// Package logistic is l generated GoMock package.
package logistic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocator is l mock of Locator interface.
type MockLocator struct {
	ctrl     *gomock.Controller
	recorder *MockLocatorMockRecorder
}

// MockLocatorMockRecorder is the mock recorder for MockLocator.
type MockLocatorMockRecorder struct {
	mock *MockLocator
}

// NewMockLocator creates l new mock instance.
func NewMockLocator(ctrl *gomock.Controller) *MockLocator {
	mock := &MockLocator{ctrl: ctrl}
	mock.recorder = &MockLocatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocator) EXPECT() *MockLocatorMockRecorder {
	return m.recorder
}

// GetAddress mocks base method.
func (m *MockLocator) GetAddress(customerID int64, addressType string) *Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", customerID, addressType)
	ret0, _ := ret[0].(*Address)
	return ret0
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockLocatorMockRecorder) GetAddress(customerID, addressType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockLocator)(nil).GetAddress), customerID, addressType)
}