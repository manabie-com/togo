package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `gorm:"unique" json:"username" `
	Role     string `json:"role" `
	Password string
	Setting  Setting `gorm:"foreignKey:UserID;references:id" json:"setting"`
}
