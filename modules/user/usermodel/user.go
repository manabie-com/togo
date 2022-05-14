package usermodel

import (
	"errors"
	"github.com/japananh/togo/common"
)

const EntityName = "User"

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)

func ErrPasswordInvalid(msg string) *common.AppError {
	return common.NewCustomError(
		errors.New(msg),
		msg,
		"ErrPasswordInvalid",
	)
}

func ErrDailyLimitTaskInvalid(msg string) *common.AppError {
	return common.NewCustomError(
		errors.New(msg),
		msg,
		"ErrDailyLimitTaskInvalid",
	)
}

type User struct {
	common.SQLModel `json:",inline"`
	DailyTaskLimit  int    `json:"daily_task_limit" gorm:"column:daily_task_limit;"`
	Status          int    `json:"status" gorm:"status;default:1;"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) Mask() {
	u.GenUID(common.DbTypeUser)
}
