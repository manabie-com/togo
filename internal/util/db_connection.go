package util

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/manabie-com/togo/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// CreateConnectionDB func
func CreateConnectionDB() (*gorm.DB, error) {
	Driver := config.Cfg.DbDriver
	DbHost := config.Cfg.DbHost
	DbUser := config.Cfg.DbUser
	DbPassword := config.Cfg.DbPassword
	DbName := config.Cfg.DbName
	DbPort := config.Cfg.DbPort
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	logLevel := logger.Warn
	if config.Cfg.Env == "local" {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Panicf("Error: %s", err)
	} else {
		fmt.Printf("Database %s connected\n", Driver)
	}
	return db, err
}
