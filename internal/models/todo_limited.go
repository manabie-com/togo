package models

import "gorm.io/gorm"

type ToDoConfig struct {
	gorm.Model
	UserID  string `gorm:"type:VARCHAR(64)"`
	Limited int64  `gorm:"type:INT"`
}

func (ToDoConfig) TableName() string {
	return "todo_config"
}
