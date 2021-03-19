package config

import (
	"fmt"
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB //database

func ConnectDB() {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		"localhost", NewEnv.DbUser, NewEnv.DbPass, NewEnv.DbName, NewEnv.DbPort, "Asia/Ho_Chi_Minh")

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println(fmt.Sprintf("Succeed to connect to database: %s", NewEnv.DbName))

	DB = conn

}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Task{})

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to migrate")
	}
}

func Seed() {
	UserRepo := repositories.UserRepository{
		DB: DB,
	}
	UserRepo.AddUser(&models.User{Username: "huyha", Password: "123456", MaxTodo: 3})
}
