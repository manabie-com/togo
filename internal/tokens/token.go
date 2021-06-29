package tokens

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/utils/constants"
	"time"
)

const (
	tokenExpiredTimeInMin = 15
)

type TokenManager interface {
	GetAuthToken(userID, password string) (string, error)
	ValidToken(token string) (userID string, valid bool)
}

type TokenManagerImpl struct {
	JWT      string
	UserRepo repositories.UserRepo
}

func NewTokenManager(jwt string, userRepo repositories.UserRepo) *TokenManagerImpl {
	return &TokenManagerImpl{
		JWT:      jwt,
		UserRepo: userRepo,
	}
}

func (t *TokenManagerImpl) GetAuthToken(userID, password string) (string, error) {
	if !t.UserRepo.ValidateUser(userID, password) {
		return "", errors.New("incorrect user_id/pwd")
	}

	token, err := t.createToken(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *TokenManagerImpl) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[constants.UserIDKey] = id
	atClaims["exp"] = time.Now().Add(time.Minute * tokenExpiredTimeInMin).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.JWT))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t *TokenManagerImpl) ValidToken(token string) (userID string, valid bool) {
	claims := make(jwt.MapClaims)
	tok, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(t.JWT), nil
	})
	if err != nil {
		return "", false
	}

	if !tok.Valid {
		return "", false
	}

	id, ok := claims[constants.UserIDKey].(string)
	if !ok {
		return "", false
	}

	return id, true
}

