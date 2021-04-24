package models

type Task struct {
	ID          uint64 `json:"id"`
	Content     string `json:"content"`
	UserID      uint64 `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
