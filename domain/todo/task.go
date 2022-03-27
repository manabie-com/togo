package todo

import (
	"context"
	"fmt"
	"time"

	"github.com/laghodessa/togo/domain"
)

// Task represents a todo task which users can add
type Task struct {
	ID      string
	UserID  string
	Message string
}

type TaskOpt func(*Task) error

func NewTask(opts ...TaskOpt) (task Task, err error) {
	for _, applyOpt := range opts {
		if err := applyOpt(&task); err != nil {
			return Task{}, err
		}
	}

	if task.UserID == "" {
		return Task{}, fmt.Errorf("%w: user id can't be blank", domain.ErrInvalidArg)
	}
	if task.Message == "" {
		return Task{}, fmt.Errorf("%w: task message can't be blank", domain.ErrInvalidArg)
	}
	task.ID = domain.NewID()
	return task, nil
}

func TaskUserID(userID string) TaskOpt {
	return func(t *Task) error {
		t.UserID = userID
		return nil
	}
}

func TaskMessage(msg string) TaskOpt {
	return func(t *Task) error {
		t.Message = msg
		return nil
	}
}

type TaskRepo interface {
	// CountByUserID returns the number of tasks the user currently has in a time range
	//
	// since is inclusive
	//
	// until is exclusive
	CountInTimeRangeByUserID(_ context.Context, userID string, since time.Time, until time.Time) (int, error)
	// AddTask add new user tasks.
	// It should also handle user daily limit logic to avoid race condition with daily limit
	AddTask(_ context.Context, _ Task, loc *time.Location, dailyLimit int) error
}
