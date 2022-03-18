package model

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	TaskLimit int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
