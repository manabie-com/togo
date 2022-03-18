package model

import "time"

type Task struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
