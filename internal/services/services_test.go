package services

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
	"github.com/manabie-com/togo/internal/storages"

)

type storeMock struct {
	mock.Mock
}

func (s *storeMock) RetrieveTasks(ctx context.Context, userID sql.NullString, createdDate sql.NullString) ([]*storages.Task, error) {
	args := s.Called(ctx, userID, createdDate)
	return args.Get(0).([]*storages.Task), args.Error(1)
}

func (s *storeMock) AddTask(ctx context.Context, t *storages.Task) error {
	args := s.Called(ctx, t)
	return args.Error(0)
}

func (s *storeMock) ValidateUser(ctx context.Context, userID sql.NullString, pwd sql.NullString) bool {
	args := s.Called(ctx, userID, pwd)
	return args.Bool(0)
}