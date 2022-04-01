package repository

import (
	"fmt"
	"github.com/ansidev/togo/domain/task"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

const DateFormat = "2006-01-02"

func TestPostgresTaskRepository(t *testing.T) {
	suite.Run(t, new(PostgresTaskRepositoryTestSuite))
}

type PostgresTaskRepositoryTestSuite struct {
	test.PostgresTestSuite
	repository task.ITaskRepository
}

func (s *PostgresTaskRepositoryTestSuite) SetupSuite() {
	s.PostgresTestSuite.SetupSuite()
	s.repository = NewPostgresTaskRepository(s.GormDb)
}

func (s *PostgresTaskRepositoryTestSuite) TestCreate_ShouldReturnCreatedTask() {
	userModel := s.insertUserIntoDb(5)

	taskModel := task.Task{
		Title:  "Task 1",
		UserID: userModel.ID,
	}

	t, err := s.repository.Create(taskModel, userModel)

	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(1), t.ID)
	require.Equal(s.T(), "Task 1", t.Title)
	require.Equal(s.T(), int64(1), t.UserID)
	require.False(s.T(), t.CreatedAt.IsZero())
	require.False(s.T(), t.UpdatedAt.IsZero())
}

func (s *PostgresTaskRepositoryTestSuite) TestGetTotalTasksByUserAndDate() {
	maxDailyTask := 5
	userModel := s.insertUserIntoDb(maxDailyTask)
	userId := userModel.ID
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	expectedTotalTasks := rand.Intn(maxDailyTask)

	for i := 0; i < expectedTotalTasks; i++ {
		s.insertTaskIntoDb(userId, date)
	}

	for j := 0; j < 10; j++ {
		s.insertTaskIntoDb(userId, date.AddDate(0, 0, 5))
	}

	total, err := s.repository.GetTotalTasksByUserAndDate(userModel, date)

	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(expectedTotalTasks), total)
}

func (s *PostgresTaskRepositoryTestSuite) insertUserIntoDb(maxDailyTask int) user.User {
	_, err := s.SqlDb.Exec(`INSERT INTO "user" (id, username, password, max_daily_task, created_at) VALUES ($1, $2, $3, $4, $5)`,
		1,
		"test_user",
		"$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		maxDailyTask,
		"2022-02-22 01:23:45")

	require.NoError(s.T(), err)

	var userModel user.User

	result := s.GormDb.First(&userModel, 1)
	require.NoError(s.T(), result.Error)

	return userModel
}

func (s *PostgresTaskRepositoryTestSuite) insertTaskIntoDb(userId int64, date time.Time) {
	_, err := s.SqlDb.Exec(`INSERT INTO "task" (title, user_id, created_at) VALUES ($1, $2, $3)`,
		fmt.Sprintf("Task %d", rand.Int()),
		userId,
		fmt.Sprintf("%s %02d:%02d:%02d", date.Format(DateFormat), rand.Intn(23), rand.Intn(59), rand.Intn(59)))

	require.NoError(s.T(), err)
}
