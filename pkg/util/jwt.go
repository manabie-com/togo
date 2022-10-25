package util

import (
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("bb64719a7a0f14cdcceda03541bfbf81054d7360f37a149900665a67d2b89f36")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, role string, userID int, duration time.Duration) (string, error) {
	expireTime := time.Now().Add(duration)
	claims := Claims{
		base64.StdEncoding.EncodeToString([]byte(username)),
		base64.StdEncoding.EncodeToString([]byte(role)),
		userID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "vschool-api",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		} else {
			return claims, err
		}
	}
	return nil, err
}
