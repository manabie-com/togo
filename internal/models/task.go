package models

type Task struct {
	ID         string `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}
