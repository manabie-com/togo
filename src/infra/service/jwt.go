package service

import (
	"errors"
	"time"
	"togo/src"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type JWT struct {
}

func (this *JWT) CreateToken(data *src.TokenData) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":     data.UserId,
		"exp":         time.Now().Add(time.Hour * 72).Unix(),
		"permissions": data.Permissions,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (this *JWT) VerifyToken(token string) (*src.TokenData, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	if t.Valid {
		tokenData := new(src.TokenData)
		if err := mapstructure.Decode(claims, tokenData); err != nil {
			return nil, err
		}
		return tokenData, nil
	}

	return nil, errors.New("token invalid")
}

func NewJWTService() src.IJWTService {
	return &JWT{}
}
