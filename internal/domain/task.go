package domain

import (
	"context"
	"errors"
)

var (
	// ErrTaskLimitExceed task limit exceed error
	ErrTaskLimitExceed = errors.New("TASK_LIMIT_EXCEED")
	// ErrTaskNotFound task not found error
	ErrTaskNotFound = errors.New("TASK_NOT_FOUND")
)

// Task model
type Task struct {
	ID      uint   `json:"id,omitempty" gorm:"primarykey"`
	UserID  uint   `json:"userId,omitempty"`
	Content string `json:"content,omitempty"`
	User    User   `json:"-"`
}

// TaskService service interface
type TaskService interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	Update(ctx context.Context, filter, update *Task) (*Task, error)
	FindByUserID(ctx context.Context, userID uint) ([]*Task, error)
}
