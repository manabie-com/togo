package model

import (
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UserID      string    `json:"user_id"`
	CreatedTime time.Time `json:"created_time"`
}
