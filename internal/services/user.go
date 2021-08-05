package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/manabie-com/togo/pkg/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

type UserService interface {
	Authentication(user *model.User) (string, error)
	CreateUser(user *model.User) error
	Authorise(token string) (*model.User, error)
}

type userService struct {
	repo   postgres.Repository
	jwtKey string
}

func NewUserService(repo postgres.Repository, jtwKey string) *userService {
	return &userService{
		repo:   repo,
		jwtKey: jtwKey,
	}
}

func (s *userService) Authentication(user *model.User) (string, error) {

	// validate userName/Password
	if user.IsEmptyAuthenticationInfo() {
		return "", errors.New("invalid userName/password")
	}

	// query database user
	dbUser, err := s.repo.GetUser(user.UserName)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if err == gorm.ErrRecordNotFound {
		return "", errors.New("userName is invalid")
	}

	// hash and check raw password
	if s.hashPassword(user.Password, dbUser.Salt) != dbUser.Password {
		return "", errors.New("password is invalid")
	}

	// generate jwt
	token, err := s.createToken(user.UserName)
	if err != nil {
		return "", errors.New("there is an unknown error occur")
	}
	return token, nil

}

func (s *userService) CreateUser(user *model.User) error {

	// validate user
	if user.IsEmptyAuthenticationInfo() {
		return errors.New("invalid userName/password")
	}

	// double check user already created
	dbUser, _ := s.repo.GetUser(user.UserName)
	if dbUser != nil {
		return errors.New("UserName already Taken")
	}

	// hash password
	salt := uuid.New().String()
	hashedPw := s.hashPassword(user.Password, salt)
	user.Salt = salt
	user.Password = hashedPw

	// save user
	if err := s.repo.SaveUser(user); err != nil {
		return err
	}

	return nil

}

func (s *userService) Authorise(token string) (*model.User, error) {

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, errors.New("invalid user")
	}

	return user, nil
}

func (s *userService) hashPassword(rawPass, salt string) string {
	key, _ := hex.DecodeString(salt)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(rawPass))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *userService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
