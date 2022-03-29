package models

import (
	"time"
	"togo/globals/database"
)

type Task struct {
	ID				uint	`gorm:"primaryKey" json:"id"`
	UserID 			uint	`json:"user_id"`
	Detail			string	`json:"detail"`
	CreatedAt		*time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt		*time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func CreateTask(task Task) (*Task, error){
	err := database.SQL.Create(&task).Error
	if err != nil {
		return &Task{}, err
	}

	return &task, nil
}