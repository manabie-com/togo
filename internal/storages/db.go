package storages

import (
	"context"
	"database/sql"
)

// Interface for accessing database
type Repository interface {
	// RetrieveTasks returns tasks if match userID AND createDate.
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	// AddTask adds a new task to DB
	AddTask(ctx context.Context, t *Task) error
	// ValidateUser returns tasks if match userID AND password
	ValidateUser(ctx context.Context, username, pwd sql.NullString) (*User, error)
	// GetUserById returns user by userId
	GetUserById(ctx context.Context, userID sql.NullString) (*User, error)
	GetUserByUsername(ctx context.Context, username sql.NullString) (*User, error)
}
