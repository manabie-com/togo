package task

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

type TaskUseCase interface {
	ListTasks(context.Context, uint, sql.NullString) ([]*storages.Task, error)
	AddTask(context.Context, *storages.Task) error
	GetUserByUsername(ctx context.Context, username sql.NullString) (*storages.User, error)
	IsMaximumTasks(ctx context.Context, userID uint, createdDate sql.NullString, maxTodo uint) (bool, error)
}

type taskUseCase struct {
	storeRepository storages.Repository
}

func NewTaskUseCase(storeRepository storages.Repository) TaskUseCase {
	return &taskUseCase{storeRepository: storeRepository}
}

func (t *taskUseCase) ListTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*storages.Task, error) {
	return t.storeRepository.RetrieveTasks(
		ctx,
		userID,
		createdDate,
	)
}

func (t *taskUseCase) GetUserByUsername(ctx context.Context, username sql.NullString) (*storages.User, error) {
	return t.storeRepository.GetUserByUsername(ctx, username)
}

func (t *taskUseCase) AddTask(ctx context.Context, task *storages.Task) error {
	return t.storeRepository.AddTask(ctx, task)
}

func (t *taskUseCase) IsMaximumTasks(ctx context.Context, userID uint, createdDate sql.NullString, maxTodo uint) (bool, error) {
	tasks, err := t.storeRepository.RetrieveTasks(
		ctx,
		userID,
		createdDate,
	)
	if err != nil {
		return false, err
	}
	return len(tasks) >= int(maxTodo), nil
}
