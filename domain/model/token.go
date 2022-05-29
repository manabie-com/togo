package model

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	jwt.StandardClaims
	UserId int
	Limit  int
}
