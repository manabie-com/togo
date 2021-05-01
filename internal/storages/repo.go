package storages

import (
	"context"
)

// Repository interface contains all methods to interact with storages db entities
type Repository interface {
	// RetrieveTasks returns tasks by userID AND createDate.
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*Task, error)
	// AddTask adds a new task to DB
	AddTask(ctx context.Context, t *Task) error
	// ValidateUser returns tasks by userID AND password
	ValidateUser(ctx context.Context, userID, pwd string) bool
	// LoadTasksCount return number of task count by userID and createdDate
	LoadTasksCount(ctx context.Context, userID, createdDate string) (cnt int, err error)
	// MaxTodo returns a user's max_todo (max daily tasks create requests) by userID
	MaxTodo(ctx context.Context, userID string) (int, error)
}
