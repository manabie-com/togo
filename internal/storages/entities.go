package storages

import "time"

// Task reflects tasks in DB
type Task struct {
	ID          int64     `json:"id"`
	Content     string    `json:"content"`
	UserID      string    `json:"user_id"`
	CreatedDate time.Time `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}
