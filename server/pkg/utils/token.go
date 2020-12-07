package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetClaimsFromToken(token, jwtKey string)  (jwt.MapClaims, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("token invalid")
	}

	return claims, nil
}

func CreateToken(id string, jwtKey string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}