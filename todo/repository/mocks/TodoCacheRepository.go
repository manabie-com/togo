// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/triet-truong/todo/todo/model"
)

// TodoCacheRepository is an autogenerated mock type for the TodoCacheRepository type
type TodoCacheRepository struct {
	mock.Mock
}

// GetCachedUser provides a mock function with given fields: id
func (_m *TodoCacheRepository) GetCachedUser(id uint) (model.UserRedisModel, error) {
	ret := _m.Called(id)

	var r0 model.UserRedisModel
	if rf, ok := ret.Get(0).(func(uint) model.UserRedisModel); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.UserRedisModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetUser provides a mock function with given fields: user
func (_m *TodoCacheRepository) SetUser(user model.UserRedisModel) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.UserRedisModel) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
