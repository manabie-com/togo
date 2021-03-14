package services

import (
	"context"
	"fmt"
	"github.com/banhquocdanh/togo/internal/storages"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/google/uuid"
)

var Now = time.Now

// ToDoService implement HTTserver
type ToDoService struct {
	Store storages.StoreInterface
}

type Option func(service *ToDoService)

func NewToDoService(opts ...Option) *ToDoService {
	srv := &ToDoService{}
	for _, opt := range opts {
		opt(srv)
	}

	return srv

}

func WithStore(db storages.StoreInterface) Option {
	return func(srv *ToDoService) {
		srv.Store = db
	}
}

func (s *ToDoService) ListTasks(ctx context.Context, userID, createDate string) ([]*storages.Task, error) {
	if userID == "" {
		return nil, fmt.Errorf("user id invalid")
	}
	if createDate == "" {
		return nil, fmt.Errorf("created date invalid")
	}

	return s.Store.RetrieveTasks(
		ctx,
		userID,
		createDate,
	)
}

func (s *ToDoService) AddTask(ctx context.Context, userID, content string) (*storages.Task, error) {
	if userID == "" {
		return nil, fmt.Errorf("invalid userID")
	}
	if content == "" {
		return nil, fmt.Errorf("invalid task's content")
	}

	now := Now()
	t := &storages.Task{
		ID:          uuid.New().String(),
		Content:     content,
		UserID:      userID,
		CreatedDate: now.Format("2006-01-02"),
	}

	return t, s.Store.AddTask(ctx, t)
}

func (s *ToDoService) validateUser(ctx context.Context, userID, pw string) bool {
	if userID == "" || pw == "" {
		return false
	}

	return s.Store.ValidateUser(ctx, userID, pw)
}

func (s *ToDoService) Login(ctx context.Context, userID, pw string, jwtKey string) (string, error) {
	if userID == "" || pw == "" {
		return "", fmt.Errorf("invalid user/pw")
	}

	if !s.validateUser(ctx, userID, pw) {
		return "", fmt.Errorf("user/pw is incorrect")
	}

	return s.createToken(userID, jwtKey)
}

func (s *ToDoService) createToken(id string, jwtKey string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) ValidToken(token, jwtKey string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is empty")
	}
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", fmt.Errorf("invalid token")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("not found userID")
	}

	exp, ok := claims["exp"].(int64)
	if !ok {
		return "", fmt.Errorf("not found expired time")
	}

	if Now().Unix() > exp {
		return "", fmt.Errorf("token is expired")
	}

	return id, nil
}
