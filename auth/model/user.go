package model

import (
	"errors"
	"github.com/manabie-com/togo/model"
	apperror "github.com/manabie-com/togo/shared/app_error"
)

const (
	EntityName = "User"
)

var (
	ErrIdOrPasswordInvalid = apperror.NewCustomError(errors.New("id or password invalid"),
		"id or password invalid", "ErrIdOrPasswordInvalid")

	ErrUserHasBeenBlock = apperror.NewCustomError(errors.New("user has been block"),
		"user has been block", "ErrUserHasBeenBlock")
)

type User struct {
	model.BaseModel `json:",inline"`
	LoginId         string `json:"id" gorm:"column:login_id"`
	Password        string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func (u User) GetUserId() int {
	return u.Id
}
