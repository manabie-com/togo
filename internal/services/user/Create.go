package services

import (
	"togo/internal/connect"
	"togo/internal/model"
)

func Create() int64 {

	user := model.User{UserName: "test", Password: "test", MaxTodos: 10}

	result := connect.DB.Create(&user)
	return result.RowsAffected
}
