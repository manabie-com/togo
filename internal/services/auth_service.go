package services

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/ent"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, credential model.LoginCredential) (*model.AccessToken, error)

	CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error)
}

type authServiceImpl struct {
	jwtKey         string
	userRepository storages.UserRepository
}

func NewAuthService(jwtKey string, taskRepository storages.UserRepository) AuthService {

	auth := &authServiceImpl{jwtKey: jwtKey, userRepository: taskRepository}
	return auth
}

func (a *authServiceImpl) Login(ctx context.Context, credential model.LoginCredential) (*model.AccessToken, error) {
	foundUser, err := a.userRepository.FindByUsername(ctx, credential.UserName)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credential.Password))
	if err != nil {
		return nil, err
	}

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = foundUser.UserID
	atClaims["max_todo"] = foundUser.MaxTodo

	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(a.jwtKey))
	if err != nil {
		return nil, err
	}
	return &model.AccessToken{Token: token}, nil

}

func (a *authServiceImpl) CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error) {
	cipherPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return a.userRepository.CreateUser(ctx, username, string(cipherPwd))
}
