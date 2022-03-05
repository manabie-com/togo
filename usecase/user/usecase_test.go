package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	userRepo "github.com/khangjig/togo/repository/user"
	userUC "github.com/khangjig/togo/usecase/user"
	"github.com/khangjig/togo/util/jwt"
)

type ClaimTest string

type TestSuite struct {
	suite.Suite

	ctx       context.Context
	claimTest ClaimTest
	useCase   *userUC.UseCase

	mockUserRepo *userRepo.MockRepository
}

func (suite *TestSuite) SetupTest() {
	userClaims := jwt.DataClaim{
		UserID: 1,
	}

	suite.ctx = context.Background()
	suite.claimTest = "MyUserClaimTest"
	suite.ctx = context.WithValue(suite.ctx, suite.claimTest, &userClaims)

	suite.mockUserRepo = &userRepo.MockRepository{}

	suite.useCase = &userUC.UseCase{
		UserRepo: suite.mockUserRepo,
	}
}

func TestUseCaseUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
