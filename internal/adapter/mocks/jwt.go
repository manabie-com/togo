// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// JWTAdapter is an autogenerated mock type for the JWTAdapter type
type JWTAdapter struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: ctx, userID
func (_m *JWTAdapter) CreateToken(ctx context.Context, userID string) (string, error) {
	ret := _m.Called(ctx, userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: ctx, token
func (_m *JWTAdapter) VerifyToken(ctx context.Context, token string) (string, error) {
	ret := _m.Called(ctx, token)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
