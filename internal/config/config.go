package config

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Config is the global config
	HOST string = "http://localhost:8080"
	PORT string

	ADMIN string

	DATABASE_URL string
	DB_HOST      string
	DB_PORT      string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string

	FREE_LIMIT int
	VIP_LIMIT  int
)

func Load() error {
	var err error
	projectDirName := "togo"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err = godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")

	ADMIN = os.Getenv("ADMIN")

	DATABASE_URL = os.Getenv("DATABASE_URL")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	FREE_LIMIT, err = strconv.Atoi(os.Getenv("FREE_LIMIT"))
	if err != nil {
		return err
	}

	VIP_LIMIT, err = strconv.Atoi(os.Getenv("VIP_LIMIT"))
	if err != nil {
		return err
	}

	return nil
}
