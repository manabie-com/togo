package model

import (
	"context"
)

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	IsAllowedToAddTask(ctx context.Context, userId string) bool
}