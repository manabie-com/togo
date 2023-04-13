package usermodel

import "github.com/phathdt/libs/go-sdk/sdkcm"

type User struct {
	sdkcm.SQLModel `json:",inline"`
	Email          string `json:"email" gorm:"column:email;"`
	Password       string `json:"password" gorm:"column:password;"`
	Salt           string `json:"-" gorm:"column:salt;"`
	LimitTask      string `json:"limit_task" gorm:"column:limit_task"`
}

func (u User) TableName() string {
	return "users"
}
