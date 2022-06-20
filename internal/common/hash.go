package common

import (
	"golang.org/x/crypto/bcrypt"
)

const hashCost = bcrypt.DefaultCost + 5

func HashPassword(pwd string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pwd), hashCost)
	return string(hashed)
}
