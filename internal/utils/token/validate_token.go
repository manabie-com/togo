package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Validate(tokenString string, jwtSecretKey string) (bool, *CustomClaims) {
	tokenVal, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	fmt.Println(err)
	if err != nil {
		return false, nil
	}
	claims, ok := tokenVal.Claims.(*CustomClaims)
	if !ok || !tokenVal.Valid {
		return false, nil
	}
	return true, claims
}
