package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint64 `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
}
