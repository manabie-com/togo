package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	absolutePath := "/home/jmramos02/Code/go/akaru/databases/tasks.db" //change this to your directory
	db, err := gorm.Open(sqlite.Open(absolutePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
