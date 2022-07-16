package models

import "time"

type Setting struct {
	ID        int        `json:"id"`
	MemberID  int        `json:"member_id"`
	LimitTask int        `json:"limit_task"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at,omitempty"`
}
