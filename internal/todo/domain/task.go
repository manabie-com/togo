package domain

import (
	"context"
	"errors"
	"time"
)

// Task reflects tasks in DB
type Task struct {
	ID        int        `json:"id"`
	Content   string     `json:"content"`
	UserID    int        `json:"user_id" db:"user_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type TaskCreateParam struct {
	Content string
}

type TaskRepository interface {
	CreateTaskForUser(context.Context, int, TaskCreateParam) (*Task, error)
	GetTasksForUser(context.Context, int, string) ([]*Task, error)
}

var ErrTaskLimitReached = errors.New("Task Limit Reached")
