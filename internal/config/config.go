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

	DATABASE_URL      string
	DATABASE_HOST     string
	DATABASE_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	DB_NAME           string

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
		log.Fatal("Error loading .env file")
	}
	DATABASE_URL = os.Getenv("DATABASE_URL")
	PORT = os.Getenv("PORT")
	ADMIN = os.Getenv("ADMIN")

	DATABASE_URL = os.Getenv("DATABASE_URL")
	DATABASE_HOST = os.Getenv("DATABASE_HOST")
	DATABASE_PORT = os.Getenv("DATABASE_PORT")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	DB_NAME = os.Getenv("POSTGRES_DB")

	FREE_LIMIT, err = strconv.Atoi(os.Getenv("FREE_LIMIT"))
	if err != nil {
		return err
	}

	VIP_LIMIT, err = strconv.Atoi(os.Getenv("VIP_LIMIT"))
	if err != nil {
		return err
	}

	HOST = os.Getenv("LOCALHOST")

	return nil
}
