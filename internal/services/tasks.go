package services

import (
	"context"
	"fmt"
	"github.com/banhquocdanh/togo/internal/storages"
	"time"

	"github.com/google/uuid"
)

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
	now := time.Now()
	t := &storages.Task{
		ID:          uuid.New().String(),
		Content:     content,
		UserID:      userID,
		CreatedDate: now.Format("2006-01-02"),
	}

	return t, s.Store.AddTask(ctx, t)
}

func (s *ToDoService) ValidateUser(ctx context.Context, userID, pw string) bool {
	if userID == "" || pw == "" {
		return false
	}

	return s.Store.ValidateUser(ctx, userID, pw)
}
