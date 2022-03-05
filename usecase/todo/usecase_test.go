package todo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/khangjig/togo/model"
	todoRepo "github.com/khangjig/togo/repository/todo"
	userRepo "github.com/khangjig/togo/repository/user"
	todoUC "github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/util/jwt"
)

type TestSuite struct {
	suite.Suite

	ctx        context.Context
	userClaims *model.User
	useCase    *todoUC.UseCase

	mockUserRepo      *userRepo.MockRepository
	mockUserCacheRepo *userRepo.MockCacheRepository
	mockTodoRepo      *todoRepo.MockRepository
	mockTodoCacheRepo *todoRepo.MockCacheRepository
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
	suite.mockUserCacheRepo = &userRepo.MockCacheRepository{}
	suite.mockTodoRepo = &todoRepo.MockRepository{}
	suite.mockTodoCacheRepo = &todoRepo.MockCacheRepository{}

	suite.useCase = &todoUC.UseCase{
		UserRepo:      suite.mockUserRepo,
		UserCacheRepo: suite.mockUserCacheRepo,
		TodoRepo:      suite.mockTodoRepo,
		TodoCacheRepo: suite.mockTodoCacheRepo,
	}
}

func TestUseCaseUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
