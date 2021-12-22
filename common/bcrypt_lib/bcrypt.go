package bcrypt_lib

import (
	"github.com/shanenoi/togo/config"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	data, err := bcrypt.GenerateFromPassword([]byte(password), config.BCRYPT_COST)
	return string(data), err
}

func CompareHashAndPassword(passwordDigest string, rawPassword string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(rawPassword))
	return result == nil
}
