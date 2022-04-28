package dao

import (
	"testing"
	"togo/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCountDailyTask(t *testing.T) {
	DB_test, err := gorm.Open(sqlite.Open("../database/todo_test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB_test.AutoMigrate(&model.User{}, &model.Task{})

	count := CountDailyTask(1, DB_test)
	if count != 0 {
		t.Error("Expected 0, got ", count)
	}

}
