package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/khangjig/togo/model"
	userRepo "github.com/khangjig/togo/repository/user"
	userUC "github.com/khangjig/togo/usecase/user"
	"github.com/khangjig/togo/util/jwt"
)

type TestSuite struct {
	suite.Suite

	ctx        context.Context
	userClaims *model.User
	useCase    *userUC.UseCase

	mockUserRepo *userRepo.MockRepository
}

func (suite *TestSuite) SetupTest() {
	suite.userClaims = &model.User{
		ID:      1,
		MaxTodo: 10,
	}

	suite.ctx = context.Background()
	// nolint:staticcheck
	suite.ctx = context.WithValue(suite.ctx, jwt.MyUserClaim, suite.userClaims)

	suite.mockUserRepo = &userRepo.MockRepository{}

	suite.useCase = &userUC.UseCase{
		UserRepo: suite.mockUserRepo,
	}
}

func TestUseCaseUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
