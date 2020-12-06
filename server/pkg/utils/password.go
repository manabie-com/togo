package utils

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(raw, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	if err != nil {
		return err
	}
	return nil
}

func GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
