package entity

import "time"

type Task struct {
	ID          int32     `json:"id"`
	Content     string    `json:"content"`
	UserID      int32     `json:"user_id"`
	CreatedDate time.Time `json:"created_date"`
	IsDone      bool      `json:"is_done"`
}
