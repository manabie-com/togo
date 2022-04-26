package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=3,max=32"`
	Description string `json:"description" validate:"required,min=3,max=32"`
	UserID      uint   `json:"user_id" validate:"required,number"`
	User        User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
