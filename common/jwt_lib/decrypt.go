package jwt_lib

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/shanenoi/togo/config"
)

func Decrypt(tokenString string) (data map[string]interface{}, err error) {
	var token *jwt.Token
	data = map[string]interface{}{}

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(config.JWT_WRONG_METHOD, token.Header["alg"])
		}

		return []byte(config.PrivateKey()), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data = claims["data"].(map[string]interface{})
	}

	return
}
