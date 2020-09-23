package usecase

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal"
	"github.com/manabie-com/togo/internal/storages"
)

type todoUsecase struct {
	todoRepository internal.Repository
}

func NewTodoUsecase(t internal.Repository) internal.Usecase {
	return &todoUsecase{
		todoRepository: t,
	}
}

func (s *todoUsecase) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return s.todoRepository.ValidateUser(ctx, userID, pwd)
}

func (s *todoUsecase) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	tasks, err := s.todoRepository.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *todoUsecase) AddTask(ctx context.Context, t *storages.Task) error {
	user, err := s.todoRepository.FindUserByID(ctx, t.UserID)
	if err != nil {
		return err
	}
	tasks, err := s.RetrieveTasks(ctx, sql.NullString{t.UserID, true}, sql.NullString{t.CreatedDate, true})
	if err != nil {
		return err
	}
	if len(tasks) >= user.MaxTodo {
		return ErrMaxTaskReached
	}
	err = s.todoRepository.AddTask(ctx, t)
	if err != nil {
		return err
	}
	return nil
}
