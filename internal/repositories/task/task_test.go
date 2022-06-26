package task

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func Test_repository_AddTask(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()

	taskRepository := NewTaskRepository(db)

	{ //Success case
		input := NewTask(models.Task{
			Content:    "Test Interview",
			CreateDate: time.Now().Format("2006-01-02"),
			UserID:     "manabie",
		})
		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("id","content","create_date","user_id") VALUES ($1,$2,$3,$4) RETURNING "tasks"."id"`)).
			WithArgs(input.ID, input.Content, input.CreateDate, input.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := taskRepository.AddTask(input)
		require.Nil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}

	{ //Fail case
		input := NewTask(models.Task{
			Content:    "Test Interview",
			CreateDate: time.Now().Format("2006-01-02"),
			UserID:     "manabie",
		})
		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("id","content","create_date","user_id") VALUES ($1,$2,$3,$4) RETURNING "tasks"."id"`)).
			WithArgs("", input.Content, input.CreateDate, input.UserID).
			WillReturnError(fmt.Errorf("nil task_id"))
		mock.ExpectCommit()

		err := taskRepository.AddTask(input)
		require.NotNil(t, err)
		require.NotNil(t, mock.ExpectationsWereMet())
	}

}

func Test_repository_FindTaskByUser(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()

	taskRepository := NewTaskRepository(db)

	{ //Success case
		input := &models.Task{
			Content:    "task_manabie_1",
			CreateDate: "2022-26-06",
			UserID:     "manabie_1",
		}
		// Create mock data for test
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(1, input.Content, input.CreateDate, input.UserID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err := recordStats(db, *input)
		require.Nil(t, err)

		rows := sqlmock.NewRows([]string{"id", "content", "create_date", "user_id"}).
			AddRow("1", input.Content, input.CreateDate, input.UserID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE (user_id = $1 and create_date = $2)`)).
			WithArgs(input.UserID, input.CreateDate).
			WillReturnRows(rows)

		tasks, err := taskRepository.FindTaskByUser(input.UserID, input.CreateDate)
		require.Nil(t, err)
		require.NotNil(t, tasks)
		require.Equal(t, 1, len(tasks))
		require.Equal(t, input.UserID, tasks[0].UserID)
		require.Equal(t, input.Content, tasks[0].Content)
		require.Equal(t, input.CreateDate, tasks[0].CreateDate)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}

	{ //fail case
		input := &models.Task{
			Content:    "fail_task_manabie_1",
			CreateDate: "2022-26-06",
			UserID:     "fail_manabie_1",
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE (user_id = $1 and create_date = $2)`)).
			WithArgs(input.UserID, input.CreateDate).
			WillReturnError(fmt.Errorf("invalid task"))

		tasks, err := taskRepository.FindTaskByUser(input.UserID, input.CreateDate)
		require.NotNil(t, err)
		require.Nil(t, tasks)
		require.Equal(t, 0, len(tasks))

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("at least 1 expectation was not met: %s", err)
		}
	}

}
