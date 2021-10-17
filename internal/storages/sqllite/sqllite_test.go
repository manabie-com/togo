package sqllite

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
)

var (
	username      = "nohattee"
	password      = "1qaz@WSX"
	wrongUsername = "wrong_username"
	wrongPassword = "wrong_password"
	userID        = 1
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

	repository := NewLiteRepository(db)
	expected := true
	actual := repository.ValidateUser(
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

	mock.ExpectQuery("SELECT id FROM users").
		WillReturnError(fmt.Errorf("the username is wrong"))

	repository := NewLiteRepository(db)
	expected := false
	actual := repository.ValidateUser(
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

	mock.ExpectQuery("SELECT id FROM users").
		WillReturnError(fmt.Errorf("the password is wrong"))

	repository := NewLiteRepository(db)
	expected := false
	actual := repository.ValidateUser(
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

func TestRetrieveTasks(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)
	listTaskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(listTaskIds[0], "This is just the first test", 1, currentDate).
		AddRow(listTaskIds[1], "This is just the second test", 1, currentDate).
		AddRow(listTaskIds[2], "This is just the third test", 1, currentDate)

	mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = (.+) AND created_date = (.+)").
		WillReturnRows(rows)

	repository := NewLiteRepository(db)
	expected := []*storages.Task{
		{
			ID:          listTaskIds[0],
			Content:     "This is just the first test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{

			ID:          listTaskIds[1],
			Content:     "This is just the second test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{
			ID:          listTaskIds[2],
			Content:     "This is just the third test",
			UserID:      1,
			CreatedDate: currentDate,
		},
	}
	actual, err := repository.RetrieveTasks(
		context.TODO(),
		uint(userID),
		utils.SqlString(currentDate),
	)
	if err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Not expected")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddTask(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)

	task := storages.Task{
		ID:          uuid.New().String(),
		Content:     "This is just for the test",
		UserID:      1,
		CreatedDate: currentDate,
	}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs(task.ID, task.Content, task.UserID, task.CreatedDate).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewLiteRepository(db)
	err := repository.AddTask(
		context.TODO(),
		&task,
	)
	if err != nil {
		t.Error("Not expected")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
