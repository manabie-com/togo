package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	DbUser        string
	DbPass        string
	DbName        string
	DbNameTesting string
	DbPort        string
	ServerPort    string
}

var NewEnv *Env

func LoadEnv(name string) *Env {
	if name == "" {
		name = ".env"
	}

	err := godotenv.Load(name)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var envMap = Env{
		DbUser:        os.Getenv("POSTGRES_USER"),
		DbPass:        os.Getenv("POSTGRES_PASSWORD"),
		DbName:        os.Getenv("POSTGRES_DB"),
		DbNameTesting: os.Getenv("POSTGRES_DB_TESTING"),
		DbPort:        os.Getenv("POSTGRES_PORT"),
		ServerPort:    os.Getenv("SERVER_PORT"),
	}

	NewEnv = &envMap

	return NewEnv
}
