package entities

import "time"

// users table
type User struct {
	ID        int
	Email     string
	Password  string
	LimitTodo int
	CreatedAt time.Time
	UpdatedAt time.Time
}
