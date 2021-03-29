package core

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	authSubKey = "sub"
	authExpKey = "exp"
)

var (
	AuthTokenIsInvalid = errors.New("token is invalid")
	JwtKeyIsEmpty      = errors.New("jwt is empty")
)

type AppAuthenticator interface {
	CreateToken(userId int64) (string, error)
	ValidateToken(req *http.Request) (int64, error)
}

func NewAppAuthenticator() (AppAuthenticator, error) {
	jwtKey := os.Getenv("JWT_TOKEN")
	if len(jwtKey) == 0 {
		return nil, JwtKeyIsEmpty
	}
	return AppAuthenticate{
		jwtKey: jwtKey,
	}, nil
}

type AppAuthenticate struct {
	jwtKey string
}

func (a AppAuthenticate) CreateToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		authSubKey: userId,
		authExpKey: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtKey))
}

func (a AppAuthenticate) ValidateToken(req *http.Request) (int64, error) {
	authToken := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	parsedToken, err := jwt.ParseWithClaims(authToken, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(a.jwtKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, AuthTokenIsInvalid
	}

	id, err := strconv.ParseInt(fmt.Sprintf("%v", claims[authSubKey]), 10, 64)
	if err != nil {
		return 0, AuthTokenIsInvalid
	}

	return id, nil
}
