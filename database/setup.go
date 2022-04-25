package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/manabie-com/togo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "manabie.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Todo{})
	database.AutoMigrate(&models.User{})

	DB = database
}

func DisconnectDatabase() {
	DB.Close()
}
