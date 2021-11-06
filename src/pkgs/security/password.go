package security

import (
	"os"
	"strconv"

	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	saltRounds, _ := strconv.Atoi(os.Getenv(settings.SaltRounds))
	return bcrypt.GenerateFromPassword([]byte(password), saltRounds)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
