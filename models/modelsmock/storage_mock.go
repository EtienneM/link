// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/v2/models (interfaces: Storage)

// Package modelsmock is a generated GoMock package.
package modelsmock

import (
	context "context"
	reflect "reflect"

	models "github.com/Scalingo/link/v2/models"
	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddIP mocks base method.
func (m *MockStorage) AddIP(arg0 context.Context, arg1 models.IP) (models.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIP", arg0, arg1)
	ret0, _ := ret[0].(models.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIP indicates an expected call of AddIP.
func (mr *MockStorageMockRecorder) AddIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIP", reflect.TypeOf((*MockStorage)(nil).AddIP), arg0, arg1)
}

// GetCurrentHost mocks base method.
func (m *MockStorage) GetCurrentHost(arg0 context.Context) (models.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentHost", arg0)
	ret0, _ := ret[0].(models.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentHost indicates an expected call of GetCurrentHost.
func (mr *MockStorageMockRecorder) GetCurrentHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentHost", reflect.TypeOf((*MockStorage)(nil).GetCurrentHost), arg0)
}

// GetIPHosts mocks base method.
func (m *MockStorage) GetIPHosts(arg0 context.Context, arg1 models.IP) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPHosts", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPHosts indicates an expected call of GetIPHosts.
func (mr *MockStorageMockRecorder) GetIPHosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPHosts", reflect.TypeOf((*MockStorage)(nil).GetIPHosts), arg0, arg1)
}

// GetIPs mocks base method.
func (m *MockStorage) GetIPs(arg0 context.Context) ([]models.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPs", arg0)
	ret0, _ := ret[0].([]models.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPs indicates an expected call of GetIPs.
func (mr *MockStorageMockRecorder) GetIPs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPs", reflect.TypeOf((*MockStorage)(nil).GetIPs), arg0)
}

// LinkIPWithCurrentHost mocks base method.
func (m *MockStorage) LinkIPWithCurrentHost(arg0 context.Context, arg1 models.IP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkIPWithCurrentHost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkIPWithCurrentHost indicates an expected call of LinkIPWithCurrentHost.
func (mr *MockStorageMockRecorder) LinkIPWithCurrentHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkIPWithCurrentHost", reflect.TypeOf((*MockStorage)(nil).LinkIPWithCurrentHost), arg0, arg1)
}

// RemoveIP mocks base method.
func (m *MockStorage) RemoveIP(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIP", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIP indicates an expected call of RemoveIP.
func (mr *MockStorageMockRecorder) RemoveIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIP", reflect.TypeOf((*MockStorage)(nil).RemoveIP), arg0, arg1)
}

// SaveHost mocks base method.
func (m *MockStorage) SaveHost(arg0 context.Context, arg1 models.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveHost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveHost indicates an expected call of SaveHost.
func (mr *MockStorageMockRecorder) SaveHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveHost", reflect.TypeOf((*MockStorage)(nil).SaveHost), arg0, arg1)
}

// UnlinkIPFromCurrentHost mocks base method.
func (m *MockStorage) UnlinkIPFromCurrentHost(arg0 context.Context, arg1 models.IP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlinkIPFromCurrentHost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlinkIPFromCurrentHost indicates an expected call of UnlinkIPFromCurrentHost.
func (mr *MockStorageMockRecorder) UnlinkIPFromCurrentHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlinkIPFromCurrentHost", reflect.TypeOf((*MockStorage)(nil).UnlinkIPFromCurrentHost), arg0, arg1)
}

// UpdateIP mocks base method.
func (m *MockStorage) UpdateIP(arg0 context.Context, arg1 models.IP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIP", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIP indicates an expected call of UpdateIP.
func (mr *MockStorageMockRecorder) UpdateIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIP", reflect.TypeOf((*MockStorage)(nil).UpdateIP), arg0, arg1)
}
