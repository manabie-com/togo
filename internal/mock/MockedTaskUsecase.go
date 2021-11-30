package mock

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/mock"
)

type MockedTaskUsecase struct {
	mock.Mock
}

func (m *MockedTaskUsecase) ListTasks(ctx context.Context, userID string, createdDate time.Time) ([]storages.Task, error) {
	args := m.Called(ctx, userID, createdDate)
	var r0 []storages.Task
	if rf, ok := args.Get(0).(func(context.Context, string, time.Time) []storages.Task); ok {
		r0 = rf(ctx, userID, createdDate)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]storages.Task)
		}
	}
	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, userID, createdDate)
	} else {
		r1 = args.Error(1)
	}
	return r0, r1
}
func (m *MockedTaskUsecase) AddTask(ctx context.Context, task *storages.Task) error {
	args := m.Called(ctx, task)
	var r0 error
	if rf, ok := args.Get(0).(func(context.Context, *storages.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = args.Error(0)
	}
	return r0
}
func (m *MockedTaskUsecase) ValidateUser(ctx context.Context, userID string, password string) (bool, error) {
	args := m.Called(ctx, userID, password)
	var r0 bool
	if rf, ok := args.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, userID, password)
	} else {
		r0 = args.Bool(0)
	}
	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, password)
	} else {
		r1 = args.Error(1)
	}
	return r0, r1
}
func (m *MockedTaskUsecase) CountTaskPerDay(ctx context.Context, userID string, createdDate time.Time) (uint8, error) {
	args := m.Called(ctx, userID, createdDate)
	var r0 uint8
	if rf, ok := args.Get(0).(func(context.Context, string, time.Time) uint8); ok {
		r0 = rf(ctx, userID, createdDate)
	} else {
		r0 = args.Get(0).(uint8)
	}
	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, userID, createdDate)
	} else {
		r1 = args.Error(1)
	}
	return r0, r1
}
