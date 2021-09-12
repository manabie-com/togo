// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"

	storages "github.com/manabie-com/togo/internal/storages"
)

// TaskService is an autogenerated mock type for the TaskService type
type TaskService struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: ctx, t
func (_m *TaskService) AddTask(ctx context.Context, t *storages.Task) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *storages.Task) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveTasks provides a mock function with given fields: ctx, userID, createdDate
func (_m *TaskService) RetrieveTasks(ctx context.Context, userID sql.NullString, createdDate sql.NullString) ([]*storages.Task, error) {
	ret := _m.Called(ctx, userID, createdDate)

	var r0 []*storages.Task
	if rf, ok := ret.Get(0).(func(context.Context, sql.NullString, sql.NullString) []*storages.Task); ok {
		r0 = rf(ctx, userID, createdDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storages.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, sql.NullString, sql.NullString) error); ok {
		r1 = rf(ctx, userID, createdDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateUser provides a mock function with given fields: ctx, userID, pwd
func (_m *TaskService) ValidateUser(ctx context.Context, userID sql.NullString, pwd sql.NullString) (bool, error) {
	ret := _m.Called(ctx, userID, pwd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, sql.NullString, sql.NullString) bool); ok {
		r0 = rf(ctx, userID, pwd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, sql.NullString, sql.NullString) error); ok {
		r1 = rf(ctx, userID, pwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
