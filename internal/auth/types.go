package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	RateLimit int    `json:rate_limit`
	jwt.StandardClaims
}
