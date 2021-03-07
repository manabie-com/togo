package sqllite

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/model"
	"testing"
)

func TestUserStorage_ValidateUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' when opening a stub database connection", err)
	}
	defer db.Close()

	userID, pwd := "10" , "123456"
	s := NewUserLiteDBStorage(db)
	tests := []struct {
		name    string
		s       model.UserStorage
		mock    func()
		want    bool
		wantErr bool
	} {
		{
			name: "OK",
			s: s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("10")
				mock.ExpectQuery("SELECT id FROM users WHERE id = \\? AND password = \\?").WithArgs(userID, pwd).WillReturnRows(rows)
			},
			want: true,
			wantErr: false,
		},

		{
			name: "Invalid Id/Not Found Id",
			s: s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT id FROM users WHERE id = \\? AND password = \\? ").WithArgs(userID, pwd).WillReturnRows(rows)
			},
			wantErr: true,
		},

		{
			name: "Wrong table name",
			s: s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT id FROM user WHERE id = \\? AND password = \\? ").WithArgs(userID, pwd).WillReturnRows(rows)
			},
			wantErr: true,
		},

		}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.s.ValidateUser(context.TODO(),userID, pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil &&  !(got == tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

