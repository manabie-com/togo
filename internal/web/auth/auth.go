// Package auth provides authentication support.
package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidToken  = errors.New("JWT token missing or invalid")
	ErrInvalidClaims = errors.New("Invalid claims")
)

// Claims represents the authorization claims transmitted via a JWT.
type Claims struct {
	jwt.RegisteredClaims
}

// Config represents information required to initialize auth.
type Config struct {
	SigningKey string
}

// Auth is used to authenticate clients. It can generate a token for a
// set of user claims and recreate the claims by parsing the token.
type Auth struct {
	method     jwt.SigningMethod
	signingKey []byte
}

// New creates an Auth to support authentication.
func New(cfg Config) (*Auth, error) {
	a := Auth{
		method:     jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		signingKey: []byte(cfg.SigningKey),
	}

	return &a, nil
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func (a *Auth) GenerateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)

	str, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return str, nil
}

func (a *Auth) EchoMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: a.signingKey,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &Claims{}
		},
	})
}

func (a *Auth) GetClaims(c echo.Context) (*Claims, error) {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}

func (a *Auth) GetUserID(c echo.Context) (uuid.UUID, error) {
	claims, err := a.GetClaims(c)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("getting claims: %w", err)
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parsing subject: %w", err)
	}

	return userID, nil
}
