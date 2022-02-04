package tasks

import (
	"log"

	"github.com/kozloz/togo"
	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/users"
)

type TaskStore interface {
	Create(userID int64, task string) (*togo.Task, error)
}
type Operation struct {
	store         TaskStore
	userOperation *users.Operation
}

func NewOperation(store TaskStore, userOp *users.Operation) *Operation {
	return &Operation{
		store:         store,
		userOperation: userOp,
	}
}

// Create the task for the user
func (o *Operation) Create(userID int64, taskName string) (*togo.Task, error) {
	log.Printf("Creating task '%s' for user '%d'.", taskName, userID)
	// Get user object
	user, err := o.userOperation.Get(userID)
	if err != nil {
		return nil, err
	}

	// Check if can still create tasks
	if user.DailyCount > user.DailyLimit {
		log.Printf("Max daily limit reached")
		return nil, errors.MaxLimit
	}

	// Insert task
	task, err := o.store.Create(user.ID, taskName)
	if err != nil {
		return nil, err
	}
	return task, nil
}
