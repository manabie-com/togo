package models

import "time"

// Task reflects tasks in DB
type Task struct {
	ID          string    `json:"id"`
	Content     string    `json:"content"`
	UserID      string    `json:"userId"`
	CreatedDate time.Time `json:"createdDate"`
}
