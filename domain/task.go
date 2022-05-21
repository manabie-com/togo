package domain

import (
	"time"
)

type TaskParams struct {
	Content      string `json:"content" validate:"required"`
	UserEmail    string `json:"user_email"`
	UserMaxTasks int32  `json:"user_max_tasks"`
}

type Task struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" validate:"required"`
	UserId    *int      `json:"user_id"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type ITaskService interface {
	Create(params TaskParams) (Task, error)
}

type ITaskRepository interface {
	Create(task Task) (Task, error)
}
