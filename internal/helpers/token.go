package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type tokenProvider struct {
	secretKey string
}

type TokenProvider interface {
	CreateToken(userId string) (string, error)
	ValidateToken(token string) (string, bool)
}

func NewTokenProvider(injectedSecretKey string) TokenProvider {
	return &tokenProvider{
		secretKey: injectedSecretKey,
	}
}

func (p *tokenProvider) CreateToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(p.secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (p *tokenProvider) ValidateToken(token string) (string, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(p.secretKey), nil
	})

	if err != nil {
		log.Println(err)
		return "", false
	}

	if !t.Valid {
		return "", false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", false
	}

	return id, true
}
