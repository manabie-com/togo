package usecase

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/utils"
)

type taskUsecase struct {
	TaskDB         postgres.TaskDB
	contextTimeout time.Duration
}

type TaskUsecase interface {
	ListTasks(ctx context.Context, userID string, createdDate time.Time) ([]storages.Task, error)
	AddTask(ctx context.Context, task *storages.Task) error
	ValidateUser(ctx context.Context, userID string, password string) (bool, error)
}

func NewTaskUsecase(taskDB postgres.TaskDB, timeout time.Duration) TaskUsecase {
	return &taskUsecase{
		TaskDB:         taskDB,
		contextTimeout: timeout,
	}
}

func (t *taskUsecase) ListTasks(ctx context.Context, userID string, createdDate time.Time) (res []storages.Task, err error) {
	ctx, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	res, err = t.TaskDB.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return
	}
	return
}
func (t *taskUsecase) AddTask(ctx context.Context, task *storages.Task) (err error) {
	ctx, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()
	userID, _ := utils.GetUserIDFromCtx(ctx)
	numTasks, err := t.TaskDB.CountTaskPerDay(ctx, userID, time.Now())
	if err != nil {
		return err
	}
	if numTasks > 5 {
		return domain.ErrTooMany
	}
	err = t.TaskDB.AddTask(ctx, task)
	return
}
func (t *taskUsecase) ValidateUser(ctx context.Context, userID string, password string) (res bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()
	res, err = t.TaskDB.ValidateUser(ctx, userID, password)
	return
}
