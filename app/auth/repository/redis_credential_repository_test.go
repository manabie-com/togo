package repository

import (
	"context"
	"encoding/json"
	"github.com/ansidev/togo/domain/auth"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestRedisCredentialRepository(t *testing.T) {
	suite.Run(t, new(RedisCredentialRepositoryTestSuite))
}

type RedisCredentialRepositoryTestSuite struct {
	test.RedisTestSuite
	tokenTTL   time.Duration
	repository auth.ICredRepository
}

func (s *RedisCredentialRepositoryTestSuite) SetupSuite() {
	s.RedisTestSuite.SetupSuite()
	s.tokenTTL = 24 * time.Hour
	s.repository = NewRedisCredentialRepository(s.Rdb, s.tokenTTL)
}

func (s *RedisCredentialRepositoryTestSuite) TestSave_ShouldSave() {
	userModel := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		MaxDailyTask: 5,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
	token, err := s.repository.Save(userModel)

	require.NoError(s.T(), err)

	_, err1 := uuid.Parse(token)
	require.NoError(s.T(), err1)
	require.True(s.T(), len(token) > 0)
}

func (s *RedisCredentialRepositoryTestSuite) TestGet_ShouldReturnAuth() {
	expectedAuthModel := auth.AuthenticationCredential{
		ID:           1,
		MaxDailyTask: 5,
	}

	bytes, err1 := json.Marshal(expectedAuthModel)
	require.NoError(s.T(), err1)

	token := uuid.NewString()
	cmd := s.RedisClient.Set(context.Background(), token, bytes, s.tokenTTL)
	_, err2 := cmd.Result()
	require.NoError(s.T(), err2)

	authModel, err3 := s.repository.Get(token)
	require.NoError(s.T(), err3)
	require.Equal(s.T(), expectedAuthModel.ID, authModel.ID)
	require.Equal(s.T(), expectedAuthModel.MaxDailyTask, authModel.MaxDailyTask)
}
