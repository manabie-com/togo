package grpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	todoRepo "github.com/khangjig/togo/repository/todo"
	userRepo "github.com/khangjig/togo/repository/user"
	grpcUC "github.com/khangjig/togo/usecase/grpc"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase *grpcUC.UseCase

	mockUserRepo      *userRepo.MockRepository
	mockUserCacheRepo *userRepo.MockCacheRepository
	mockTodoRepo      *todoRepo.MockRepository
	mockTodoCacheRepo *todoRepo.MockCacheRepository
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.mockUserRepo = &userRepo.MockRepository{}
	suite.mockUserCacheRepo = &userRepo.MockCacheRepository{}
	suite.mockTodoRepo = &todoRepo.MockRepository{}
	suite.mockTodoCacheRepo = &todoRepo.MockCacheRepository{}

	suite.useCase = &grpcUC.UseCase{
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
