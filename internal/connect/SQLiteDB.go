package connect

import (
	"togo/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(sqlite.Open("../database/todo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{}, &model.Task{})
}
