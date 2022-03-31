package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	UserID  string `gorm:"type:VARCHAR(64)"`
	Content string `gorm:"type:TEXT"`
}

func (ToDo) TableName() string {
	return "todo"
}
