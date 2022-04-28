package dao

import (
	"testing"
	"togo/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_LoginDao(t *testing.T) {

	DB_test, err := gorm.Open(sqlite.Open("../database/todo_test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB_test.AutoMigrate(&model.User{}, &model.Task{})

	// Test CheckCredential
	if CheckCredential("admin_kier", "admin123", DB_test) != true {
		t.Error("Expected true, got false")
	}

}
