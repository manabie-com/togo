package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string	`json:"id"       gorm:"primary_key;column:id;not null"  mapstructure:"id"`
	Password string	`json:"password" gorm:"column:password;not null"        mapstructure:"password"`
	MaxTodo  int32 	`json:"max_todo" gorm:"column:max_todo;not null"        mapstructure:"max_todo"`
}

func(User) TableName() string{
	return "users"
}

func(User) BeforeCreate(tx *gorm.DB) error {
	return nil
}
