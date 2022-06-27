package domain

import (
	"context"

	"manabie/togo/internal/model"

	"github.com/stretchr/testify/mock"
)

type mockTaskStore struct {
	mock.Mock
}

func (s *mockTaskStore) RetrieveTasks(ctx context.Context, task *model.Task) ([]*model.Task, error) {
	args := s.Called(task)
	if t := args.Get(0); t != nil {
		return t.([]*model.Task), nil
	}
	return nil, args.Error(1)
}

func (s *mockTaskStore) AddTask(ctx context.Context, task *model.Task) error {
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
func (s *mockTaskCountStore) Inc(ctx context.Context, userID, date string) (int, error) {
	args := s.Called(userID, date)
	return args.Int(0), args.Error(1)
}

func (s *mockTaskCountStore) Desc(ctx context.Context, userID, date string) error {
	args := s.Called(userID, date)
	return args.Error(0)
}

type mockUserStore struct {
	mock.Mock
}

func (s *mockUserStore) Create(ctx context.Context, user *model.User) error {
	args := s.Called(user)
	return args.Error(0)
}
func (s *mockUserStore) FindUser(ctx context.Context, userID string) (*model.User, error) {
	args := s.Called(userID)
	if u := args.Get(0); u != nil {
		return u.(*model.User), nil
	}
	return nil, args.Error(1)
}
