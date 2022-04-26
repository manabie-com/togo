package user

import (
	"togo/internal/connect"
	"togo/internal/model"
)

var users model.User

func GetMaxDailyTodo(id uint) int64 {

	connect.DB.Table("users").Select("max_todos").Where("id = ?", id).Scan(&users)

	return int64(users.MaxTodos)

}
