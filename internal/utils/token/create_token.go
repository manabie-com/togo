package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Create new JWT token with user id
func NewToken(userId string, jwtSecretKey string, tokenTimeOut time.Duration, issuer string) (string, error) {
	//Create token class with signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTimeOut).Unix(),
			Issuer:    issuer,
		},
	})

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
