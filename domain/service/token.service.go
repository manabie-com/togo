package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
	"togo/domain/errdef"
	"togo/domain/model"
)

type TokenService interface {
	CreateToken(ctx context.Context, u model.User) (string, error)
	ValidateToken(ctx context.Context, token string) (*model.UserClaim, error)
}

type tokenServiceImpl struct {
	secret string
}

func (this *tokenServiceImpl) CreateToken(ctx context.Context, u model.User) (string, error) {
	t := model.UserClaim{
		UserId: u.Id,
		Limit:  u.Limit,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).UnixMilli(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	signedToken, err := token.SignedString([]byte(this.secret))
	return signedToken, err
}

func (this *tokenServiceImpl) ValidateToken(ctx context.Context, signedToken string) (*model.UserClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.UserClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(this.secret), nil
		},
	)
	claims, ok := token.Claims.(*model.UserClaim)
	if !ok {
		return nil, errdef.TokenWrongFormat
	}
	return claims, err
}

func NewTokenService(secret string) TokenService {
	return &tokenServiceImpl{
		secret: secret,
	}
}
