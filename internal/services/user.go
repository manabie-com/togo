package services

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
)

type UserAuthKey int8

type UserService struct {
	JWTKey string
	storages.Storage
}

func (s *UserService) UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

func (s *UserService) GetAuthToken(ctx context.Context, userID string, password string) (string, error) {
	id := value(userID)

	if !s.Storage.ValidateUser(ctx, id, value(password)) {
		return "", errors.New("incorrect user_id/pwd")
	}

	token, err := s.createToken(id.String)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) ValidToken(token string) (string, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", err
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("cannot get user id")
	}

	return id, nil
}

func (s *UserService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
