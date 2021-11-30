package mock

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/mock"
)

type MockedTaskDB struct {
	mock.Mock
}

func (m *MockedTaskDB) RetrieveTasks(ctx context.Context, userID string, createdDate time.Time) ([]storages.Task, error) {
	args := m.Called(ctx, userID, createdDate)
	var res0 []storages.Task
	if rf, ok := args.Get(0).(func(context.Context, string, time.Time) []storages.Task); ok {
		res0 = rf(ctx, userID, createdDate)
	} else {
		if args.Get(0) != nil {
			res0 = args.Get(0).([]storages.Task)
		}
	}
	var res1 error
	if rf, ok := args.Get(1).(func(context.Context, string, time.Time) error); ok {
		res1 = rf(ctx, userID, createdDate)
	} else {
		res1 = args.Error(1)
	}
	return res0, res1
}

func (m *MockedTaskDB) AddTask(ctx context.Context, task *storages.Task) error {
	args := m.Called(ctx, task)
	var res0 error
	if rf, ok := args.Get(0).(func(context.Context, *storages.Task) error); ok {
		res0 = rf(ctx, task)
	} else {
		res0 = args.Error(0)
	}
	return res0
}

func (m *MockedTaskDB) CountTaskPerDay(ctx context.Context, userID string, createdDate time.Time) (uint8, error) {
	args := m.Called(ctx, userID, createdDate)
	var res0 uint8
	if rf, ok := args.Get(0).(func(context.Context, string, time.Time) uint8); ok {
		res0 = rf(ctx, userID, createdDate)
	}
	var res1 error
	if rf, ok := args.Get(1).(func(context.Context, string, time.Time) error); ok {
		res1 = rf(ctx, userID, createdDate)
	} else {
		res1 = args.Error(1)
	}
	return res0, res1
}

func (m *MockedTaskDB) ValidateUser(ctx context.Context, userID string, pw string) (bool, error) {
	args := m.Called(ctx, userID, pw)
	var res0 bool
	if rf, ok := args.Get(0).(func(context.Context, string, string) bool); ok {
		res0 = rf(ctx, userID, pw)
	}
	var res1 error
	if rf, ok := args.Get(1).(func(context.Context, string, string) error); ok {
		res1 = rf(ctx, userID, pw)
	} else {
		res1 = args.Error(1)
	}
	return res0, res1
}
