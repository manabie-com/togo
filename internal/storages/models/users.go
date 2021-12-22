package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"uniqueIndex" validate:"required"`
	Password string `json:",omitempty" validate:"required"`
}

func (_ User) TableName() string {
	return "users"
}
