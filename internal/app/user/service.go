package user

import (
	"context"
	"fmt"

	"github.com/dinhquockhanh/togo/internal/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, req *CreateUserReq) (*UserSafe, error)
	UpdateUserTier(ctx context.Context, req *UpdateUserTierReq) (*UserSafe, error)
	GetByUserName(ctx context.Context, req *GetUserByUserNameReq) (*User, error)
	List(ctx context.Context, req *ListUsersReq) ([]*UserSafe, error)
	Delete(ctx context.Context, req *DeleteUserByNameReq) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, req *CreateUserReq) (*UserSafe, error) {
	op := "service.Create"
	usr, err := s.GetByUserName(ctx, &GetUserByUserNameReq{
		UserName: req.UserName,
	})

	if err != nil {
		if !errors.IsSQLNotFound(err) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	if usr != nil {
		return nil, &errors.Error{
			Code:    409,
			Message: fmt.Sprintf("the username=%s already taken", req.UserName),
		}
	}

	a, err := s.repo.Create(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return a, nil
}

func (s *service) UpdateUserTier(ctx context.Context, req *UpdateUserTierReq) (*UserSafe, error) {
	a, err := s.repo.UpdateUserTier(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("service.UpdateUserTier: %w", err)
	}

	return a, nil
}

func (s *service) List(ctx context.Context, req *ListUsersReq) ([]*UserSafe, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetByUserName(ctx context.Context, req *GetUserByUserNameReq) (*User, error) {
	usr, err := s.repo.GetByUserName(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("service.GetByUserName: %w", err)
	}

	return usr, nil
}

func (s *service) Delete(ctx context.Context, req *DeleteUserByNameReq) error {
	err := s.repo.Delete(ctx, req)
	if err != nil {
		return fmt.Errorf("service.Delete: %w", err)
	}

	return nil
}
