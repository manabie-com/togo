package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	secretKey string
}

const keySize = 32

func NewJwt(secretKey string) (Tokenizer, error) {
	if len(secretKey) < keySize {
		return nil, fmt.Errorf("key size must be at least %d characters", keySize)
	}
	return &Jwt{secretKey}, nil
}

//Create creates the jwt token
func (j *Jwt) Create(payload *Payload) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(j.secretKey))
}

//Verify verifies the jwt token
func (j *Jwt) Verify(tokenString string) (*Payload, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) && errors.Is(ve.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken

	}

	payload, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
