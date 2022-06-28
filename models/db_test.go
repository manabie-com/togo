package models

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestDB(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
	DB_URI := os.Getenv("DB_URI")
	db := Connect(DB_URI)
	if db == nil {
		t.Fatal("Error connect database")
	}
}

