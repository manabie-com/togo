package auth_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/khangjig/togo/config"
	userRepo "github.com/khangjig/togo/repository/user"
	authUC "github.com/khangjig/togo/usecase/auth"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase *authUC.UseCase

	mockUserRepo *userRepo.MockRepository
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.mockUserRepo = &userRepo.MockRepository{}

	suite.useCase = &authUC.UseCase{
		UserRepo: suite.mockUserRepo,
		Config:   &config.Config{},
	}
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
