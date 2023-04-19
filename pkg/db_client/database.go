package db_client

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	//db.AutoMigrate(&models.User{})
	//db.AutoMigrate(&models.Task{})
	return Handler{db}
}
