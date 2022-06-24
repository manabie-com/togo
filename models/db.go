package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func Connect() { // connect to database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	DB, err = sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Connect to database failed")
	}
}
func Hash(password string) (string, error) { // Hash password into a crypt text
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error { // check a crypted text and a password user enter
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
