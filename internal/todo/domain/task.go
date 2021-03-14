package domain

import (
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
	CreateTaskForUser(int, TaskCreateParam) (*Task, error)
	GetTasksForUser(int, string) ([]*Task, error)
	GetTaskCount(int, string) (int, error)
}

var ErrTaskLimitReached = errors.New("Task Limit Reached")
