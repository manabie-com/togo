package userstorage

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	store *sqlStore
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

	s.store = NewSQLStore(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_store_FindUser() {
	var (
		id    = 1
		email = "abc@mail.com"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow(id, email))

	res, err := s.store.FindUser(context.TODO(), map[string]interface{}{"id": 1})

	require.NoError(s.T(), err)
	require.Equal(s.T(), res.ID, id)
	require.Equal(s.T(), res.Email, email)
}
