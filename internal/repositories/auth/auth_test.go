package auth

import (
	"fmt"
	"regexp"
	"testing"

	"example.com/m/v2/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func Test_repository_ValidateUser(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()

	type userValidate struct {
		Username string
		Password string
		Expect   bool
	}

	authRepository := NewAuthRepository(db)

	{ //Success case
		user := userValidate{
			Username: "manabie321",
			Password: "123456",
			Expect:   true,
		}

		// Create mock data for test
		mock.ExpectExec("INSERT INTO users").
			WithArgs(1, user.Username, user.Password, 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err := recordStats(db, user.Username, user.Password)
		require.Nil(t, err)

		rows := sqlmock.NewRows([]string{"id", "username", "password", "max_task_per_day"}).
			AddRow("1", user.Username, user.Password, 5)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(user.Username).
			WillReturnRows(rows)

		actual, err := authRepository.ValidateUser(user.Username)
		require.Nil(t, err)
		require.Equal(t, user.Expect, actual)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}
	{ // Fail case
		user := userValidate{
			Username: "fail_name",
			Password: "fail_pass",
			Expect:   false,
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(user.Username).
			WillReturnError(fmt.Errorf("invalid user"))

		actual, err := authRepository.ValidateUser(user.Username)
		require.NotNil(t, err)
		require.Equal(t, user.Expect, actual)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}
}

func Test_repository_GenerateToken(t *testing.T) {
	t.Parallel()
	//Load Env
	utils.LoadEnv("../../../.env")

	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	type args struct {
		userID        string
		maxTaskPerday string
	}

	authRepository := NewAuthRepository(db)
	{ // Success case
		input := args{
			userID:        "1",
			maxTaskPerday: "5",
		}
		tokenStr, err := authRepository.GenerateToken(input.userID, input.maxTaskPerday)
		require.Nil(t, err)
		require.NotNil(t, tokenStr)

		require.NotEqual(t, "", utils.SafeString(tokenStr))
	}

	{ // Fail case
		input := args{
			userID:        "",
			maxTaskPerday: "",
		}
		tokenStr, err := authRepository.GenerateToken(input.userID, input.maxTaskPerday)
		require.NotNil(t, err)
		require.Nil(t, tokenStr)

		require.Equal(t, "", utils.SafeString(tokenStr))
	}
}
