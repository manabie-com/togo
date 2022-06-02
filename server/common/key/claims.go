package key

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
