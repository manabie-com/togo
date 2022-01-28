package model

import "time"

type User struct {
	ID        int
	Password  string
	LimitTask int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
