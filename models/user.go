package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := uuid.Parse(user.ID); err != nil {
		user.ID = uuid.New().String()
	}
	return nil
}
