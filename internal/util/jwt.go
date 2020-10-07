package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims struct
type Claims struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// CreateToken func
func CreateToken(claims *Claims, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

// CreateAuthTokenPair func
func CreateAuthTokenPair(email string, id uint64, jwtKey string, expMinute int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expMinute) * time.Minute)
	claims := &Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	authToken, err := CreateToken(claims, jwtKey)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

// DecodeToken func
func DecodeToken(token string, claims *Claims, jwtKey string) error {

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return jwt.ErrSignatureInvalid
	}
	return nil
}
