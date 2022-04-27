package login

import (
	"fmt"
	"togo/internal/connect"
	"togo/internal/model"
	"togo/internal/utils"
)

var user model.User

func Validate(uname string, pass string) bool {
	connect.DB.Table("users").Select("user_name", "password").Where("user_name = ?", uname).Scan(&user)
	fmt.Println(fmt.Sprintf("username: %s", user.UserName))
	fmt.Println(fmt.Sprintf("password: %s", user.Password))

	fmt.Println(user.Password)
	if user.UserName == "" {
		return false
	}

	if utils.CheckPasswordHash(pass, user.Password) {
		return true
	}

	return false
}
