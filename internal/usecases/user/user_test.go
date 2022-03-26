package user

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository/postgres"
	"github.com/manabie-com/togo/internal/utils"
)

func TestValidateUser_Success(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT id FROM users").
		WithArgs(username, password).
		WillReturnRows(rows)

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUsecase := NewUserUsecase(repository)
	expect := true
	actual := userUsecase.ValidateUser(
		context.Background(),
		utils.SqlNullString(username),
		utils.SqlNullString(password),
	)

	if actual != expect {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestValidateUser_WrongUsername(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	mock.ExpectQuery("SELECT id FROM users").
		WithArgs(wrongUsername, password).
		WillReturnError(fmt.Errorf("invalid username"))

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUseCase := NewUserUsecase(repository)
	expect := false
	actual := userUseCase.ValidateUser(
		context.Background(),
		utils.SqlNullString(wrongUsername),
		utils.SqlNullString(password),
	)

	if actual != expect {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestValidateUser_WrongPassword(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	mock.ExpectQuery("SELECT id FROM users").
		WithArgs(username, wrongPassword).
		WillReturnError(fmt.Errorf("invalid password"))

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUsecase := NewUserUsecase(repository)
	expect := false
	actual := userUsecase.ValidateUser(
		context.Background(),
		utils.SqlNullString(username),
		utils.SqlNullString(wrongPassword),
	)

	if expect != actual {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestGetUserByUsername_Success(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"id", "username", "max_task_per_day"}).
		AddRow(1, username, 5)

	mock.ExpectQuery("SELECT id, username, max_task_per_day FROM users").
		WithArgs(username).
		WillReturnRows(rows)

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUsecase := NewUserUsecase(repository)

	expect := &models.User{
		ID:            1,
		Username:      username,
		MaxTaskPerDay: 5,
	}
	actual, err := userUsecase.GetUserByUserName(context.Background(), utils.SqlNullString(username))
	if err != nil {
		t.Errorf("fail get user by username")
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect: %#v, actual: %#v", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestGetUserByUserName_FailWithNoRow(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	mock.ExpectQuery("SELECT id, username, max_task_per_day FROM users").
		WithArgs(wrongUsername).
		WillReturnError(errors.New("sql: no rows in result set"))

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUsecase := NewUserUsecase(repository)

	user, err := userUsecase.GetUserByUserName(context.Background(), utils.SqlNullString(wrongUsername))
	if err == nil || user != nil {
		t.Fatal("expect 'sql: no rows in result set' and nil user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestGenerateToken_Success(t *testing.T) {
	db, _ := setupMock()
	defer db.Close()

	// set env for test
	{
		os.Setenv("JWT_TIMEOUT", "60")
		os.Setenv("JWT_KEY", "alk123@!opjO")
	}

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUsecase := NewUserUsecase(repository)

	token, err := userUsecase.GenerateToken(uint(1), uint(5))
	if err != nil || token == "" {
		t.Errorf("fail generate token")
	}

	// test passed, unset env
	{
		os.Unsetenv("JWT_TIMEOUT")
		os.Unsetenv("JWT_KEY")
	}
}

func TestGenerateToken_Fail(t *testing.T) {
	db, _ := setupMock()
	defer db.Close()

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	userUsecase := NewUserUsecase(repository)

	token, err := userUsecase.GenerateToken(uint(1), uint(5))
	if err == nil || token != "" {
		t.Errorf("expect error not nil and token is empty")
	}
}
