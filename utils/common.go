package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	hashedPass := string(hashedByte)

	return hashedPass, err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func MapErrors(err validator.ValidationErrors) map[string]string {
	var message string

	for _, e := range err {
		if message != "" {
			message += fmt.Sprintf(" Invalid field %s.", e.Field())
			break
		}
		message += fmt.Sprintf("Invalid field %s.", e.Field())
	}

	return map[string]string{
		"message": message,
	}
}
