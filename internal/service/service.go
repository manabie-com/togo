package service

import (
	"main/config"
	"main/internal/store"
)

type TogoService struct {
	cfg   *config.Config
	store *store.Store
}

func NewService(cfg *config.Config, store *store.Store) (*TogoService, error) {
	return &TogoService{
		cfg:   cfg,
		store: store,
	}, nil
}

func (s *TogoService) CreateTodo() {

}
