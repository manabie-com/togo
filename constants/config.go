package constants

import (
	"fmt"
	"os"
)

type DBType string

const (
	POSTGRES DBType = "POSTGRES"
	SQLITE DBType = "SQLITE"
)

type PostgreConfig struct {
	HOST string
	PORT string
	USER string
	PASSWORD string
	DB_NAME string
}

func GetPostgreConnectionString() string {
	config := PostgreConfig{
		HOST:     getEnv("DB_HOST","localhost"),
		PORT:     getEnv("DB_PORT", "5432"),
		USER:     getEnv("DB_USER", "postgres"),
		PASSWORD: getEnv("DB_PASSWORD", "mysecretpassword"),
		DB_NAME:  getEnv("DB_NAME", "togo-db"),
	}
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", config.HOST, config.PORT, config.USER, config.PASSWORD, config.DB_NAME)
}
var DB_TYPE = DBType(getEnv("DB_TYPE", string(POSTGRES)))

const (
	JWT_KEY = "wqGyEBBfPK9w3Lxw"
)

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}