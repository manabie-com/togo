package model

import "time"

// Subscription represents the user subscription model
type Subscription struct {
	UserID  int
	PlanID  int
	StartAt time.Time
	EndAt   time.Time
}
