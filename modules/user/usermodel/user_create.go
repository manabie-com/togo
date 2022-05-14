package usermodel

import (
	"github.com/japananh/togo/common"
	"strings"
	"unicode"
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
	if errMsg := verifyPassword(u.Password); errMsg != "" {
		return ErrPasswordInvalid(errMsg)
	}

	return nil
}

func verifyPassword(s string) string {
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
		return "password has invalid characters"
	}

	if letterCount < 8 {
		return "password must have at least 8 characters"
	}

	if !hasNumber {
		return "password must have at least 1 number"
	}

	if !hasLetter {
		return "password must have at least 1 letter"
	}

	if !hasSpecial {
		return "password must have at least 1 speicial character"
	}

	return ""
}
