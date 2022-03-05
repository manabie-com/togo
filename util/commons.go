package util

import (
	"regexp"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// nolint:lll
func IsEmail(str string) (bool, error) {
	match, err := regexp.MatchString(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, str)
	if err != nil {
		return false, err
	}

	return match, nil
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("empty password")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "hash password")
	}

	return string(hashBytes), nil
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
