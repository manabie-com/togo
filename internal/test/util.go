package test

import (
	"github.com/jmramos02/akaru/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDB() *gorm.DB {
	absolutePath := "/home/jmramos02/Code/go/akaru/databases/tasks_test.db" //change this to your directory
	db, err := gorm.Open(sqlite.Open(absolutePath), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&model.Task{}, &model.User{})
	return db
}

func CreateTaskTestData(db *gorm.DB, userID int, name string) model.Task {
	task := model.Task{
		UserID: userID,
		Name:   "Buy Milk",
	}

	db.Save(&task)

	return task
}

func CreateUserTestData(db *gorm.DB, username string, limit int) model.User {
	user := model.User{
		Username: username,
		Limit:    limit,
	}

	db.Save(&user)
	return user
}
