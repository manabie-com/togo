// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/user/handler.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockService) Create(ctx context.Context, req *CreateUserReq) (*UserSafe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(*UserSafe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, req)
}

// Delete mocks base method.
func (m *MockService) Delete(ctx context.Context, req *DeleteUserByNameReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), ctx, req)
}

// GetByUserName mocks base method.
func (m *MockService) GetByUserName(ctx context.Context, req *GetUserByUserNameReq) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserName", ctx, req)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserName indicates an expected call of GetByUserName.
func (mr *MockServiceMockRecorder) GetByUserName(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserName", reflect.TypeOf((*MockService)(nil).GetByUserName), ctx, req)
}

// List mocks base method.
func (m *MockService) List(ctx context.Context, req *ListUsersReq) ([]*UserSafe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, req)
	ret0, _ := ret[0].([]*UserSafe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceMockRecorder) List(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), ctx, req)
}

// UpdateUserTier mocks base method.
func (m *MockService) UpdateUserTier(ctx context.Context, req *UpdateUserTierReq) (*UserSafe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserTier", ctx, req)
	ret0, _ := ret[0].(*UserSafe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserTier indicates an expected call of UpdateUserTier.
func (mr *MockServiceMockRecorder) UpdateUserTier(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserTier", reflect.TypeOf((*MockService)(nil).UpdateUserTier), ctx, req)
}
