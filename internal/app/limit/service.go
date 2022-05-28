package limit

import (
	"context"
	"fmt"
)

type Repository interface {
	GetLimit(ctx context.Context, req *GetLimitReq) (*Limit, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetLimit(ctx context.Context, req *GetLimitReq) (*Limit, error) {
	limit, err := s.repo.GetLimit(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("service.GetTaskByID: %w", err)
	}

	return limit, nil
}
