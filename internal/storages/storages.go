package storages

import (
	"context"
	"database/sql"
)

type Repository interface {
	RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	GetUserByUsername(ctx context.Context, username sql.NullString) (*User, error)
	ValidateUser(ctx context.Context, username, pwd sql.NullString) bool
}
