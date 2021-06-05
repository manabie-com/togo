package storages

import (
	"context"
	"database/sql"
)

type TaskWriteRepository interface {
	AddTask(ctx context.Context, t *Task) error
}

type TaskListRepository interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
}

type TaskWriteListRepository interface {
	TaskListRepository
	TaskWriteRepository
}

type UserReadRepository interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
