package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error : '%s'", err)
	}
	time.Now().Format("2006-01-02")
	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(1, "Learning Go", "firstUser", time.Now())
	query := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\$1 AND to_char(created_date,'YYYY-MM-DD') = \\$2"
	mock.ExpectQuery(query).WillReturnRows(rows)

	pDB := postgres.NewPostgresDB(db)
	userID := "firstUser"
	created_date := time.Now()
	listTask, err := pDB.RetrieveTasks(context.TODO(), userID, created_date)
	assert.NoError(t, err)
	assert.Len(t, listTask, 1)
}

func TestAddTask(t *testing.T) {
	ta := &storages.Task{
		Content:     "Reading book",
		CreatedDate: time.Now(),
		UserID:      "firstUser",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error : '%s'", err)
	}
	query := "INSERT INTO tasks (content, user_id, created_date) VALUES (\\?, \\?, \\?)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ta.ID, ta.Content, ta.UserID, ta.CreatedDate).WillReturnResult(sqlmock.NewResult(0, 1))
	pDB := postgres.NewPostgresDB(db)
	pDB.AddTask(context.TODO(), ta)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), ta.ID)

}
