package models

import (
	"time"
)

type Task struct {
	TaskID      string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy   string    `json:"user_id,omitempty" bson:"user_ids,omitempty"`
}
