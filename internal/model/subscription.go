package model

import "time"

// Subscription represents the user subscription model
type Subscription struct {
	UserID  int        `json:"user_id"`
	PlanID  int        `json:"plan_id"`
	StartAt time.Time  `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}
