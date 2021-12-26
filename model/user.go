package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int
	Username string
	Password string
	MaxTodo  int
}
