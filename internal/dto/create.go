package dto

import "time"

type CreateTaskDto struct {
	Description string    `json:"description"`
	EndedAt     time.Time `json:"ended_at"`
}

type CreateUserDto struct {
	Name       string `json:"name"`
	LimitCount int    `json:"limit_count"`
}
