// Code generated by MockGen. DO NOT EDIT.
// Source: go.etcd.io/etcd/client/v3 (interfaces: KV)

// Package etcdmock is a generated GoMock package.
package etcdmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v3 "go.etcd.io/etcd/client/v3"
)

// MockKV is a mock of KV interface.
type MockKV struct {
	ctrl     *gomock.Controller
	recorder *MockKVMockRecorder
}

// MockKVMockRecorder is the mock recorder for MockKV.
type MockKVMockRecorder struct {
	mock *MockKV
}

// NewMockKV creates a new mock instance.
func NewMockKV(ctrl *gomock.Controller) *MockKV {
	mock := &MockKV{ctrl: ctrl}
	mock.recorder = &MockKVMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKV) EXPECT() *MockKVMockRecorder {
	return m.recorder
}

// Compact mocks base method.
func (m *MockKV) Compact(arg0 context.Context, arg1 int64, arg2 ...v3.CompactOption) (*v3.CompactResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Compact", varargs...)
	ret0, _ := ret[0].(*v3.CompactResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compact indicates an expected call of Compact.
func (mr *MockKVMockRecorder) Compact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compact", reflect.TypeOf((*MockKV)(nil).Compact), varargs...)
}

// Delete mocks base method.
func (m *MockKV) Delete(arg0 context.Context, arg1 string, arg2 ...v3.OpOption) (*v3.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*v3.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockKVMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKV)(nil).Delete), varargs...)
}

// Do mocks base method.
func (m *MockKV) Do(arg0 context.Context, arg1 v3.Op) (v3.OpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(v3.OpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockKVMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockKV)(nil).Do), arg0, arg1)
}

// Get mocks base method.
func (m *MockKV) Get(arg0 context.Context, arg1 string, arg2 ...v3.OpOption) (*v3.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v3.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKVMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKV)(nil).Get), varargs...)
}

// Put mocks base method.
func (m *MockKV) Put(arg0 context.Context, arg1, arg2 string, arg3 ...v3.OpOption) (*v3.PutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*v3.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockKVMockRecorder) Put(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockKV)(nil).Put), varargs...)
}

// Txn mocks base method.
func (m *MockKV) Txn(arg0 context.Context) v3.Txn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Txn", arg0)
	ret0, _ := ret[0].(v3.Txn)
	return ret0
}

// Txn indicates an expected call of Txn.
func (mr *MockKVMockRecorder) Txn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Txn", reflect.TypeOf((*MockKV)(nil).Txn), arg0)
}
