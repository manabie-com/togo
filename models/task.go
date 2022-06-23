package models

import "time"

type Task struct {
	ID        int32     `json:"id" validate:"omitempty"`
	Name      string    `json:"name" validate:"required"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UserId    uint32    `json:"userId"`
}
