package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is an database connection instanace that will be used in entity repository
var DB *gorm.DB

func ConnectPostgre() {
	dbHost := os.Getenv("POSTGRE_DB_HOST")
	dbPort := os.Getenv("POSTGRE_DB_PORT")
	dbName := os.Getenv("POSTGRE_DB_NAME")
	dbUsername := os.Getenv("POSTGRE_DB_USERNAME")
	dbPassword := os.Getenv("POSTGRE_DB_PASSWORD")

	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort)

	env := os.Getenv("ENV_MODE")
	dbLogger := logger.Default.LogMode(logger.Silent)
	if env == "dev" {
		dbLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		Logger: dbLogger,
	})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Cannot connect to Postgre database", err)
		panic(err)
	} else {
		fmt.Println("Connected to Postgre database!")
	}

	// Set update default connections pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	DB = db
	fmt.Println("Database connection has been established")
}
