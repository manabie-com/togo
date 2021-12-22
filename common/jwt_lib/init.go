package jwt_lib

import "github.com/golang-jwt/jwt"

type baseClaims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}
