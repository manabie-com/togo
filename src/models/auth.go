package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
