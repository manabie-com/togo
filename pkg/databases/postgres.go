package databases

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		sslMode = "require"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_PORT"),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
