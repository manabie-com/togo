package task

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

type TaskUseCase interface {
	ListTasks(context.Context, uint, sql.NullString) ([]*storages.Task, error)
	AddTask(context.Context, *storages.Task) error
	IsMaximumTasks(ctx context.Context, userID uint, createdDate sql.NullString, maxTodo uint) (bool, error)
}

type taskUseCase struct {
	storeRepository storages.Repository
}

func NewTaskUseCase(storeRepository storages.Repository) TaskUseCase {
	return &taskUseCase{storeRepository: storeRepository}
}

// ListTasks returns tasks by userID AND createDate.
func (t *taskUseCase) ListTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*storages.Task, error) {
	return t.storeRepository.RetrieveTasks(
		ctx,
		userID,
		createdDate,
	)
}

// AddTask adds a new task to DB
func (t *taskUseCase) AddTask(ctx context.Context, task *storages.Task) error {
	return t.storeRepository.AddTask(ctx, task)
}

// IsMaximumTasks check if number of user tasks reached the limited
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
