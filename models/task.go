package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID      string `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	UserID  string `json:"user_id" gorm:"size:191"`
	User    User
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := uuid.Parse(task.ID); err != nil {
		task.ID = uuid.NewString()
	}
	return nil
}
