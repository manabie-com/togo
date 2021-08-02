package models

import (
	"time"

	"github.com/google/uuid"
)

// Task reflects tasks in DB
type Task struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"size:255; not null" json:"content"`
	AccountID uuid.UUID `gorm:"column:account_id" json:"account_id"`
	//default timestamps
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`
}
