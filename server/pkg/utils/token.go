package utils

import "github.com/dgrijalva/jwt-go"

func GetClaimsFromToken(token, jwtKey string)  (*jwt.MapClaims, error) {
	return nil, nil
}

func CreateToken(id string, jwtKey string) (string, error) {
	return "", nil
}