package models

import (
	"time"
)

type BaseModelID struct {
	ID uint64 ` json:"id" gorm:"primaryKey;autoIncrement"`
}

type BaseModelTime struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`
}
