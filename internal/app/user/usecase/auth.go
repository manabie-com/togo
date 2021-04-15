package usecase

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/util"
	"github.com/rs/zerolog/log"
	"time"
)

const tokenExpiredDuration = 15 * time.Minute

var generateToken = util.GenerateToken
var parseToken = util.ParseToken
var (
	ErrInvalidUser           = errors.New("incorrect user_id/pwd")
	ErrUnableToGenerateToken = errors.New("unable to generate token")
)

type AuthService struct {
	userStorage UserStorage
	jwtKey      string
}

func NewAuthService(storage UserStorage, jwtKey string) *AuthService {
	return &AuthService{userStorage: storage, jwtKey: jwtKey}
}

//go:generate mockgen -package mock -destination mock/auth_mock.go github.com/manabie-com/togo/internal/app/user/usecase UserStorage
type UserStorage interface {
	ValidateUser(ctx context.Context, userID, pwd string) error
}

func (a AuthService) GetAuthToken(ctx context.Context, userID, pwd string) (string, error) {
	err := a.userStorage.ValidateUser(ctx, userID, pwd)
	if err != nil {
		log.Error().Str("userID", userID).Err(err).Msg(ErrInvalidUser.Error())
		return "", ErrInvalidUser
	}
	token, err := generateToken(a.jwtKey, userID, tokenExpiredDuration)
	if err != nil {
		log.Error().Str("userID", userID).Err(err).Msg("unable to generate token")
		return "", ErrUnableToGenerateToken
	}
	return token, nil
}

func (a AuthService) Authorize(token string) (string, error) {
	return parseToken(a.jwtKey, token)
}
