package dao

import (
	"togo/internal/model"
	"togo/internal/utils"

	"gorm.io/gorm"
)

type LoginDao interface {
	CheckCredential(uname string, pass string, db *gorm.DB) bool
}

func CheckCredential(uname string, pass string, db *gorm.DB) bool {

	var user model.User

	db.Table("users").Select("user_name", "password").Where("user_name = ?", uname).Scan(&user)

	if user.UserName == "" {
		return false
	}

	if utils.CheckPasswordHash(pass, user.Password) {
		return true
	}

	return false
}
