package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/utils"
)

var (
	username      = "nohattee"
	password      = "1qaz@WSX"
	wrongUsername = "wrong_username"
	wrongPassword = "wrong_password"
)

func setupMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestValidateUser(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery("SELECT id FROM users").
		WithArgs(username, password).
		WillReturnRows(rows)

	repository := postgres.NewPostgresRepository(db)
	usecase := NewUserUseCase(repository)

	expected := true
	actual := usecase.ValidateUser(
		context.TODO(),
		utils.SqlString(username),
		utils.SqlString(password),
	)

	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestValidateUserWithWrongUsername(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	mock.ExpectQuery("SELECT id FROM users WHERE username = (.+) AND password = (.+)").
		WillReturnError(fmt.Errorf("the username is wrong"))

	repository := postgres.NewPostgresRepository(db)
	usecase := NewUserUseCase(repository)

	expected := false
	actual := usecase.ValidateUser(
		context.TODO(),
		utils.SqlString(wrongUsername),
		utils.SqlString(password),
	)

	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestValidateUserWithWrongPassword(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	mock.ExpectQuery("SELECT id FROM users WHERE username = (.+) AND password = (.+)").
		WillReturnError(fmt.Errorf("the password is wrong"))

	repository := postgres.NewPostgresRepository(db)
	usecase := NewUserUseCase(repository)
	expected := false
	actual := usecase.ValidateUser(
		context.TODO(),
		utils.SqlString(username),
		utils.SqlString(wrongPassword),
	)

	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByUsername(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "max_todo"}).
		AddRow(1, username, 5)

	mock.ExpectQuery("SELECT id, username, max_todo FROM users").
		WithArgs(username).
		WillReturnRows(rows)

	repository := postgres.NewPostgresRepository(db)
	usecase := NewUserUseCase(repository)

	expected := &storages.User{
		ID:       1,
		Username: username,
		MaxTodo:  5,
	}
	actual, err := usecase.GetUserByUsername(
		context.TODO(),
		utils.SqlString(username),
	)
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Not expected")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
