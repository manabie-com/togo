package models

import "time"

type TaskStatus int8

const (
	ActiveTaskStatus TaskStatus = iota + 1
	DeletedTaskStatus
)

type Task struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	Content   string     `json:"content"`
	UserID    string     `json:"user_id"`
	Status    TaskStatus `json:"status"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (Task) TableName() string {
	return "tasks"
}
