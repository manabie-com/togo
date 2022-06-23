package utils

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token hết hạn sau 12 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func UseToken(tokenString string) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}
