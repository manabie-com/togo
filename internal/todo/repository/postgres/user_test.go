package postgres

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/stretchr/testify/assert"
)

func TestPGUserRepository_GetByCredentials(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		username string
		password string
	}

	fixedTime := time.Date(2021, 3, 17, 0, 0, 0, 0, time.Now().Location())
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *d.User
		wantErr bool
	}{
		{
			"System Error",
			args{"test", "test"},
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WithArgs("test", "test").
					WillReturnError(errors.New("System ERror"))
			},
			nil,
			true,
		},
		{
			"No data",
			args{"test", "test"},
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WithArgs("test", "test").
					WillReturnError(sql.ErrNoRows)
			},
			nil,
			false,
		},
		{
			"Success",
			args{"test", "test"},
			func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at"}).
					AddRow(1, "test", "test", &fixedTime)
				mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WithArgs("test", "test").
					WillReturnRows(rows)
			},
			&d.User{ID: 1, Username: "test", Password: "test", CreatedAt: &fixedTime},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, _ := sqlmock.New()
			defer mockDB.Close()
			dbConn := sqlx.NewDb(mockDB, "sqlmock")
			tr := &PGUserRepository{
				PGRepository: PGRepository{DBConn: dbConn},
			}
			tt.mock(mock)
			got, err := tr.GetByCredentials(context.Background(), tt.args.username, tt.args.password)

			assert.NoError(mock.ExpectationsWereMet())
			assert.Equal(tt.want, got)
			assert.True((err != nil) == tt.wantErr)
		})
	}
}
