package dao

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserID  int64 `gorm:"foreignKey:user.ID"`
	Name    string
	Content string
}
