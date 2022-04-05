package model

import "time"

type Task struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    int        `json:"user_id"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
}
