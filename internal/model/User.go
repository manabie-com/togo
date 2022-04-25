package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
	MaxTodos int    `json:"max_todos"`
}
