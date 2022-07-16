package models

import "time"

type Task struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	Content    string     `json:"content"`
	TargetDate time.Time  `json:"target_date"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdateAt   *time.Time `json:"update_at,omitempty"`
}
