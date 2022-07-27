package dto

import "time"

type UserResponse struct {
	ID         int    ` json:"id"`
	Name       string `json:"name"`
	LimitCount int    `json:"limit_count"`
}

type TaskResponse struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	EndedAt     time.Time `json:"ended_at"`
}
