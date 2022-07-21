// Code generated by MockGen. DO NOT EDIT.
// Source: ./querier.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/tonghia/togo/internal/store"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetTaskByID mocks base method.
func (m *MockQuerier) GetTaskByID(ctx context.Context, id uint64) (store.TodoTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, id)
	ret0, _ := ret[0].(store.TodoTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockQuerierMockRecorder) GetTaskByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockQuerier)(nil).GetTaskByID), ctx, id)
}

// GetTaskByUserID mocks base method.
func (m *MockQuerier) GetTaskByUserID(ctx context.Context, userID uint64) ([]store.TodoTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByUserID", ctx, userID)
	ret0, _ := ret[0].([]store.TodoTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByUserID indicates an expected call of GetTaskByUserID.
func (mr *MockQuerierMockRecorder) GetTaskByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByUserID", reflect.TypeOf((*MockQuerier)(nil).GetTaskByUserID), ctx, userID)
}

// GetTotalTaskByUserID mocks base method.
func (m *MockQuerier) GetTotalTaskByUserID(ctx context.Context, userID uint64) (store.GetTotalTaskByUserIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalTaskByUserID", ctx, userID)
	ret0, _ := ret[0].(store.GetTotalTaskByUserIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalTaskByUserID indicates an expected call of GetTotalTaskByUserID.
func (mr *MockQuerierMockRecorder) GetTotalTaskByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalTaskByUserID", reflect.TypeOf((*MockQuerier)(nil).GetTotalTaskByUserID), ctx, userID)
}

// InsertTask mocks base method.
func (m *MockQuerier) InsertTask(ctx context.Context, arg store.InsertTaskParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTask", ctx, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTask indicates an expected call of InsertTask.
func (mr *MockQuerierMockRecorder) InsertTask(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTask", reflect.TypeOf((*MockQuerier)(nil).InsertTask), ctx, arg)
}
