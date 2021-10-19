package task

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/utils"
)

var (
	userID  = 1
	maxTodo = 5
)

// setup mock
func setupMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
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

	repository := postgres.NewPostgresRepository(db)
	usecase := NewTaskUseCase(repository)

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
	actual, err := usecase.ListTasks(
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

	repository := postgres.NewPostgresRepository(db)
	usecase := NewTaskUseCase(repository)

	err := usecase.AddTask(
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

func TestIsMaximumTasksHasReached(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)
	listTaskIds := []string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}
	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(listTaskIds[0], "This is just the first test", 1, currentDate).
		AddRow(listTaskIds[1], "This is just the second test", 1, currentDate).
		AddRow(listTaskIds[2], "This is just the third test", 1, currentDate).
		AddRow(listTaskIds[3], "This is just the four test", 1, currentDate).
		AddRow(listTaskIds[4], "This is just the five test", 1, currentDate)

	mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = (.+) AND created_date = (.+)").
		WillReturnRows(rows)

	repository := postgres.NewPostgresRepository(db)
	usecase := NewTaskUseCase(repository)

	expected := true
	actual, err := usecase.IsMaximumTasks(
		context.TODO(),
		uint(userID),
		utils.SqlString(currentDate),
		uint(maxTodo),
	)
	if err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

	if expected != actual {
		t.Error("Not expected")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestIsMaximumTasksNotReached(t *testing.T) {
	db, mock := setupMock()
	defer db.Close()

	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)
	listTaskIds := []string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}
	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(listTaskIds[0], "This is just the first test", 1, currentDate).
		AddRow(listTaskIds[1], "This is just the second test", 1, currentDate).
		AddRow(listTaskIds[2], "This is just the third test", 1, currentDate).
		AddRow(listTaskIds[3], "This is just the four test", 1, currentDate)

	mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = (.+) AND created_date = (.+)").
		WillReturnRows(rows)

	repository := postgres.NewPostgresRepository(db)
	usecase := NewTaskUseCase(repository)

	expected := false
	actual, err := usecase.IsMaximumTasks(
		context.TODO(),
		uint(userID),
		utils.SqlString(currentDate),
		uint(maxTodo),
	)
	if err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

	if expected != actual {
		t.Error("Not expected")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
