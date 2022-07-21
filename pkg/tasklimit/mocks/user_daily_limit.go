// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_daily_limit.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserLimitSvc is a mock of UserLimitSvc interface.
type MockUserLimitSvc struct {
	ctrl     *gomock.Controller
	recorder *MockUserLimitSvcMockRecorder
}

// MockUserLimitSvcMockRecorder is the mock recorder for MockUserLimitSvc.
type MockUserLimitSvcMockRecorder struct {
	mock *MockUserLimitSvc
}

// NewMockUserLimitSvc creates a new mock instance.
func NewMockUserLimitSvc(ctrl *gomock.Controller) *MockUserLimitSvc {
	mock := &MockUserLimitSvc{ctrl: ctrl}
	mock.recorder = &MockUserLimitSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserLimitSvc) EXPECT() *MockUserLimitSvcMockRecorder {
	return m.recorder
}

// GetUserLimit mocks base method.
func (m *MockUserLimitSvc) GetUserLimit(userID uint64) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLimit", userID)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetUserLimit indicates an expected call of GetUserLimit.
func (mr *MockUserLimitSvcMockRecorder) GetUserLimit(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLimit", reflect.TypeOf((*MockUserLimitSvc)(nil).GetUserLimit), userID)
}
