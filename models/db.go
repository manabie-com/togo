package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// connect to database
func Connect() *sql.DB{ 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}	
	
	db, err := sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Connect to database failed")
	}
	return db
}

// Hash password into a crypt text
func Hash(password string) (string, error) { 
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

 // check a crypted text and a password user enter
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// interface implement
// func (r *Repository)Close() {
// 	r.DB.Close()
// }