package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" `
	Password string `json:"password" `
	MaxTodos int    `json:"max_todos" `
}
