// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/api (interfaces: Client)

// Package apimock is a generated GoMock package.
package apimock

import (
	context "context"
	reflect "reflect"

	api "github.com/Scalingo/link/api"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddIP mocks base method
func (m *MockClient) AddIP(arg0 context.Context, arg1 string, arg2 api.AddIPParams) (api.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIP", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIP indicates an expected call of AddIP
func (mr *MockClientMockRecorder) AddIP(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIP", reflect.TypeOf((*MockClient)(nil).AddIP), arg0, arg1, arg2)
}

// GetIP mocks base method
func (m *MockClient) GetIP(arg0 context.Context, arg1 string) (api.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIP", arg0, arg1)
	ret0, _ := ret[0].(api.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIP indicates an expected call of GetIP
func (mr *MockClientMockRecorder) GetIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIP", reflect.TypeOf((*MockClient)(nil).GetIP), arg0, arg1)
}

// ListIPs mocks base method
func (m *MockClient) ListIPs(arg0 context.Context) ([]api.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIPs", arg0)
	ret0, _ := ret[0].([]api.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIPs indicates an expected call of ListIPs
func (mr *MockClientMockRecorder) ListIPs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIPs", reflect.TypeOf((*MockClient)(nil).ListIPs), arg0)
}

// RemoveIP mocks base method
func (m *MockClient) RemoveIP(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIP", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIP indicates an expected call of RemoveIP
func (mr *MockClientMockRecorder) RemoveIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIP", reflect.TypeOf((*MockClient)(nil).RemoveIP), arg0, arg1)
}

// TryGetLock mocks base method
func (m *MockClient) TryGetLock(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryGetLock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TryGetLock indicates an expected call of TryGetLock
func (mr *MockClientMockRecorder) TryGetLock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryGetLock", reflect.TypeOf((*MockClient)(nil).TryGetLock), arg0, arg1)
}

// Version mocks base method
func (m *MockClient) Version(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockClientMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockClient)(nil).Version), arg0)
}
