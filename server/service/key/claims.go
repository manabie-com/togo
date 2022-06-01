package key

import (
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
