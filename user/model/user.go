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
	ErrUserHasBeenBlock = apperror.NewCustomError(errors.New("user has been block"),
		"user has been block", "ErrUserHasBeenBlock")
)

type User struct {
	model.BaseModel
	LogInId  string `gorm:"column:login_id"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func (u User) GetUserId() int {
	return u.Id
}
