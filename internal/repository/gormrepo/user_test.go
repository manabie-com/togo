package gormrepo

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"togo/internal/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_userRepository_Create_Failed(t *testing.T) {
	userInput := &domain.User{
		FullName: faker.Name(),
		Username: faker.Username(),
		Password: faker.Password(),
	}
	sqlErr := errors.New("insert failed")
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"users\"").WillReturnError(sqlErr)
	mock.ExpectRollback()
	r := userRepository{gdb}
	user, err := r.Create(context.Background(), userInput)
	assert.Nil(t, user)
	assert.ErrorIs(t, err, sqlErr)
}

func Test_userRepository_Create_Successful(t *testing.T) {
	fullName := faker.Name()
	username := faker.Username()
	password := faker.Password()
	taskPerDay := 1
	userInput := &domain.User{
		FullName:    fullName,
		Username:    username,
		Password:    password,
		TasksPerDay: taskPerDay,
	}
	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "password", "tasks_per_day"})
	rows.AddRow(1, fullName, username, password, taskPerDay)
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"users\"").WillReturnRows(rows)
	mock.ExpectCommit()
	r := userRepository{gdb}
	user, err := r.Create(context.Background(), userInput)
	assert.Nil(t, err)
	assert.Equal(t, &domain.User{
		ID:          1,
		FullName:    fullName,
		Username:    username,
		Password:    password,
		TasksPerDay: taskPerDay,
	}, user)
}

func Test_userRepository_FindOne_NotFound(t *testing.T) {
	filter := &domain.User{
		Username: faker.Username(),
	}
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "password", "tasks_per_day"})
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM")).WillReturnRows(rows)
	r := userRepository{gdb}
	user, err := r.FindOne(context.Background(), filter)
	assert.Nil(t, user)
	assert.ErrorIs(t, err, domain.ErrUserNotFound)
}

func Test_userRepository_FindOne_Successful(t *testing.T) {
	username := faker.Username()
	filter := &domain.User{
		Username: username,
	}
	eUser := &domain.User{
		ID:          1,
		FullName:    faker.Name(),
		Username:    username,
		Password:    faker.Password(),
		TasksPerDay: 2,
	}
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "password", "tasks_per_day"})
	rows.AddRow(eUser.ID, eUser.FullName, eUser.Username, eUser.Password, eUser.TasksPerDay)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM")).WillReturnRows(rows)
	r := userRepository{gdb}
	user, err := r.FindOne(context.Background(), filter)
	assert.Nil(t, err)
	assert.Equal(t, eUser, user)
}
