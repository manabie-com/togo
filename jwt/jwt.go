package jwt

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Create(w http.ResponseWriter, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token expired after 12 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func CheckToken(tokenString string) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString, claims, nil)

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}
