package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/services/users"

	"github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) login(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("Invalid username or password")
	}

	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("User %s not exist", username)
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passErr != nil {
		return "", errors.New("Wrong password")
	}

	token, err := genJWT(user, s.secretJWT)
	if err != nil {
		logrus.WithError(err).Errorf("Gen JWT failed")
		return "", errors.New("Internal error")
	}

	return token, nil
}

func genJWT(user *users.User, secretJWT string) (string, error) {
	if user == nil {
		return "", errors.New("user empty")
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["expired_time"] = time.Now().Add(time.Hour * 24).Unix()
	claims["token_type"] = "Bearer"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJWT))
}

// use to hash user password when user create account
func genHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logrus.WithError(err).Errorf("Gen hash from pwd failed")
		return "", err
	}
	return string(hash), nil
}
