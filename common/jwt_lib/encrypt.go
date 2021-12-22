package jwt_lib

import (
	"github.com/golang-jwt/jwt"
	"github.com/shanenoi/togo/config"
	"time"
)

func expiredAt(duration time.Duration) int64 {
	return time.Now().Local().Add(duration).Unix()
}

func Encrypt(data interface{}) (string, error) {
	claims := baseClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: expiredAt(time.Hour * time.Duration(config.JWT_EXPIRED_HOUR)),
			Issuer:    config.JWT_ISSUER,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.PrivateKey()))
}
