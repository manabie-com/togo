package db

import (
	"fmt"
	"log"
	"os"
	"product-api/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	//connectingString:= "sqlserver://username:password@host:port?database=nameDb"

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	//dsn := fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",dbHost,dbUsername,dbPassword,dbDatabase,dbPort)


	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		panic("Failed to connect database")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Task{})

	DB = database
	return DB
}
