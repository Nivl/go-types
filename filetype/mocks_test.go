// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: ReadSeeker)

// Package filetype_test is a generated GoMock package.
package filetype_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReadSeeker is a mock of ReadSeeker interface
type MockReadSeeker struct {
	ctrl     *gomock.Controller
	recorder *MockReadSeekerMockRecorder
}

// MockReadSeekerMockRecorder is the mock recorder for MockReadSeeker
type MockReadSeekerMockRecorder struct {
	mock *MockReadSeeker
}

// NewMockReadSeeker creates a new mock instance
func NewMockReadSeeker(ctrl *gomock.Controller) *MockReadSeeker {
	mock := &MockReadSeeker{ctrl: ctrl}
	mock.recorder = &MockReadSeekerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadSeeker) EXPECT() *MockReadSeekerMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReadSeeker) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReadSeekerMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadSeeker)(nil).Read), arg0)
}

// Seek mocks base method
func (m *MockReadSeeker) Seek(arg0 int64, arg1 int) (int64, error) {
	ret := m.ctrl.Call(m, "Seek", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seek indicates an expected call of Seek
func (mr *MockReadSeekerMockRecorder) Seek(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seek", reflect.TypeOf((*MockReadSeeker)(nil).Seek), arg0, arg1)
}