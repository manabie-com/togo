package storages

import (
	"context"
	"database/sql"
)

// DB stores tasks and credentials of users
type DB interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
