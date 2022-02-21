package model

import (
	"github.com/google/uuid"
	"time"
)

var (
	StatusNew    = "TODO"
	PrioryNormal = "NORMAL"
)

type Task struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Status    string    `db:"status" json:"status"`
	Priority  string    `db:"priority" json:"priority"`
	UserID    uuid.UUID `db:"user_id" json:"userID"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type TaskRequest struct {
	Title string `json:"title"`
}
