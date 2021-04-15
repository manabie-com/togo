package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(jwtKey, userID string, expiredDuration time.Duration) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(expiredDuration).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(jwtKey, token string) (string, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errors.New("token is invalid")
	}
	id, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("unable to get user ID")
	}
	return id, nil
}
