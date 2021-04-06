package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/entities"
	"github.com/manabie-com/togo/internal/util"
)

type ToDoUsecase struct {
	Store storages.Store
}

func NewToDoUsecase(store storages.Store) *ToDoUsecase {
	return &ToDoUsecase{
		Store: store,
	}
}

// GetToken is function validate user and password then create and return token
func (s *ToDoUsecase) GetToken(id, pass string) (string, error) {
	if !s.Store.ValidateUser(context.Background(), id, pass) {
		return "", errors.New("incorrect user_id/pwd")
	}

	token, err := s.createToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}

// createToken is function create token with value is user's id
func (s *ToDoUsecase) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(util.Conf.SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoUsecase) ValidToken(token string) (string, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(util.Conf.SecretKey), nil
	})
	if err != nil {
		return "", false
	}

	if !t.Valid {
		return "", false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", false
	}

	return id, true
}

// AddTask ...
func (s *ToDoUsecase) AddTask(content, userId string) error {
	createdDate := util.GetDate()
	task := &entities.Task{
		ID:          uuid.New().String(),
		Content:     content,
		UserID:      userId,
		CreatedDate: createdDate,
	}

	return s.Store.AddTask(context.Background(), task)
}

// ListTask ...
func (s *ToDoUsecase) ListTask(createdDate, userId string, total, page int) ([]*entities.Task, error) {
	if total > 10 {
		total = 10
	}
	return s.Store.RetrieveTasks(context.Background(), userId, createdDate, total, page)
}
