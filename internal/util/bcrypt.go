package util

import "golang.org/x/crypto/bcrypt"

// HashPassword func
func HashPassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword)
}

// CompareHashPassword func
func CompareHashPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
