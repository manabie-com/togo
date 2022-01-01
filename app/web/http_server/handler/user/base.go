package user

import (
	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"

	"github.com/labstack/echo/v4"
)

// Service return list apis relate to flow
type Service interface {
	Create(c echo.Context) error
}

type service struct {
	userRepo userRepo.Repository
}

// NewService return handler instance
func NewService(
	userRepoInstance userRepo.Repository,
) Service {
	return &service{
		userRepo: userRepoInstance,
	}
}
