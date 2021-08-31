package utils

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// InitEnv is to initialize environment params
func InitEnv() error {
	var err error
	if flag.Lookup("test.v") == nil {
		err = godotenv.Load()
	} else {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		return err
	}
	return nil
}

// InitDB is to initialize DB
func InitDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"))

	return sql.Open("postgres", psqlInfo)
}
