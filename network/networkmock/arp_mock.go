// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/v2/network (interfaces: ARP)

// Package networkmock is a generated GoMock package.
package networkmock

import (
	reflect "reflect"

	network "github.com/Scalingo/link/v2/network"
	gomock "github.com/golang/mock/gomock"
)

// MockARP is a mock of ARP interface.
type MockARP struct {
	ctrl     *gomock.Controller
	recorder *MockARPMockRecorder
}

// MockARPMockRecorder is the mock recorder for MockARP.
type MockARPMockRecorder struct {
	mock *MockARP
}

// NewMockARP creates a new mock instance.
func NewMockARP(ctrl *gomock.Controller) *MockARP {
	mock := &MockARP{ctrl: ctrl}
	mock.recorder = &MockARPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockARP) EXPECT() *MockARPMockRecorder {
	return m.recorder
}

// GratuitousArp mocks base method.
func (m *MockARP) GratuitousArp(arg0 network.GratuitousArpRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GratuitousArp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GratuitousArp indicates an expected call of GratuitousArp.
func (mr *MockARPMockRecorder) GratuitousArp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GratuitousArp", reflect.TypeOf((*MockARP)(nil).GratuitousArp), arg0)
}
