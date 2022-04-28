package dao

import (
	"togo/internal/model"

	"gorm.io/gorm"
)

type UserDao interface {
	MaxDailyTodo(id uint, db *gorm.DB) int64
}

func MaxDailyTodo(id uint, db *gorm.DB) int64 {

	var users *model.User

	db.Table("users").Select("max_todos").Where("id = ?", id).Scan(&users)

	return int64(users.MaxTodos)

}

func CheckUserExist(uname string, db *gorm.DB) bool {

	var user model.User

	db.Table("users").Select("user_name").Where("user_name = ?", uname).Scan(&user)

	return user.UserName == ""

}
