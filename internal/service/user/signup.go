package userservice

import (
	"context"

	"github.com/trinhdaiphuc/togo/internal/entities"
)

func (u *userService) SignUp(ctx context.Context, user *entities.User) (*entities.User, error) {
	return u.userRepo.CreateUser(ctx, user)
}
