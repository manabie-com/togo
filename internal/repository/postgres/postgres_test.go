package postgres

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestValidateUser_Success(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT id FROM users").
		WithArgs(username, password).
		WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	expect := true
	actual := repository.ValidateUser(
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

	repository := NewPostgresRepository(db)
	expect := false
	actual := repository.ValidateUser(
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
		WillReturnError(fmt.Errorf("invalid pasword"))

	repository := NewPostgresRepository(db)
	expect := false
	actual := repository.ValidateUser(
		context.Background(),
		utils.SqlNullString(username),
		utils.SqlNullString(wrongPassword),
	)

	if actual != expect {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestGetUser_Success(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "username", "max_task_per_day"}).
		AddRow(uint(userID), username, 5)

	mock.ExpectQuery("SELECT id, username, max_task_per_day FROM users").
		WithArgs(username).
		WillReturnRows(row)

	repository := NewPostgresRepository(db)
	expect := &models.User{
		ID:            uint(userID),
		Username:      username,
		MaxTaskPerDay: 5,
	}

	actual, err := repository.GetUserByUserName(context.Background(), utils.SqlNullString(username))
	if err != nil {
		t.Fatalf("failed get user by username: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("actual does not match expect")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestGetUser_FailWithNoRow(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	mock.ExpectQuery("SELECT id, username, max_task_per_day FROM users").
		WithArgs(wrongUsername).
		WillReturnError(errors.New("sql: no rows in result set"))

	repository := NewPostgresRepository(db)

	user, err := repository.GetUserByUserName(context.Background(), utils.SqlNullString(wrongUsername))
	if err == nil || user != nil {
		t.Fatal("expect 'sql: no rows in result set' and nil user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestRetrieveTasks_Success(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")
	tasksIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
	rows := sqlmock.NewRows([]string{"id", "detail", "user_id", "created_date"}).
		AddRow(tasksIds[0], "first task", uint(userID), currentDate).
		AddRow(tasksIds[1], "second task", uint(userID), currentDate).
		AddRow(tasksIds[2], "third task", uint(userID), currentDate)

	mock.ExpectQuery("SELECT id, detail, user_id, created_date FROM tasks").
		WithArgs(uint(userID), currentDate).
		WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	expect := []*models.Task{
		{
			ID:          tasksIds[0],
			Detail:      "first task",
			UserID:      uint(userID),
			CreatedDate: currentDate,
		},
		{
			ID:          tasksIds[1],
			Detail:      "second task",
			UserID:      uint(userID),
			CreatedDate: currentDate,
		},
		{
			ID:          tasksIds[2],
			Detail:      "third task",
			UserID:      uint(userID),
			CreatedDate: currentDate,
		},
	}

	actual, err := repository.RetrieveTasks(context.Background(), uint(userID), utils.SqlNullString(currentDate))
	if err != nil {
		t.Fatalf("failed retrieve task: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("actual does not match expect")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestRetrieveTasks_FailWithTimedOutContext(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")

	mock.ExpectQuery("SELECT id, detail, user_id, created_date FROM tasks").
		WithArgs(uint(userID), currentDate).
		WillReturnError(errors.New("error context time out"))

	repository := NewPostgresRepository(db)
	timedOutCtx, cancle := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancle()

	result, err := repository.RetrieveTasks(timedOutCtx, uint(userID), utils.SqlNullString(currentDate))
	if err == nil || result != nil {
		t.Fatal("expect to fail RetrieveTasks due to context time out and the returned result is nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestAddTask_Success(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")

	task := models.Task{
		ID:          uuid.New().String(),
		Detail:      "new task",
		UserID:      uint(userID),
		CreatedDate: currentDate,
	}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs(&task.ID, &task.Detail, &task.UserID, &task.CreatedDate).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewPostgresRepository(db)
	err := repository.AddTask(context.Background(), &task)
	if err != nil {
		t.Fatalf("failed adding task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestAddTask_Fail(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")

	task := models.Task{
		ID:          uuid.New().String(),
		Detail:      "new task",
		CreatedDate: currentDate,
	}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs(&task.ID, &task.Detail, &task.UserID, currentDate).
		WillReturnError(fmt.Errorf("nil user_id"))

	repository := NewPostgresRepository(db)
	err := repository.AddTask(context.Background(), &task)
	if err == nil {
		t.Fatalf("AddTask should fail due to nil user_id")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}
