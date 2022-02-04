package tasks

import (
	"github.com/kozloz/togo"
	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/users"
)

type TaskStore interface {
	Create(task string) (*togo.Task, error)
}
type Operation struct {
	store         TaskStore
	userOperation users.Operation
}

// Create the task for the user
func (o *Operation) Create(userID int64, taskName string) (*togo.Task, error) {
	// Get user object
	user, err := o.userOperation.Get(userID)
	if err != nil {
		return nil, err
	}

	// Check if can still create tasks
	if user.DailyCount > user.DailyLimit {
		return nil, errors.MaxLimit
	}

	// Insert task
	task, err := o.store.Create(taskName)
	if err != nil {
		return nil, err
	}
	return task, nil
}
