package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Env struct {
	DbUser        string
	DbPass        string
	DbName        string
	DbNameTesting string
	DbPort        string
	ServerPort    string
	JwtSecret     string
	JwtExpires    int64
}

var NewEnv *Env

func LoadEnv() *Env {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	expires, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRES"), 10, 64)

	if err != nil {
		log.Fatal("Invalid JWT_EXPIRES in .env file")
	}

	var envMap = Env{
		DbUser:        os.Getenv("POSTGRES_USER"),
		DbPass:        os.Getenv("POSTGRES_PASSWORD"),
		DbName:        os.Getenv("POSTGRES_DB"),
		DbNameTesting: os.Getenv("POSTGRES_DB_TESTING"),
		DbPort:        os.Getenv("POSTGRES_PORT"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		JwtSecret:     os.Getenv("JWT_SECRET"),
		JwtExpires:    expires,
	}

	NewEnv = &envMap

	return NewEnv
}
