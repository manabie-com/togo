package token

import (
	"ManabieProject/helper"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type JWTMaker struct {
	secretKey string
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

const minSecretKeySize = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// NewSecretKey secret returns 32 bytes AES key
func (maker *JWTMaker) NewSecretKey() (string, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	secretKey, err := helper.GenKeyAEStoBase64()
	maker.secretKey = secretKey
	return secretKey, err
}

// CreateToken Create Token JWTMaker
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := InitPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken Verify Token JWTMaker
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		very, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(very.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
