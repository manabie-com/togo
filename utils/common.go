package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.NewEnv.JwtExpires)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.NewEnv.JwtSecret))
}

func ExtractToken(r *http.Request) string {
	token := strings.TrimSpace(r.Header.Get("Authorization"))
	return token
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.NewEnv.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

func ExtractTokenMetadata(token *jwt.Token) string {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%v", claims["username"])
	}
	return ""
}

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
