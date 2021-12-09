package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type hashPassword struct {

}

type Hash interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) error
}

func NewHashPassword() Hash {
	return &hashPassword{}
}

func(h *hashPassword) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func(h *hashPassword) CheckPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
