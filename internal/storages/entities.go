package storages

import (
	"context"
	"database/sql"
)

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	Status      string `json:"status"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}

// DBConnecter ...
type DBConnecter interface {
	Connect() (*sql.DB, error)
}

// IToGoDB interface for database, database management systems should implement all the method bellow
type IToGoDB interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	UpdateStatusTask(ctx context.Context, userID, taskID, status sql.NullString) error
	UpdateAllStatusTasks(ctx context.Context, userID, createdDate, status sql.NullString) error
	DeleteTask(ctx context.Context, userID, taskID sql.NullString) error
	DeleteTasks(ctx context.Context, userID sql.NullString, createdDate sql.NullString) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
	GetUserMaxTask(ctx context.Context, userID sql.NullString) (int, error)
	GetUserTodayTask(ctx context.Context, userID sql.NullString) (int, error)
}
