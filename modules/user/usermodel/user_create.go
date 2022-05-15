package usermodel

import (
	"github.com/japananh/togo/common"
	"strings"
	"unicode"
)

var (
	ErrInvalidCharacterMsg         = "password has invalid characters"
	ErrNotEnoughCharacterMsg       = "password must have at least 8 characters"
	ErrMustHaveNumberMsg           = "password must have at least 1 number"
	ErrMustHaveLetterMsg           = "password must have at least 1 letter"
	ErrMustHaveSpecialCharacterMsg = "password must have at least 1 special character"
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Status          int    `json:"status" gorm:"column:status;default:1;"`
	DailyTaskLimit  int    `json:"daily_task_limit" gorm:"column:daily_task_limit;"`
	Email           string `json:"email" form:"email" binding:"required" gorm:"column:email;"`
	Password        string `json:"password" form:"password" binding:"required" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask() {
	u.GenUID(common.DbTypeUser)
}

func (u *UserCreate) Validate() error {
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	if u.DailyTaskLimit <= 0 {
		u.DailyTaskLimit = 5
	}
	if errMsg := VerifyPassword(u.Password); errMsg != "" {
		return ErrPasswordInvalid(errMsg)
	}

	return nil
}

func VerifyPassword(s string) string {
	hasNumber, hasLetter, hasSpecial, hasInvalidCharacter := false, false, false, false
	letterCount := 0

	for _, c := range s {
		letterCount++
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsLetter(c):
			hasLetter = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		default:
			hasInvalidCharacter = true
		}
	}

	if hasInvalidCharacter {
		return ErrInvalidCharacterMsg
	}

	if letterCount < 8 {
		return ErrNotEnoughCharacterMsg
	}

	if !hasNumber {
		return ErrMustHaveNumberMsg
	}

	if !hasLetter {
		return ErrMustHaveLetterMsg
	}

	if !hasSpecial {
		return ErrMustHaveSpecialCharacterMsg
	}

	return ""
}
