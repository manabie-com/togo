// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	storages "github.com/perfectbuii/togo/internal/storages"
	mock "github.com/stretchr/testify/mock"
)

// TaskStore is an autogenerated mock type for the TaskStore type
type MockTaskStore struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: ctx, task
func (_m *MockTaskStore) AddTask(ctx context.Context, task *storages.Task) error {
	ret := _m.Called(ctx, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *storages.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTasks provides a mock function with given fields: ctx, task
func (_m *MockTaskStore) GetTasks(ctx context.Context, task *storages.Task) ([]*storages.Task, error) {
	ret := _m.Called(ctx, task)

	var r0 []*storages.Task
	if rf, ok := ret.Get(0).(func(context.Context, *storages.Task) []*storages.Task); ok {
		r0 = rf(ctx, task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storages.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *storages.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
