package service

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
)

type jwtGenerator interface {
  Token(user *core.User) (string, error)
}

type JwtUserService struct {
  UserRepo core.UserRepo
  JwtGenerator jwtGenerator
}

func (service *JwtUserService) Login(ctx context.Context, id, password string) (string, error) {
  user, err := service.UserRepo.Validate(ctx, id, password)
  if err != nil {
    return "", err
  }
  return service.JwtGenerator.Token(user)
}

func (service *JwtUserService) Signup(ctx context.Context, user *core.User, password string) (string, error) {
  if user.ID == "" {
    return "", core.ErrEmptyId
  }
  if password == "" {
    return "", core.ErrEmptyPassword
  }
  if user.MaxTodo <= 0 {
    return "", core.ErrInvalidMaxTodo
  }
  err := service.UserRepo.Create(ctx, user, password)
  if err != nil {
    return "", err
  }
  return service.JwtGenerator.Token(user)
}
