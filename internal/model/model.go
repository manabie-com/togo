package model

import "time"

type LoginCredential struct {
	UserName string
	Password string
}
type AccessToken struct {
	Token string `json:"access_token"`
}

type TaskCreationRequest struct {
	Content string `json:"content"`
}

type Task struct {
	TaskID      string    `json:"task_id"`
	Content     string    `json:"content"`
	UserID      string    `json:"user_id"`
	CreatedDate time.Time `json:"created_date"`
}
