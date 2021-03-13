package auth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrEmptyKeyFunc         = errors.New("empty jwt KeyFunc")
	ErrUnknownSigningMethod = errors.New("unknown signing method")
	ErrNoAccessToken        = errors.New("no access token")
	ErrUnknownUser          = errors.New("unknown user")
)

type JWTConfig struct {
	jwt.Keyfunc
	Issuer string
	jwt.SigningMethod
}

type JWTConfigOption func(*JWTConfig)

func WithKeyFunc(secret string) JWTConfigOption {
	return func(c *JWTConfig) {
		c.Keyfunc = jwt.Keyfunc(func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	}
}

func WithIssuer(issuer string) JWTConfigOption {
	return func(c *JWTConfig) {
		c.Issuer = issuer
	}
}

func WithSigningMethod(algorithm string) JWTConfigOption {
	return func(c *JWTConfig) {
		c.SigningMethod = jwt.GetSigningMethod(algorithm)
	}
}

func NewJWTConfig(options ...JWTConfigOption) (*JWTConfig, error) {
	c := &JWTConfig{}

	for _, o := range options {
		o(c)
	}

	if c.Keyfunc == nil {
		return nil, fmt.Errorf("creating jwt config: %w", ErrEmptyKeyFunc)
	}

	if c.SigningMethod == nil {
		return nil, fmt.Errorf("creating jwt config: %w", ErrUnknownSigningMethod)
	}

	return c, nil
}
