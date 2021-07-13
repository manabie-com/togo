package jwtprovider

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	key       string
	expiresIn time.Duration
}

type JWTProvider interface {
	GenerateToken(payload map[string]interface{}) (string, error)
	Parse(token string) (map[string]interface{}, bool)
}

func NewJWTProvider(key string, expiresIn time.Duration) *jwtProvider {
	return &jwtProvider{
		key:       key,
		expiresIn: expiresIn,
	}
}

func (h *jwtProvider) GenerateToken(payload map[string]interface{}) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(h.expiresIn).Unix()
	for key, value := range payload {
		atClaims[key] = value
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(h.key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *jwtProvider) Parse(token string) (map[string]interface{}, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(h.key), nil
	})
	if err != nil {
		return nil, false
	}
	if !t.Valid {
		return nil, false
	}
	return claims, true
}
