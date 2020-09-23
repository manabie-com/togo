package sqllite

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

func TestLiteDB_RetrieveTasks(t *testing.T) {
	type args struct {
		ctx         context.Context
		userID      sql.NullString
		createdDate sql.NullString
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *LiteDB
		inspect func(r *LiteDB, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1      []*storages.Task
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			name: "success",
			init: func(*testing.T) *LiteDB { //copy the conf/conf.json here to proceed the test
				db, err := sql.Open("sqlite3", "./../../../data.db")
				if err != nil {
					log.Fatal("error opening db", err)
				}
				return &LiteDB{DB: db}
			},
			args: func(*testing.T) args {
				return args{
					ctx:         context.Background(),
					userID:      sql.NullString{"firstUser", true},
					createdDate: sql.NullString{"2020-06-29", true},
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1, err := receiver.RetrieveTasks(tArgs.ctx, tArgs.userID, tArgs.createdDate)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !(len(got1) > 0) {
				t.Errorf("LiteDB.RetrieveTasks got1 = %v", got1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("LiteDB.RetrieveTasks error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestLiteDB_AddTask(t *testing.T) {
	type args struct {
		ctx context.Context
		t   *storages.Task
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *LiteDB
		inspect func(r *LiteDB, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			err := receiver.AddTask(tArgs.ctx, tArgs.t)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("LiteDB.AddTask error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestLiteDB_ValidateUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID sql.NullString
		pwd    sql.NullString
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *LiteDB
		inspect func(r *LiteDB, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1 bool
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1 := receiver.ValidateUser(tArgs.ctx, tArgs.userID, tArgs.pwd)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LiteDB.ValidateUser got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestLiteDB_FindUserByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *LiteDB
		inspect func(r *LiteDB, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1      *storages.User
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1, err := receiver.FindUserByID(tArgs.ctx, tArgs.userID)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LiteDB.FindUserByID got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("LiteDB.FindUserByID error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
