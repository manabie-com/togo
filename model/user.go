package model

import "time"

type User struct {
	ID        int
	Password  string
	MaxTodo   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
