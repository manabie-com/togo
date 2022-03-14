package usecase

import (
	"context"
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
		fmt.Printf("GetUser error %s", err)
	}
	// Set custom claims
	claims := &model.JwtCustomClaims{
		Name:  id,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
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
		return "", "Parse token error"
	}

	claims := token.Claims.(*model.JwtCustomClaims)
	fmt.Println(claims.Name)
	if token.Valid {
		fmt.Println("You look nice today")
		return claims.Name, ""
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "", "Timing is everything"
		} else {
			return "", fmt.Sprintf("Couldn't handle this token:%s", err)
		}
	} else {
		return "", fmt.Sprintf("Couldn't handle this token:%s", err)
	}

}
