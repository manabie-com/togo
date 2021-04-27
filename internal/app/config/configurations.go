package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
type App struct {
	Port      string
	SecretKey string
}
type Config struct {
	Database Database
	App      App
}

func LoadConfigs() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	return Config{
		Database{
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		}, App{
			Port:      os.Getenv("API_PORT"),
			SecretKey: os.Getenv("SECRET_KEY"),
		},
	}
}
