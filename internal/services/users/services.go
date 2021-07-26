package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
)

type ToDoService struct {
	JWTKey string
	repo   storages.Repository
}

func SetupNewService(jwtKey string, r storages.Repository) *ToDoService {
	return &ToDoService{JWTKey: jwtKey, repo: r}
}

func (s *ToDoService) GetAuthToken(username sql.NullString, password sql.NullString) (string, error) {
	user, err := s.repo.ValidateUser(username, password)
	if err != nil {
		return "", err
	}

	return s.createToken(user.ID)
}

func (s *ToDoService) FindUserById(context context.Context, id sql.NullString) bool {
	_, err := s.repo.GetUserById(context, id)
	if err != nil {
		log.Println(fmt.Sprintf("Cannot find user by id %s", id.String))
	}
	return err == nil
}

func (s *ToDoService) createToken(id string) (string, error) {
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
