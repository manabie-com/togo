package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

func ValidateRequest(req interface{}) error {
	v := validator.New()
	err := v.Struct(req)
	var result []string
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			result = append(result, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}
		return errors.New(strings.Join(result, ", "))
	}
	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidateDateFromString(dateStr string, layout string) bool {
	_, err := time.Parse(layout, dateStr)
	if err != nil {
		logger.Errorln(err)
		return false
	}
	return true
}

func ValidateInputIsInteger(input string) (int, bool) {
	if input == "" {
		return 0, false
	}
	value, err := strconv.Atoi(input)
	if err != nil {
		logger.Errorln(err)
		return value, false
	}
	return value, true
}
