package domain

import (
	"context"
	"time"
)

// Timestamp represents the timestamp of a record
type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents the user model
type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	DailyLimit int    `json:"daily_limit"`
	Timestamp
}

// Task represents the Task model
type Task struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
	Timestamp
}

// TaskRepository abstracts the database to the repository interface
type TaskRepository interface {
	Close()
	CheckTaskByUserID(ctx context.Context, userID int64) error
	CreateTask(ctx context.Context, userID int64, content string) (*Task, error)
	GetAllTaskByUserID(ctx context.Context, userID int64) ([]*Task, error)
}

// TaskService abstracts the task service
type TaskService interface {
	CreateTask(ctx context.Context, userID int64, content string) (*Task, error)
	GetAllTaskByUserID(ctx context.Context, userID int64) ([]*Task, error)
}
