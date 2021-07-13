package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/pkg/token"
	"time"
)

type Generator struct {
	Cfg *config.Config
}

func (g *Generator) CreateToken(userID string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[token.UserIDField] = userID
	atClaims[token.ExpiredField] = time.Now().Add(g.Cfg.SSExpire).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenStr, err := at.SignedString([]byte(g.Cfg.JWTKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (g *Generator) ValidateToken(tokenStr string) (string, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(tokenStr, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(g.Cfg.JWTKey), nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", errors.New(dictionary.TokenIsNotValid)
	}

	id, ok := claims[token.UserIDField].(string)
	if !ok {
		return "", errors.New(dictionary.FailedToParseUserID)
	}

	return id, nil
}
