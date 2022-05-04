package userservice

import (
	"context"
	"todo/internal/entities"
)

func (u *userService) SignUp(ctx context.Context, user *entities.User) (*entities.User, error) {
	return u.userRepo.CreateUser(ctx, user)
}
