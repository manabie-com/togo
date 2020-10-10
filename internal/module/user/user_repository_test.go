package user

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository Repository
}

func (s *Suite) SetupSuite() {
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

	s.repository, _ = NewUserRepository(s.DB)

}

func (s *Suite) TearDownSuite() {

}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestUserRepository_GetUser() {

	var (
		id    uint64 = 1
		email string = "test01@mail.com"
	)

	const sqlSelectAll = (`SELECT * FROM "users" WHERE id = $1`)

	rows := sqlmock.NewRows([]string{"id", "email", "max_todo"}).
		AddRow(id, "test01@mail.com", 5)

	s.mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WithArgs(id).
		WillReturnRows(rows)

	user, err := s.repository.GetUser(id)
	require.NoError(s.T(), err)
	if assert.NotNil(s.T(), user) {
		assert.Equal(s.T(), email, user.Email)
	}
}

func (s *Suite) TestUserRepository_GetAll() {
	const sqlSelectAll = `SELECT * FROM "users"`

	rows := sqlmock.NewRows([]string{"id", "email", "max_todo"}).
		AddRow(1, "test01@mail.com", 5).
		AddRow(2, "test02@mail.com", 5)

	s.mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
		WillReturnRows(rows)

	l, err := s.repository.GetAll()
	require.NoError(s.T(), err)
	if assert.NotNil(s.T(), l) {
		require.Len(s.T(), l, 2)
	}
}
