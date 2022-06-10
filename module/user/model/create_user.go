package model

import "togo/common"

type CreateUser struct {
	Name  *string `json:"name" gorm:"name"`
	Email *string `json:"email" gorm:"email"`
	*common.Model
}

func (CreateUser) TableName() string {
	return User{}.TableName()
}
