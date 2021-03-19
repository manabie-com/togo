package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbUser     string
	DbPass     string
	DbName     string
	DbPort     string
	ServerPort string
}

func LoadConfig(name string) *Config {
	if name == "" {
		name = ".env"
	}

	err := godotenv.Load(name)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var envMap = Config{
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbPass:     os.Getenv("POSTGRES_PASSWORD"),
		DbName:     os.Getenv("POSTGRES_DB"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	return &envMap
}

//func ConnectDB() {
//
//}
