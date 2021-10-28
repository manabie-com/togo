package storages

import (
	"time"
)

// Task reflects tasks in DB
type Task struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// User reflects users data from DB
type User struct {
	ID        int
	Email     string
	Username  string
	Password  string
	MaxTodo   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
