package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword return hashed password from input password
func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hashed password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword return result of matching between password and hashedPassword
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
