// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "togo/internal/user/dto"

	mock "github.com/stretchr/testify/mock"

	response "togo/internal/response"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: createUserDto
func (_m *UserService) Create(createUserDto *dto.CreateUserDto) (*response.UserResponse, error) {
	ret := _m.Called(createUserDto)

	var r0 *response.UserResponse
	if rf, ok := ret.Get(0).(func(*dto.CreateUserDto) *response.UserResponse); ok {
		r0 = rf(createUserDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.CreateUserDto) error); ok {
		r1 = rf(createUserDto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}