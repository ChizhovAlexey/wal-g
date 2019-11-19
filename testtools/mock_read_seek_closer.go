// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ioextensions/io.go

package testtools

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReadSeekCloser is a mock of ReadSeekCloser interface
type MockReadSeekCloser struct {
	ctrl     *gomock.Controller
	recorder *MockReadSeekCloserMockRecorder
}

// MockReadSeekCloserMockRecorder is the mock recorder for MockReadSeekCloser
type MockReadSeekCloserMockRecorder struct {
	mock *MockReadSeekCloser
}

// NewMockReadSeekCloser creates a new mock instance
func NewMockReadSeekCloser(ctrl *gomock.Controller) *MockReadSeekCloser {
	mock := &MockReadSeekCloser{ctrl: ctrl}
	mock.recorder = &MockReadSeekCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadSeekCloser) EXPECT() *MockReadSeekCloserMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReadSeekCloser) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReadSeekCloserMockRecorder) Read(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadSeekCloser)(nil).Read), p)
}

// Seek mocks base method
func (m *MockReadSeekCloser) Seek(offset int64, whence int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seek", offset, whence)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seek indicates an expected call of Seek
func (mr *MockReadSeekCloserMockRecorder) Seek(offset, whence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seek", reflect.TypeOf((*MockReadSeekCloser)(nil).Seek), offset, whence)
}

// Close mocks base method
func (m *MockReadSeekCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockReadSeekCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadSeekCloser)(nil).Close))
}

// MockFlusher is a mock of Flusher interface
type MockFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockFlusherMockRecorder
}

// MockFlusherMockRecorder is the mock recorder for MockFlusher
type MockFlusherMockRecorder struct {
	mock *MockFlusher
}

// NewMockFlusher creates a new mock instance
func NewMockFlusher(ctrl *gomock.Controller) *MockFlusher {
	mock := &MockFlusher{ctrl: ctrl}
	mock.recorder = &MockFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlusher) EXPECT() *MockFlusherMockRecorder {
	return m.recorder
}

// Flush mocks base method
func (m *MockFlusher) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockFlusherMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockFlusher)(nil).Flush))
}