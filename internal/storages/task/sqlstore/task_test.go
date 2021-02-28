package sqlstore

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/task/model"
	cmsqlmock "github.com/manabie-com/togo/pkg/common/cmsql/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mock sqlmock.Sqlmock
	taskStore *TaskStore
)

func TestMain(m *testing.M) {
	var db *sql.DB
	db, mock = cmsqlmock.SetupMock()
	taskStore = NewTaskStore(db)
	m.Run()
	cmsqlmock.TeardownMock(db)
}

func TestStore_AddTask(t *testing.T) {
	t.Run("TestStore_AddTask", func(t *testing.T) {
		task := &model.Task{
			ID:          uuid.New().String(),
			Content:     "new task 1",
			UserID:      "00001",
			CreatedDate: "2021-02-22",
		}

		mock.ExpectExec("INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)").
			WithArgs(task.ID, task.Content, task.UserID, task.CreatedDate).
			WillReturnResult(sqlmock.NewResult(1, 1))

		if err := taskStore.AddTask(context.Background(), task); err != nil {
			t.Errorf("error was not expected while creating task: %s", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestTaskStore_RetrieveTasks(t *testing.T) {
	t.Run("TestTaskStore_RetrieveTasks", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2").
			WithArgs("user_00001", "2021-02-22").
			WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
				AddRow("task_0001", "content_0001", "user_00001", "2021-02-22").
				AddRow("task_0002", "content_0002", "user_00001", "2021-02-22"))

		tasksResult, err := taskStore.RetrieveTasks(context.Background(),
			sql.NullString{String: "user_00001", Valid:  true},
			sql.NullString{String: "2021-02-22", Valid: true})
		assert.Nil(t, err)
		assert.NotNil(t, tasksResult)
		assert.Equal(t, len(tasksResult), 2)
		assert.Equal(t, tasksResult[0].UserID, "user_00001")
		assert.Equal(t, tasksResult[0].CreatedDate, "2021-02-22")
	})

	t.Run("TestTaskStore_RetrieveTasks_Empty", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2").
			WithArgs("user_0002", "2021-02-22").
			WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}))

		tasksResult, err := taskStore.RetrieveTasks(context.Background(),
			sql.NullString{String: "user_0002", Valid:  true},
			sql.NullString{String: "2021-02-22", Valid: true})
		assert.Nil(t, err)
		assert.Nil(t, tasksResult)
	})
}
