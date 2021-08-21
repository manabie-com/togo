package storages

import (
	"time"
)

// Task reflects tasks in DB
type Task struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// User reflects users data from DB
type User struct {
	ID        string    `json:"id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Login struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type CreateTask struct {
	Content string `json:"content"`
}

func GetTask() Task {
	var task Task
	return task
}

func GetUser() User {
	var user User
	return user
}

func GetLogin() Login {
	var login Login
	return login
}

func GetCreateTask() CreateTask {
	var createTask CreateTask
	return createTask
}
