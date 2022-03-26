package task

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository/postgres"
	"github.com/manabie-com/togo/internal/utils"
)

func TestRetrieveTask(t *testing.T) {
	t.Parallel()
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")
	taskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
	rows := sqlmock.NewRows([]string{"id", "detail", "user_id", "created_date"}).
		AddRow(taskIds[0], "first task", uint(userID), currentDate).
		AddRow(taskIds[1], "second task", uint(userID), currentDate).
		AddRow(taskIds[2], "third task", uint(userID), currentDate)

	mock.ExpectQuery("SELECT id, detail, user_id, created_date FROM tasks").
		WithArgs(uint(userID), currentDate).
		WillReturnRows(rows)

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	taskUsecase := NewTaskUsecase(repository)

	expect := []*models.Task{
		{
			ID:          taskIds[0],
			Detail:      "first task",
			UserID:      uint(userID),
			CreatedDate: currentDate,
		},
		{
			ID:          taskIds[1],
			Detail:      "second task",
			UserID:      uint(userID),
			CreatedDate: currentDate,
		},
		{
			ID:          taskIds[2],
			Detail:      "third task",
			UserID:      uint(userID),
			CreatedDate: currentDate,
		},
	}

	actual, err := taskUsecase.RetrieveTasks(context.Background(), uint(userID), utils.SqlNullString(currentDate))
	if err != nil {
		t.Errorf("fail retrieve task: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect: %#v, actual: %#v", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestAddTask_Success(t *testing.T) {
	t.Parallel()
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-2")
	task := models.Task{
		ID:          uuid.New().String(),
		Detail:      "new task",
		UserID:      uint(userID),
		CreatedDate: currentDate,
	}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs(task.ID, task.Detail, task.UserID, task.CreatedDate).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	taskUsecase := NewTaskUsecase(repository)

	if err := taskUsecase.AddTask(context.Background(), &task); err != nil {
		t.Errorf("fail add task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestIsMaxTaskPerDay_Reached(t *testing.T) {
	t.Parallel()
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")
	taskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}

	rows := sqlmock.NewRows([]string{"id", "detail", "user_id", "created_date"}).
		AddRow(taskIds[0], "first task", uint(userID), currentDate).
		AddRow(taskIds[1], "second task", uint(userID), currentDate).
		AddRow(taskIds[2], "third task", uint(userID), currentDate).
		AddRow(taskIds[3], "fourth task", uint(userID), currentDate).
		AddRow(taskIds[4], "fifth task", uint(userID), currentDate)

	mock.ExpectQuery("SELECT id, detail, user_id, created_date FROM tasks").
		WithArgs(uint(userID), currentDate).
		WillReturnRows(rows)

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	taskUsecase := NewTaskUsecase(repository)
	expect := true
	actual, err := taskUsecase.IsMaxTasksPerDay(
		context.Background(),
		uint(userID),
		uint(maxTaskPerDay),
		utils.SqlNullString(currentDate),
	)
	if err != nil {
		t.Errorf("fail IsMaxTasksPerDay: %s", err)
	}

	if expect != actual {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestIsMaxTaskPerDay_NotReached(t *testing.T) {
	t.Parallel()
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")
	taskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}

	rows := sqlmock.NewRows([]string{"id", "detail", "user_id", "created_date"}).
		AddRow(taskIds[0], "first task", uint(userID), currentDate).
		AddRow(taskIds[1], "second task", uint(userID), currentDate).
		AddRow(taskIds[2], "third task", uint(userID), currentDate).
		AddRow(taskIds[3], "fourth task", uint(userID), currentDate)

	mock.ExpectQuery("SELECT id, detail, user_id, created_date FROM tasks").
		WithArgs(uint(userID), currentDate).
		WillReturnRows(rows)

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	taskUsecase := NewTaskUsecase(repository)
	expect := false
	actual, err := taskUsecase.IsMaxTasksPerDay(
		context.Background(),
		uint(userID),
		uint(maxTaskPerDay),
		utils.SqlNullString(currentDate),
	)
	if err != nil {
		t.Errorf("fail IsMaxTasksPerDay: %s", err)
	}

	if expect != actual {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}

func TestIsMaxTaskPerDay_FailDueToTimedOutContext(t *testing.T) {
	t.Parallel()
	db, mock := setupMock()
	defer db.Close()

	currentDate := time.Now().Format("2006-01-02")

	mock.ExpectQuery("SELECT id, detail, user_id, created_date FROM tasks").
		WithArgs(uint(userID), currentDate).
		WillReturnError(errors.New("error context time out"))

	// Test with Postgres repo
	repository := postgres.NewPostgresRepository(db)
	taskUsecase := NewTaskUsecase(repository)
	expect := false
	timedOutCtx, cancle := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancle()

	actual, err := taskUsecase.IsMaxTasksPerDay(
		timedOutCtx,
		uint(userID),
		uint(maxTaskPerDay),
		utils.SqlNullString(currentDate),
	)
	if err == nil {
		t.Error("expect to fail IsMaxTasksPerDay due to context time out")
	}

	if expect != actual {
		t.Errorf("expect: %t, actual: %t", expect, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("at least 1 expectation was not met: %s", err)
	}
}
