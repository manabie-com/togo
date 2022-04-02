package model

import "time"

// Subscription represents the user subscription model
type Subscription struct {
	UserID  int `gorm:"primaryKey;autoIncrement:false"`
	PlanID  int `gorm:"primaryKey;autoIncrement:false"`
	StartAt time.Time
	EndAt   time.Time
}
