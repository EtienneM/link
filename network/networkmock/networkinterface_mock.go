// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/network (interfaces: NetworkInterface)

// Package networkmock is a generated GoMock package.
package networkmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNetworkInterface is a mock of NetworkInterface interface
type MockNetworkInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkInterfaceMockRecorder
}

// MockNetworkInterfaceMockRecorder is the mock recorder for MockNetworkInterface
type MockNetworkInterfaceMockRecorder struct {
	mock *MockNetworkInterface
}

// NewMockNetworkInterface creates a new mock instance
func NewMockNetworkInterface(ctrl *gomock.Controller) *MockNetworkInterface {
	mock := &MockNetworkInterface{ctrl: ctrl}
	mock.recorder = &MockNetworkInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkInterface) EXPECT() *MockNetworkInterfaceMockRecorder {
	return m.recorder
}

// EnsureIP mocks base method
func (m *MockNetworkInterface) EnsureIP(arg0 string) error {
	ret := m.ctrl.Call(m, "EnsureIP", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIP indicates an expected call of EnsureIP
func (mr *MockNetworkInterfaceMockRecorder) EnsureIP(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIP", reflect.TypeOf((*MockNetworkInterface)(nil).EnsureIP), arg0)
}

// RemoveIP mocks base method
func (m *MockNetworkInterface) RemoveIP(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveIP", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIP indicates an expected call of RemoveIP
func (mr *MockNetworkInterfaceMockRecorder) RemoveIP(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIP", reflect.TypeOf((*MockNetworkInterface)(nil).RemoveIP), arg0)
}
