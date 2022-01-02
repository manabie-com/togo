package user

import (
	"github.com/labstack/echo/v4"

	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"

	utilsToken "github.com/manabie-com/togo/app/utils/token"
)

// Service return list apis relate to flow
type Service interface {
	Create(c echo.Context) error
	Login(c echo.Context) error
}

type service struct {
	userRepo   userRepo.Repository
	tokenMaker utilsToken.Maker
}

// NewService return handler instance
func NewService(
	userRepoInstance userRepo.Repository,
	tokenMaker utilsToken.Maker,
) Service {
	return &service{
		userRepo:   userRepoInstance,
		tokenMaker: tokenMaker,
	}
}
