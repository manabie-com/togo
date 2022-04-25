package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
