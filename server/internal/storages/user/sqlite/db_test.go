package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	userEntity "github.com/HoangVyDuong/togo/internal/storages/user"
	"reflect"
	"testing"
)

var mockUser = userEntity.User{
	ID: 1234,
	Name: "vydh2",
	Password: "abc",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("error happen when new mock")
	}

	return db, mock
}

func TestGetUserByName(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, password FROM user WHERE name = \\?"

	rows := sqlmock.NewRows([]string{"id", "name", "password"}).
		AddRow(mockUser.ID, mockUser.Name, mockUser.Password)

	mock.ExpectQuery(query).WithArgs(mockUser.Name).WillReturnRows(rows)

	user, err := repo.GetUserByName(context.Background(), mockUser.Name)
	if err != nil {
		t.Errorf("Data Error")
	}
	if !reflect.DeepEqual(user, mockUser) {
		t.Errorf("")
	}
}

func TestGetUserByNameErr(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, password FROM user WHERE name = \\?"

	mock.ExpectQuery(query).WithArgs(mockUser.Name).WillReturnError(errors.New("Database error"))

	_, err := repo.GetUserByName(context.Background(), mockUser.Name)
	if err == nil {
		t.Errorf("Data Error")
	}
}