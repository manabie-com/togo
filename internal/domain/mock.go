package domain

import (
	"context"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/mock"
)

type mockTaskStore struct {
	mock.Mock
}

func (s *mockTaskStore) RetrieveTasks(ctx context.Context, task *storages.Task) ([]*storages.Task, error) {
	args := s.Called(task)
	if t := args.Get(0); t != nil {
		return t.([]*storages.Task), nil
	}
	return nil, args.Error(1)
}

func (s *mockTaskStore) AddTask(ctx context.Context, task *storages.Task) error {
	args := s.Called(task)
	return args.Error(0)
}

type mockTaskCountStore struct {
	mock.Mock
}

func (s *mockTaskCountStore) CreateIfNotExists(ctx context.Context, userID, date string) error {
	args := s.Called(userID, date)
	return args.Error(0)
}
func (s *mockTaskCountStore) UpdateAndGet(ctx context.Context, userID, date string) (int, error) {
	args := s.Called(userID, date)
	return args.Int(0), args.Error(1)
}

type mockUserStore struct {
	mock.Mock
}

func (s *mockUserStore) Create(ctx context.Context, user *storages.User) error {
	args := s.Called(user)
	return args.Error(0)
}
func (s *mockUserStore) FindUser(ctx context.Context, userID string) (*storages.User, error) {
	args := s.Called(userID)
	if u := args.Get(0); u != nil {
		return u.(*storages.User), nil
	}
	return nil, args.Error(1)
}
