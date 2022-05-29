// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/limit/handler.go

// Package limit is a generated GoMock package.
package limit

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

// GetLimit mocks base method.
func (m *MockService) GetLimit(ctx context.Context, req *GetLimitReq) (*Limit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimit", ctx, req)
	ret0, _ := ret[0].(*Limit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLimit indicates an expected call of GetLimit.
func (mr *MockServiceMockRecorder) GetLimit(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimit", reflect.TypeOf((*MockService)(nil).GetLimit), ctx, req)
}
