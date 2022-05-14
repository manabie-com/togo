package usermodel

import (
	"github.com/japananh/togo/component/tokenprovider"
)

type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required" gorm:"column:email;"`
	Password string `json:"password" form:"password" binding:"required" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}
