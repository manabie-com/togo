package db

import (
	"fmt"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectDB() *gorm.DB {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		"localhost", config.NewEnv.DbUser, config.NewEnv.DbPass, config.NewEnv.DbName, config.NewEnv.DbPort, "Asia/Ho_Chi_Minh")

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println(fmt.Sprintf("Succeed to connect to database: %s", config.NewEnv.DbName))

	return conn
}

func DisconnectDB(db *gorm.DB) {
	conn, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database")
	}

	conn.Close()
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Task{})

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to migrate")
	}
}

func Seed(db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userRepo.AddUser(&models.User{Username: "huyha", Password: "123456", MaxTodo: 3})
}
