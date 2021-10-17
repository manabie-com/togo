package user

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
)

type UserUseCase interface {
	ValidateUser(ctx context.Context, username, password sql.NullString) bool
	GenerateToken(userID, maxTodo uint) (string, error)
	GetUserByUsername(ctx context.Context, username sql.NullString) (*storages.User, error)
}

type userUseCase struct {
	storeRepository storages.Repository
}

func NewUserUseCase(storeRepository storages.Repository) UserUseCase {
	return &userUseCase{storeRepository: storeRepository}
}

func (u *userUseCase) GetUserByUsername(ctx context.Context, username sql.NullString) (*storages.User, error) {
	return u.storeRepository.GetUserByUsername(ctx, username)
}

func (u *userUseCase) ValidateUser(ctx context.Context, username, password sql.NullString) bool {
	return u.storeRepository.ValidateUser(ctx, username, password)
}

func (u *userUseCase) GenerateToken(userID uint, maxTodo uint) (string, error) {
	return u.createToken(userID, maxTodo)
}

func (u *userUseCase) createToken(id uint, maxTodo uint) (string, error) {
	atClaims := jwt.MapClaims{}

	timeout, err := strconv.Atoi(os.Getenv("JWT_TIMEOUT"))
	if err != nil {
		return "", err
	}

	atClaims["user_id"] = id
	atClaims["max_todo"] = maxTodo

	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
