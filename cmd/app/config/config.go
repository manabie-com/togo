package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBEnv struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func GetDBEnv() DBEnv {
	return DBEnv{
		DBHost: GetValue("DB_HOST"),
		DBPort: GetValue("DB_PORT"),
		DBUser: GetValue("DB_USER"),
		DBPass: GetValue("DB_PASS"),
		DBName: GetValue("DB_NAME"),
	}
}

// GetValue similar to os.Getenv("") but handles missing configs
func GetValue(configName string) string {
	config, exist := os.LookupEnv(configName)
	if !exist {
		log.Fatal(configName + " not set in .env")
	}
	return config
}
