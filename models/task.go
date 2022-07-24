package models

import "time"

type Task struct {
	BaseModelID
	UserID      int       `json:"user_id" gorm:"index; not null"`
	Description string    `json:"description" gorm:"not null"`
	EndedAt     time.Time `json:"ended_at" gorm:"not null"`
	BaseModelTime
}

func (Task) TableName() string {
	return "public.tasks"
}
