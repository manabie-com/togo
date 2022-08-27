package models

import "time"

type Tasks struct {
	ID          string    `json:"id"`
	AssignedTo  string    `json:"assigned_to"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Users struct {
	UserID     string    `json:"user_id"`
	Name       string    `json:"name"`
	TaskPerDay int       `json:"task_per_day"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
