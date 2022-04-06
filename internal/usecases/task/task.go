package task

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository"
)

// TaskUsecase is the definition for collection of methods related to the `tasks` table use case
type TaskUsecase interface {
	AddTask(ctx context.Context, task *models.Task) error
	RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*models.Task, error)
	IsMaxTasksPerDay(ctx context.Context, userID, maxTaskPerDay uint, createdDate sql.NullString) (bool, error)
}

type taskUsecase struct {
	repository repository.DatabaseRepository
}

// NewTaskUsecase returns a TaskUsecase attached with methods related to the `tasks` table use case
func NewTaskUsecase(repository repository.DatabaseRepository) TaskUsecase {
	return &taskUsecase{repository}
}

// AddTask is a wrapper of repository.AddTask that interact directly with the connected database
func (u *taskUsecase) AddTask(ctx context.Context, task *models.Task) error {
	return u.repository.AddTask(ctx, task)
}

// RetrieveTasks is a wrapper of repository.RetrieveTasks that interact directly with the connected database
func (u *taskUsecase) RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*models.Task, error) {
	return u.repository.RetrieveTasks(ctx, userID, createdDate)
}

// IsMaxTasksPerDay returns true if the specified user has reached the threshold for the maximum number of tasks perday and vice versa
func (u *taskUsecase) IsMaxTasksPerDay(ctx context.Context, userID, maxTaskPerDay uint, createdDate sql.NullString) (bool, error) {
	tasks, err := u.repository.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return false, err
	}

	return len(tasks) >= int(maxTaskPerDay), nil
}
