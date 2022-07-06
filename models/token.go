package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId        uint32
	LimitDayTasks uint
	jwt.StandardClaims
}

func (t *Token) CreateToken() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	return tokenString
}
