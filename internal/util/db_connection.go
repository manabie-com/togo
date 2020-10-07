package util

import (
	"fmt"
	"log"

	"github.com/manabie-com/togo/internal/config"

	"github.com/jinzhu/gorm"
)

// CreateConnectionDB func
func CreateConnectionDB() (*gorm.DB, error) {
	Driver := config.Cfg.DbDriver
	DbHost := config.Cfg.DbHost
	DbUser := config.Cfg.DbUser
	DbPassword := config.Cfg.DbPassword
	DbName := config.Cfg.DbName
	DbPort := config.Cfg.DbPort
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Driver, url)
	// defer db.Close()

	if err != nil {
		log.Panicf("Error: %s", err)
	} else {
		fmt.Printf("Database %s connected\n", Driver)

		if config.Cfg.Env == "local" {
			db.LogMode(true)
		}
	}
	return db, err
}
