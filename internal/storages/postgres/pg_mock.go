package postgres

import (
	"context"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/mock"
	"time"
)

type DatabaseMock struct {
	mock.Mock
}

func (m *DatabaseMock) ValidateUser(ctx context.Context, username, password string) (*storages.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(*storages.User), args.Error(1)
}

func (m *DatabaseMock) GetTasks(ctx context.Context, usrId int, createAt time.Time) ([]*storages.Task, error) {
	args := m.Called(ctx, usrId, createAt)
	return args.Get(0).([]*storages.Task), args.Error(1)
}

func (m *DatabaseMock) InsertTask(ctx context.Context, task *storages.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}
