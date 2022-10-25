package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
	"togo.com/pkg/model"
	"togo.com/pkg/repository"
)

const (
	Signature = "secret"
)

type AuthorizeUseCase interface {
	Login(ctx context.Context, request model.LoginRequest) (resp model.LoginResponse, err error)
}

type authorizeUseCase struct {
	repo repository.Repository
}

func NewAuthorizeUseCase(repo repository.Repository) AuthorizeUseCase {
	return authorizeUseCase{repo: repo}
}

func (a authorizeUseCase) Login(ctx context.Context, request model.LoginRequest) (resp model.LoginResponse, err error) {
	//validate  user
	id, err := a.repo.GetUser(ctx, request)
	if err != nil {
		return model.LoginResponse{}, err
	}
	if id == "" {
		return model.LoginResponse{}, errors.New("invalid user")
	}
	// Set custom claims
	claims := &model.JwtCustomClaims{
		Name:  id,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(Signature))
	if err != nil {
		return model.LoginResponse{}, err
	}

	return model.LoginResponse{Token: t}, err
}

func ValidateToken(tokenString string) (userId string, errString string) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(Signature), nil
	})
	if token == nil || err != nil {
		fmt.Printf("Error %s", err)
		return "", "Unauthenticated"
	}
	if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
		return claims.Name, ""
	}
	return "", "Unauthenticated"

}
