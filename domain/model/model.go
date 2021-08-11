package model

import "time"

type User struct {
	Id       string
	Username string
	Password string
	MaxTodo  int
}

type Task struct {
	ID          string
	Content     string
	UserID      string
	CreatedDate time.Time
}
