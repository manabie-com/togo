package db

import (
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/manabie-com/togo/internal/models"
)

var (
	DB *gorm.DB
)

func InitDB(host, port, user, password, database string) {
	var err error

	dns := "host=" + host +" user="+ user +" password="+ password +" dbname="+ database +" port="+ port +" sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate([]models.Task{})
	db.AutoMigrate([]models.User{})
	DB = db
}

func getDB() *gorm.DB  {
	return DB
}
