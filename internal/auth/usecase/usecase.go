package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"kienmatu/go-todos/internal/auth"
	"kienmatu/go-todos/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type authUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL int64) auth.UseCase {
	return &authUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * time.Duration(tokenTTL),
	}
}

func (a *authUseCase) SignUp(ctx context.Context, username, password string, limit int) (*models.User, error) {
	fmtusername := strings.ToLower(username)
	euser, _ := a.userRepo.GetUserByUsername(ctx, fmtusername)

	if euser != nil {
		return nil, auth.ErrUserExisted
	}
	user := &models.User{
		Id:       uuid.New().String(),
		Username: fmtusername,
		Password: password,
		Limit:    limit,
	}
	user.HashPassword()
	err := a.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return a.userRepo.GetUserByUsername(ctx, username)
}

func (a *authUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, _ := a.userRepo.GetUserByUsername(ctx, username)
	if user == nil {
		return "", auth.ErrUserNotFound
	}

	if !user.ComparePassword(password) {
		return "", auth.ErrWrongPassword
	}

	claims := AuthClaims{
		Username: user.Username,
		UserId:   user.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "go-todos",
			ExpiresAt: time.Now().Add(a.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *authUseCase) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", auth.ErrInvalidAccessToken
}
