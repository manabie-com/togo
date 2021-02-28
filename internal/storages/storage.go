package storages

import (
	"context"
	"database/sql"
)

type Storage interface {
	RetrieveTasks(ctx context.Context, userID sql.NullString, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	ValidateUser(ctx context.Context, userID sql.NullString, pwd sql.NullString) bool
}
