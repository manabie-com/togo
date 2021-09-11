package sqlite

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/manabie-com/togo/internal/storages"
)

const password = "example"

func TestValidateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	testCases := map[string]struct{
		id string
		password string
		isValid bool
	}{
		"valid": {
			id:       "1",
			password: password,
			isValid:  true,
		},
		"invalid": {
			id:       "2",
			password: "123456",
			isValid:  false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			u := &storages.User{
				ID:       tc.id,
				Password: tc.password,
			}
			rows := sqlmock.NewRows([]string{"id", "password"}).AddRow(u.ID, password)
			mock.ExpectQuery(regexp.QuoteMeta(sqlValidateUser)).WithArgs(u.ID, u.Password).WillReturnRows(rows)

			liteDB := NewLiteDB(db)
			isValid, err := liteDB.ValidateUser(
				context.Background(),
				sql.NullString{
					String: u.ID,
					Valid:  true,
				},
				sql.NullString{
					String: u.Password,
					Valid:  true,
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.isValid, isValid)
		})
	}
}

func TestAddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	task := &storages.Task{
		ID:          "1",
		Content:     "hash password",
		UserID:      "1",
		CreatedDate: time.Now().String(),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sqlAddTask)).WithArgs(task.ID, task.Content, task.UserID, task.CreatedDate).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	liteDB := NewLiteDB(db)
	require.NoError(t, liteDB.AddTask(context.Background(), task))
}

func TestRetrieveTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	task := &storages.Task{
		ID:          "1",
		Content:     "hash password",
		UserID:      "1",
		CreatedDate: time.Now().String(),
	}
	rows := sqlmock.NewRows([]string{"id", "content", "userID", "createdDate"}).AddRow(task.ID, task.Content, task.UserID, task.CreatedDate)
	mock.ExpectQuery(regexp.QuoteMeta(sqlRetrieveTasks)).WithArgs(task.UserID, task.CreatedDate).WillReturnRows(rows)

	liteDB := NewLiteDB(db)
	tasks, err := liteDB.RetrieveTasks(
		context.Background(),
		sql.NullString{
			String: task.UserID,
			Valid:  true,
		},
		sql.NullString{
			String: task.CreatedDate,
			Valid:  true,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, "1", tasks[0].ID)
	assert.Equal(t, "hash password", tasks[0].Content)
}
