package repository

import (
	"ansidev.xyz/pkg/tm"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPostgresAuthorRepository(t *testing.T) {
	suite.Run(t, new(PostgresUserRepositoryTestSuite))
}

type PostgresUserRepositoryTestSuite struct {
	test.PostgresTestSuite
	repository user.IUserRepository
}

func (s *PostgresUserRepositoryTestSuite) SetupSuite() {
	s.PostgresTestSuite.SetupSuite()
	s.repository = NewPostgresUserRepository(s.GormDb)
}

func (s *PostgresUserRepositoryTestSuite) TestFindFirstByID_ShouldReturnNoRecord() {
	_, err := s.repository.FindFirstByID(1)

	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrRecordNotFound, errors.Cause(err))
}

func (s *PostgresUserRepositoryTestSuite) TestFindFirstByID_ShouldReturnRecord() {
	u := s.insertUserIntoDb()

	userModel, err2 := s.repository.FindFirstByID(u.ID)

	require.NoError(s.T(), err2)
	require.Equal(s.T(), userModel.ID, u.ID)
	require.Equal(s.T(), userModel.Username, u.Username)
	require.Equal(s.T(), userModel.Password, u.Password)
	require.Equal(s.T(), userModel.MaxDailyTask, u.MaxDailyTask)
	require.Equal(s.T(), userModel.CreatedAt.Format(tm.DateTimeFormat), u.CreatedAt.Format(tm.DateTimeFormat))
	require.Equal(s.T(), userModel.UpdatedAt.Format(tm.DateTimeFormat), u.UpdatedAt.Format(tm.DateTimeFormat))
}

func (s *PostgresUserRepositoryTestSuite) TestFindFirstByUsername_ShouldReturnNoRecord() {
	_, err := s.repository.FindFirstByUsername("test_user")

	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrRecordNotFound, errors.Cause(err))
}

func (s *PostgresUserRepositoryTestSuite) TestFindFirstByUsername_ShouldReturnRecord() {
	u := s.insertUserIntoDb()

	userModel, err2 := s.repository.FindFirstByUsername(u.Username)

	require.NoError(s.T(), err2)
	require.Equal(s.T(), userModel.ID, u.ID)
	require.Equal(s.T(), userModel.Username, u.Username)
	require.Equal(s.T(), userModel.Password, u.Password)
	require.Equal(s.T(), userModel.MaxDailyTask, u.MaxDailyTask)
	require.Equal(s.T(), userModel.CreatedAt.Format(tm.DateTimeFormat), u.CreatedAt.Format(tm.DateTimeFormat))
	require.Equal(s.T(), userModel.UpdatedAt.Format(tm.DateTimeFormat), u.UpdatedAt.Format(tm.DateTimeFormat))
}

func (s *PostgresUserRepositoryTestSuite) insertUserIntoDb() user.User {
	_, err := s.SqlDb.Exec(`INSERT INTO "user" (id, username, password, max_daily_task, created_at) VALUES ($1, $2, $3, $4, $5)`,
		1,
		"test_user",
		"$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		5,
		"2022-02-22 01:23:45")

	require.NoError(s.T(), err)

	var userModel user.User

	result := s.GormDb.First(&userModel, 1)
	require.NoError(s.T(), result.Error)

	return userModel
}
