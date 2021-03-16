package postgres

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/stretchr/testify/assert"
)

func TestPGTaskRepository_CreateTaskForUser(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	os.Setenv("MAX_TASKS_DAILY", "5")
	type args struct {
		userID    int
		taskParam d.TaskCreateParam
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *d.Task
		wantErr bool
	}{
		{
			"System Error",
			args{1, d.TaskCreateParam{Content: "content"}},
			func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("SELECT (.+) FROM users WHERE").WithArgs(1).WillReturnError(errors.New("System Error"))
				mock.ExpectRollback()
			},
			nil,
			true,
		},
		{
			"System Error 2",
			args{1, d.TaskCreateParam{Content: "content"}},
			func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("SELECT (.+) FROM users WHERE").WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(`INSERT INTO tasks (.+) RETURNING`).WithArgs(1, "content", 5).
					WillReturnError(errors.New("System Error"))
				mock.ExpectRollback()
			},
			nil,
			true,
		},
		{
			"Empty result",
			args{1, d.TaskCreateParam{Content: "content"}},
			func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("SELECT (.+) FROM users WHERE").WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(`INSERT INTO tasks (.+) RETURNING`).WithArgs(1, "content", 5).
					WillReturnRows(sqlmock.NewRows(nil))
				mock.ExpectRollback()
			},
			nil,
			false,
		},
		{
			"Success",
			args{1, d.TaskCreateParam{Content: "content"}},
			func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).
					AddRow(1, "content", 1, &now)
				mock.ExpectBegin()
				mock.ExpectExec("SELECT (.+) FROM users WHERE").WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(`INSERT INTO tasks (.+) RETURNING`).WithArgs(1, "content", 5).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
			&d.Task{ID: 1, Content: "content", UserID: 1, CreatedAt: &now},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, _ := sqlmock.New()
			defer mockDB.Close()

			dbConn := sqlx.NewDb(mockDB, "sqlmock")
			tr := &PGTaskRepository{
				PGRepository: PGRepository{DBConn: dbConn},
			}
			tt.mock(mock)
			got, err := tr.CreateTaskForUser(context.Background(), tt.args.userID, tt.args.taskParam)

			assert.NoError(mock.ExpectationsWereMet())
			assert.Equal(tt.want, got)
			assert.True((err != nil) == tt.wantErr)
		})
	}
}

func TestPGTaskRepository_GetTasksForUser(t *testing.T) {
	assert := assert.New(t)
	fixedTime := time.Now()
	type args struct {
		userID  int
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    []*d.Task
		wantErr bool
	}{
		{
			"System Error",
			args{1, "2021-03-20"},
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE").WithArgs(1, "2021-03-20 00:00:00", "2021-03-20 23:59:59").
					WillReturnError(errors.New("System Error"))
			},
			nil,
			true,
		},
		{
			"Success",
			args{1, "2021-03-20"},
			func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).
					AddRow(1, "abc", 1, &fixedTime).
					AddRow(2, "bcd", 1, &fixedTime)
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE").WithArgs(1, "2021-03-20 00:00:00", "2021-03-20 23:59:59").
					WillReturnRows(rows)
			},
			[]*d.Task{
				{ID: 1, Content: "abc", UserID: 1, CreatedAt: &fixedTime},
				{ID: 2, Content: "bcd", UserID: 1, CreatedAt: &fixedTime},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, _ := sqlmock.New()
			defer mockDB.Close()

			dbConn := sqlx.NewDb(mockDB, "sqlmock")
			tr := &PGTaskRepository{
				PGRepository: PGRepository{DBConn: dbConn},
			}
			tt.mock(mock)
			got, err := tr.GetTasksForUser(context.Background(), tt.args.userID, tt.args.dateStr)

			assert.NoError(mock.ExpectationsWereMet())
			assert.Equal(tt.want, got)
			assert.True((err != nil) == tt.wantErr)
		})
	}
}
