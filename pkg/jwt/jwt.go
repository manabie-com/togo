package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie/project/config"
	"time"
)

type tokenUser struct {
	conf *config.Config
}

type TokenUser interface {
	GenerateToken(username string) (string, error)
	ParseToken(tokenStr string) (string, error)
}

func NewTokenUser(conf *config.Config) TokenUser{
	return &tokenUser{
		conf:conf,
	}
}

func (t *tokenUser) GenerateToken(username string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"]= username
	atClaims["exp"] = time.Now().Add(time.Hour * 60 * 60 *60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.conf.Secret.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t *tokenUser) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.conf.Secret.JwtSecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "Missing Authentication Token", err
	}
}
