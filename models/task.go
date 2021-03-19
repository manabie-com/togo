package models

import "time"

type Task struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Content   string    `gorm:"type:text;not null;" json:"content"`
	Username  string    `gorm:"index;not null" json:"username"` // Foreign key (belongs to)
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
