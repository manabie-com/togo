// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// IsUserExisted provides a mock function with given fields: userID
func (_m *UserRepository) IsUserExisted(userID int64) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsUserHavingMaxTodo provides a mock function with given fields: userID, date
func (_m *UserRepository) IsUserHavingMaxTodo(userID int64, date time.Time) error {
	ret := _m.Called(userID, date)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, time.Time) error); ok {
		r0 = rf(userID, date)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
