package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/core/port"
)

const (
	gJwtUserIdKey      = "user_id"
	gJwtExpiredTimeKey = "exp"
)

func NewJwtService(jwtKey string) port.JwtService {
	return &jwtService{
		jwtKey: jwtKey,
	}
}

type jwtService struct {
	jwtKey string
}

func (p *jwtService) CreateToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[gJwtUserIdKey] = userId
	atClaims[gJwtExpiredTimeKey] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(p.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *jwtService) ParseToken(token string) (string, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(p.jwtKey), nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", errors.New("invalid token")
	}

	userId, ok := claims[gJwtUserIdKey].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	return userId, nil
}
