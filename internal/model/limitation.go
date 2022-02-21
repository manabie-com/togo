package model

import (
	"time"
)

type Limitation struct {
	ID        uint64    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"userID"`
	LimitTask int64     `db:"limit_tasks" json:"limitTask"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
