package task

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository"
)

type TaskUsecase interface {
	AddTask(ctx context.Context, task *models.Task) error
	RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*models.Task, error)
	IsMaxTasksPerDay(ctx context.Context, userID, maxTaskPerDay uint, createdDate sql.NullString) (bool, error)
}

type taskUsecase struct {
	repository repository.DatabaseRepository
}

func NewTaskUsecase(repository repository.DatabaseRepository) TaskUsecase {
	return &taskUsecase{repository}
}

func (u *taskUsecase) AddTask(ctx context.Context, task *models.Task) error {
	return u.repository.AddTask(ctx, task)
}

func (u *taskUsecase) RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*models.Task, error) {
	return u.repository.RetrieveTasks(ctx, userID, createdDate)
}

func (u *taskUsecase) IsMaxTasksPerDay(ctx context.Context, userID, maxTaskPerDay uint, createdDate sql.NullString) (bool, error) {
	tasks, err := u.repository.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return false, err
	}

	return len(tasks) >= int(maxTaskPerDay), nil
}
