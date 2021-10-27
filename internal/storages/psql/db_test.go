package psql

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	jwtKey   = "wqGyEBBfPK9w3Lxw"
	email    = "me@here.com"
	password = "password"
	//hashpassword for mock
	hashPassword = "$2a$12$orZppdmhH.KRrxcZcjx0NeLPtIDpaf2GNUben4Rz7w53e5dSQJgdq"
	userID       = 1
)

func TestValidateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	testCases := map[string]struct {
		email    string
		password string
		isValid  bool
	}{
		"valid": {
			email:    "me@here.com",
			password: password,
			isValid:  true,
		},
		"invalid": {
			email:    "test@gmail.com",
			password: "123456",
			isValid:  false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			u := &storages.User{
				Email:    tc.email,
				Password: tc.password,
			}

			rows := sqlmock.NewRows([]string{"password", "email"}).AddRow(hashPassword, email)
			mock.ExpectQuery(regexp.QuoteMeta(sqlValidateUser)).WithArgs(u.Email).WillReturnRows(rows)
			liteDB := NewModels(db)
			isValid := liteDB.ValidateUser(
				sql.NullString{
					String: u.Email,
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

func TestGetUserFromEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	testCases := map[string]struct {
		id       int
		max_todo int
		email    string
		isValid  bool
	}{
		"UserOne": {
			id:       1,
			max_todo: 5,
			email:    "test2@here.com",
		},
		"UserTwo": {
			id:       2,
			max_todo: 10,
			email:    "test@here.com",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			u := &storages.User{
				ID:      tc.id,
				MaxTodo: tc.max_todo,
				Email:   tc.email,
			}

			rows := sqlmock.NewRows([]string{"id", "max_todo", "email"}).AddRow(u.ID, u.MaxTodo, u.Email)
			mock.ExpectQuery(regexp.QuoteMeta(sqlGetUserFromEmail)).WithArgs(u.Email).WillReturnRows(rows)
			liteDB := NewModels(db)
			user, err := liteDB.GetUserFromEmail(
				tc.email,
			)
			require.NoError(t, err)
			require.Equal(t, user.MaxTodo, u.MaxTodo)
			require.Equal(t, user.ID, u.ID)

		})
	}
}

func TestRetrieveTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	u := &storages.User{
		Email:   "test@gmail.com",
		ID:      1,
		MaxTodo: 0,
	}
	task := &storages.Task{
		ID:        1,
		Content:   "hash password",
		UserID:    1,
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	rows := sqlmock.NewRows([]string{"id", "max_todo", "email"}).AddRow(u.ID, u.MaxTodo, u.Email)
	mock.ExpectQuery(regexp.QuoteMeta(sqlGetUserFromEmail)).WithArgs(u.Email).WillReturnRows(rows)
	liteDB := NewModels(db)
	user, err := liteDB.GetUserFromEmail(
		u.Email,
	)
	require.NoError(t, err)
	require.Equal(t, user.MaxTodo, u.MaxTodo)
	require.Equal(t, user.ID, u.ID)

	rowsTask := sqlmock.NewRows([]string{"id", "max_todo", "email"}).AddRow(u.ID, u.MaxTodo, u.Email)
	mock.ExpectQuery(regexp.QuoteMeta(sqlGetUserFromEmail)).WithArgs(u.Email).WillReturnRows(rowsTask)
	rowsTask = sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).AddRow(task.ID, task.Content, task.UserID, task.CreatedAt)
	mock.ExpectQuery(regexp.QuoteMeta(sqlRetrieveTasks)).WithArgs(task.UserID, task.CreatedAt.String()).WillReturnRows(rowsTask)
	tasks, err := liteDB.RetrieveTasks(
		sql.NullString{
			String: u.Email,
			Valid:  true,
		},
		sql.NullString{
			String: task.CreatedAt.String(),
			Valid:  true,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, 1, tasks[0].ID)
	assert.Equal(t, "hash password", tasks[0].Content)
}

func TestAddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	u := &storages.User{
		Email:   "test@gmail.com",
		ID:      1,
		MaxTodo: 0,
	}
	task := &storages.Task{
		ID:        1,
		Content:   "hash password",
		UserID:    1,
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}
	rows := sqlmock.NewRows([]string{"id", "max_todo", "email"}).AddRow(u.ID, u.MaxTodo, u.Email)
	mock.ExpectQuery(regexp.QuoteMeta(sqlGetUserFromEmail)).WithArgs(u.Email).WillReturnRows(rows)
	liteDB := NewModels(db)
	user, err := liteDB.GetUserFromEmail(
		u.Email,
	)
	require.NoError(t, err)
	require.Equal(t, user.MaxTodo, u.MaxTodo)
	require.Equal(t, user.ID, u.ID)

	rowsTask := sqlmock.NewRows([]string{"id", "max_todo", "email"}).AddRow(u.ID, u.MaxTodo, u.Email)
	mock.ExpectQuery(regexp.QuoteMeta(sqlGetUserFromEmail)).WithArgs(u.Email).WillReturnRows(rowsTask)
	rowsTask = sqlmock.NewRows([]string{"id"}).AddRow(task.ID)
	dateToday := time.Now().Format("2006-01-02")
	//incrementTodo here since it creates task for the user
	incrementTodo := u.MaxTodo + 1
	mock.ExpectQuery(regexp.QuoteMeta(`select count(id) FROM tasks where user_id = $1 AND DATE(created_at) = $2`)).WithArgs(u.ID, dateToday).WillReturnRows(rowsTask)

	mock.ExpectBegin()
	//another increment todo here for inserting task from user
	incrementTodo = incrementTodo + 1
	mock.ExpectExec(regexp.QuoteMeta(sqlAddTask)).WithArgs(task.Content, task.UserID, task.CreatedAt, task.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`update users set max_todo = $1, updated_at = $2 where id = $3`)).WithArgs(incrementTodo, task.UpdatedAt, task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	require.NoError(t, liteDB.AddTask(task, u.Email))
}

func TestUpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	task := &storages.Task{
		ID:        1,
		Content:   "hash password",
		UserID:    1,
		CreatedAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sqlUpdateTask)).WithArgs(task.Content, task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	liteDB := NewModels(db)
	require.NoError(t, liteDB.UpdateTask(task))
}

func TestDeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	task := &storages.Task{
		ID:        1,
		Content:   "hash password",
		UserID:    1,
		CreatedAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sqlDeleteTask)).WithArgs(task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	liteDB := NewModels(db)
	require.NoError(t, liteDB.DeleteTask(task.ID))
}
