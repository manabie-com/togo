package task

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestTaskRepositorySuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
}

func (s *TestTaskRepositorySuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository, _ = NewTaskRepository(s.DB)

}

func (s *TestTaskRepositorySuite) TearDownSuite() {

}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(TestTaskRepositorySuite))
}

func (s *TestTaskRepositorySuite) TestTaskRepository_AddTask() {
	var (
		userID  uint64 = 1
		content string = "example"
	)

	const sqlSelectAll = `INSERT INTO "tasks" ("content","status","user_id") VALUES ($1,$2,$3)`

	rows := sqlmock.NewRows(nil)

	s.mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
		WithArgs(content, StatusActive, userID).
		WillReturnRows(rows)

	task, err := s.repository.AddTask(userID, content)
	require.NoError(s.T(), err)
	if assert.NotNil(s.T(), task) {
		assert.Equal(s.T(), userID, task.UserID)
		assert.Equal(s.T(), content, task.Content)
	}
}

func (s *TestTaskRepositorySuite) TestTaskRepository_AddManyTasks() {

	var (
		userID   uint64 = 1
		contents        = []string{
			"content 01",
			"content 02",
		}
	)

	myMockRows1 := sqlmock.NewRows([]string{"created_date", "id"}).AddRow(time.Now(), 1).AddRow(time.Now(), 2)

	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks"`)).WillReturnRows(myMockRows1)

	l, err := s.repository.AddManyTasks(userID, contents)
	require.NoError(s.T(), err)

	if assert.NotNil(s.T(), l) {
		require.Len(s.T(), l, 2)
	}
}

func (s *TestTaskRepositorySuite) TestTaskRepository_RetrieveTasks() {

	var (
		userID      uint64 = 1
		createdDate string = time.Now().Format("2006-01-02")
	)

	const sqlSelectAll = `SELECT * FROM "tasks" WHERE user_id = $1 AND to_char(created_date,'YYYY-MM-DD') = $2`

	rows := sqlmock.NewRows([]string{"id", "content", "status", "user_id", "created_date"}).
		AddRow(1, "content 01", StatusActive, userID, time.Now()).
		AddRow(2, "content 02", StatusActive, userID, time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WithArgs(userID, createdDate).WillReturnRows(rows)

	l, err := s.repository.RetrieveTasks(userID, createdDate)
	require.NoError(s.T(), err)
	if assert.NotNil(s.T(), l) {
		require.Len(s.T(), l, 2)
	}
}

func (s *TestTaskRepositorySuite) TestTaskRepository_NumTasksToday() {
	var (
		userID        uint64 = 1
		createdDate   string = time.Now().Format("2006-01-02")
		numTasksToday int64  = 3
	)

	const sqlSelectAll = `SELECT count(*) FROM tasks WHERE status = $1 AND user_id = $2 AND to_char(created_date,'YYYY-MM-DD') = $3`

	rows := sqlmock.NewRows([]string{"count"}).AddRow(numTasksToday)

	s.mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WithArgs(StatusActive, userID, createdDate).
		WillReturnRows(rows)

	l, err := s.repository.NumTasksToday(userID)
	require.NoError(s.T(), err)
	if assert.NotNil(s.T(), l) {
		assert.Equal(s.T(), l, numTasksToday)
	}
}
