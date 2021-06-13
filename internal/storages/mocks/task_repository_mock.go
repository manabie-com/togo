package mocks

import (
	"context"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/stretchr/testify/mock"
	"time"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) CreateTask(ctx context.Context, content string, owner *ent.User) (*ent.Task, error) {
	ret := m.Called(ctx, content, owner)
	return ret.Get(0).(*ent.Task), ret.Error(1)
}

func (m *TaskRepositoryMock) GetTaskByDate(ctx context.Context, userId string, gte time.Time, lt time.Time) ([]*ent.Task, error) {
	ret := m.Called(ctx, userId, gte, lt)
	return ret.Get(0).([]*ent.Task), ret.Error(1)
}
