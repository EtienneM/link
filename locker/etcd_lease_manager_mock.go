// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/locker (interfaces: EtcdLeaseManager)

// Package locker is a generated GoMock package.
package locker

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clientv3 "go.etcd.io/etcd/clientv3"
)

// MockEtcdLeaseManager is a mock of EtcdLeaseManager interface
type MockEtcdLeaseManager struct {
	ctrl     *gomock.Controller
	recorder *MockEtcdLeaseManagerMockRecorder
}

// MockEtcdLeaseManagerMockRecorder is the mock recorder for MockEtcdLeaseManager
type MockEtcdLeaseManagerMockRecorder struct {
	mock *MockEtcdLeaseManager
}

// NewMockEtcdLeaseManager creates a new mock instance
func NewMockEtcdLeaseManager(ctrl *gomock.Controller) *MockEtcdLeaseManager {
	mock := &MockEtcdLeaseManager{ctrl: ctrl}
	mock.recorder = &MockEtcdLeaseManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEtcdLeaseManager) EXPECT() *MockEtcdLeaseManagerMockRecorder {
	return m.recorder
}

// GetLease mocks base method
func (m *MockEtcdLeaseManager) GetLease(arg0 context.Context) (clientv3.LeaseID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLease", arg0)
	ret0, _ := ret[0].(clientv3.LeaseID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLease indicates an expected call of GetLease
func (mr *MockEtcdLeaseManagerMockRecorder) GetLease(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLease", reflect.TypeOf((*MockEtcdLeaseManager)(nil).GetLease), arg0)
}

// MarkLeaseAsDirty mocks base method
func (m *MockEtcdLeaseManager) MarkLeaseAsDirty(arg0 context.Context, arg1 clientv3.LeaseID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkLeaseAsDirty", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkLeaseAsDirty indicates an expected call of MarkLeaseAsDirty
func (mr *MockEtcdLeaseManagerMockRecorder) MarkLeaseAsDirty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkLeaseAsDirty", reflect.TypeOf((*MockEtcdLeaseManager)(nil).MarkLeaseAsDirty), arg0, arg1)
}

// Start mocks base method
func (m *MockEtcdLeaseManager) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockEtcdLeaseManagerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEtcdLeaseManager)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockEtcdLeaseManager) Stop(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop
func (mr *MockEtcdLeaseManagerMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockEtcdLeaseManager)(nil).Stop), arg0)
}

// SubscribeToLeaseChange mocks base method
func (m *MockEtcdLeaseManager) SubscribeToLeaseChange(arg0 context.Context, arg1 LeaseChangedCallback) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToLeaseChange", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToLeaseChange indicates an expected call of SubscribeToLeaseChange
func (mr *MockEtcdLeaseManagerMockRecorder) SubscribeToLeaseChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToLeaseChange", reflect.TypeOf((*MockEtcdLeaseManager)(nil).SubscribeToLeaseChange), arg0, arg1)
}

// UnsubscribeToLeaseChange mocks base method
func (m *MockEtcdLeaseManager) UnsubscribeToLeaseChange(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeToLeaseChange", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeToLeaseChange indicates an expected call of UnsubscribeToLeaseChange
func (mr *MockEtcdLeaseManagerMockRecorder) UnsubscribeToLeaseChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeToLeaseChange", reflect.TypeOf((*MockEtcdLeaseManager)(nil).UnsubscribeToLeaseChange), arg0, arg1)
}
