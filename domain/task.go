package domain

import (
	"gorm.io/gorm"
	"time"
)

type TaskParams struct {
	Content   string `json:"content" validate:"required"`
	UserEmail string `json:"user_email" validate:"omitempty,required_without=user_id"`
	UserId    int    `json:"user_id" validate:"omitempty,required_without=user_email"`
	TaskLimit int    `json:"task_limit"`
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
	SetTx(tx *gorm.DB) *gorm.DB
	Create(task Task) (Task, error)
	Save(task Task) error
}
