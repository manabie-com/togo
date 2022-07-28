package dto

import "time"

type CreateTaskDto struct {
	Description string    `json:"description" validate:"required"`
	EndedAt     time.Time `json:"ended_at" validate:"required"`
}
