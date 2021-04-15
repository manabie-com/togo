package rdbms

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/util"
	"testing"
)

func TestUserStorage_ValidateUser(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT id FROM users WHERE id = \? AND password = \?`).WithArgs("firstUser", "example").WillReturnRows(mock.NewRows([]string{"id"}).AddRow("firstUser"))
	mock.ExpectQuery(`SELECT id FROM users WHERE id = \? AND password = \?`).WithArgs("invalidUser", "example").WillReturnError(errors.New("sql: no rows in result set"))

	type args struct {
		userID string
		pwd    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid user",
			args: args{
				userID: "firstUser",
				pwd:    "example",
			},
			wantErr: false,
		},
		{
			name: "invalid user",
			args: args{
				userID: "invalidUser",
				pwd:    "example",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := UserStorage{
				db:         db,
				driverName: util.DriverSQLite3,
			}
			if got := us.ValidateUser(context.Background(), tt.args.userID, tt.args.pwd); (got != nil) != tt.wantErr {
				t.Errorf("ValidateUser() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
