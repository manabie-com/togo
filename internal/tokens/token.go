package tokens

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/utils/constants"
	"time"
)

const (
	tokenExpiredTimeInMin = 15
)

type UserManager interface {
	ValidateUser(userID, password string) bool
}

type TokenManager struct {
	JWT      string
	UserManager UserManager
}

func NewTokenManager(jwt string, userManager UserManager) *TokenManager {
	return &TokenManager{
		JWT:      jwt,
		UserManager: userManager,
	}
}

func (t *TokenManager) GetAuthToken(userID, password string) (string, error) {
	if !t.UserManager.ValidateUser(userID, password) {
		return "", errors.New("incorrect user_id/pwd")
	}

	token, err := t.createToken(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *TokenManager) createToken(id string) (string, error) {
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

func (t *TokenManager) ValidToken(token string) (userID string, valid bool) {
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

