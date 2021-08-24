package adapter

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/common"
)

const (
	jwtClaimsUserID = "user_id"
	jwtClaimsExp    = "exp"
)

type (
	JWTAdapter interface {
		CreateToken(ctx context.Context, userID string) (string, error)
		VerifyToken(ctx context.Context, token string) (string, error)
	}

	jwtAdapter struct {
		jwtKey string
	}
)

func NewJWTAdapter(jwtKey string) JWTAdapter {
	return &jwtAdapter{
		jwtKey: jwtKey,
	}
}

func (a *jwtAdapter) CreateToken(ctx context.Context, userID string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[jwtClaimsUserID] = userID
	atClaims[jwtClaimsExp] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(a.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *jwtAdapter) VerifyToken(ctx context.Context, token string) (string, error) {
	claims := make(jwt.MapClaims)
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(a.jwtKey), nil
	})
	if err != nil {
		log.Printf("parse with claims error: %v", err)
		return "", errors.New(common.ReasonInvalidToken.Code())
	}

	if !jwtToken.Valid {
		log.Println("jwtToken invalid")
		return "", errors.New(common.ReasonInvalidToken.Code())
	}

	userID, ok := claims[jwtClaimsUserID].(string)
	if !ok {
		log.Println("claims user id not ok")
		return "", errors.New(common.ReasonInvalidToken.Code())
	}
	return userID, nil
}
