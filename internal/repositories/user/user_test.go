package user

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/require"
)

func Test_repository_GetByUsername(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()

	userRepository := NewUserRepository(db)

	{ //Success case
		input := &models.User{
			Username: "manabie_1",
			Password: "manabie_1",
		}
		// Create mock data for test
		mock.ExpectExec("INSERT INTO users").
			WithArgs(1, input.Username, input.Password, 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err := recordStats(db, input.Username, input.Password)
		require.Nil(t, err)

		rows := sqlmock.NewRows([]string{"id", "username", "password", "max_task_per_day"}).
			AddRow("1", input.Username, input.Password, 5)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnRows(rows)

		user, err := userRepository.GetByUsername(input.Username)
		require.Nil(t, err)
		require.NotNil(t, user)
		require.Equal(t, input.Username, user.Username)
		require.Equal(t, input.Password, user.Password)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}
	{ // Fail case
		input := &models.User{
			Username: "manabie_2",
			Password: "manabie_2",
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnError(fmt.Errorf("invalid username"))

		user, err := userRepository.GetByUsername(input.Username)
		require.NotNil(t, err)
		require.Nil(t, user)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}
}

func Test_repository_Create(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()

	userRepository := NewUserRepository(db)

	{ //Success case
		user := New(&models.User{
			Username:      "manabie_1",
			Password:      "manabie_1",
			MaxTaskPerDay: 5,
		})
		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("id","username","password","max_task_per_day") VALUES ($1,$2,$3,$4) RETURNING "users"."id"`)).
			WithArgs(user.ID, user.Username, user.Password, user.MaxTaskPerDay).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := userRepository.Create(user)
		require.Nil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}

	{
		//Fail case
		user := New(&models.User{
			Username:      "manabie_1",
			Password:      "manabie_1",
			MaxTaskPerDay: 5,
		})
		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("id","username","password","max_task_per_day") VALUES ($1,$2,$3,$4) RETURNING "users"."id"`)).
			WithArgs("1", user.Username, user.Password, user.MaxTaskPerDay).
			WillReturnError(fmt.Errorf("nil user_id"))
		mock.ExpectCommit()

		err := userRepository.Create(user)
		require.NotNil(t, err)
		require.NotNil(t, mock.ExpectationsWereMet())
	}
}
