package user

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository"
)

type UserUsecase interface {
	ValidateUser(ctx context.Context, username, password sql.NullString) bool
	GetUserByUserName(ctx context.Context, username sql.NullString) (*models.User, error)
	GenerateToken(userID, maxTaskPerDay uint) (string, error)
}

type userUsecase struct {
	repository repository.DatabaseRepository
}

func NewUserUsecase(repository repository.DatabaseRepository) UserUsecase {
	return &userUsecase{repository}
}

func (u *userUsecase) ValidateUser(ctx context.Context, username, password sql.NullString) bool {
	return u.repository.ValidateUser(ctx, username, password)
}

func (u *userUsecase) GetUserByUserName(ctx context.Context, username sql.NullString) (*models.User, error) {
	return u.repository.GetUserByUserName(ctx, username)
}

func (u *userUsecase) GenerateToken(userID, maxTaskPerDay uint) (string, error) {
	// Init a map claim for storing essential info
	claims := jwt.MapClaims{}

	timeout, err := strconv.Atoi(os.Getenv("JWT_TIMEOUT"))
	if err != nil {
		return "", err
	}

	claims["user_id"] = userID
	claims["max_task_per_day"] = maxTaskPerDay
	// Must use 'exp' key for storing timeout info
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
