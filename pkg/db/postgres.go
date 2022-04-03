package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New creates new PostgreSQL database connection
func New(dsn string, enableLog bool) (*gorm.DB, error) {
	dbConfig := &gorm.Config{}
	if enableLog {
		logger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: false,
				Colorful:                  false,
			},
		)
		dbConfig.Logger = logger
	}
	db, err := gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
