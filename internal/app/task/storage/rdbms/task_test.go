package rdbms

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/util"
	"reflect"
	"regexp"
	"testing"
)

func TestTaskStorage_RetrieveTasks(t1 *testing.T) {
	t1.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t1.Error(err)
		return
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`)).
		WithArgs("firstUser", "2020-06-29").
		WillReturnRows(mock.NewRows([]string{"id", "content", "user_id", "created_date"}).
			AddRow("e1da0b9b-7ecc-44f9-82ff-4623cc50446a", "first content", "firstUser", "2020-06-29"))
	mock.ExpectQuery(`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \? AND created_date = \?`).
		WithArgs("firstUser", "2020-06-30").
		WillReturnError(errors.New("sql: no rows in result set"))

	type args struct {
		userID      string
		createdDate string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Task
		wantErr bool
	}{
		{
			name: "retrieve tasks successfully",
			args: args{
				userID:      "firstUser",
				createdDate: "2020-06-29",
			},
			want: []model.Task{
				{
					ID:          "e1da0b9b-7ecc-44f9-82ff-4623cc50446a",
					Content:     "first content",
					UserID:      "firstUser",
					CreatedDate: "2020-06-29",
				},
			},
			wantErr: false,
		},
		{
			name: "not found any task",
			args: args{
				userID:      "firstUser",
				createdDate: "2020-06-30",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskStorage{
				db:         db,
				driverName: util.DriverSQLite3,
			}
			got, err := t.RetrieveTasks(context.Background(), tt.args.userID, tt.args.createdDate)
			if (err != nil) != tt.wantErr {
				t1.Errorf("RetrieveTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("RetrieveTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskStorage_AddTask(t1 *testing.T) {
	t1.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t1.Error(err)
		return
	}
	defer db.Close()
	mock.ExpectExec(`INSERT INTO tasks \(id, content, user_id, created_date\) VALUES \(\?, \?, \?, \?\)`).
		WithArgs("e1da0b9b-7ecc-44f9-82ff-4623cc50446a", "first content", "firstUser", "2020-06-29").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO tasks \(id, content, user_id, created_date\) VALUES \(\?, \?, \?, \?\)`).
		WithArgs("e1da0b9b-7ecc-44f9-82ff-4623cc50446a", "second content", "firstUser", "2020-06-29").
		WillReturnError(errors.New("mock error"))

	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "add task successfully",
			args: args{
				task: model.Task{
					ID:          "e1da0b9b-7ecc-44f9-82ff-4623cc50446a",
					Content:     "first content",
					UserID:      "firstUser",
					CreatedDate: "2020-06-29",
				},
			},
			wantErr: false,
		},
		{
			name: "unable to add task",
			args: args{
				task: model.Task{
					ID:          "e1da0b9b-7ecc-44f9-82ff-4623cc50446a",
					Content:     "second content",
					UserID:      "firstUser",
					CreatedDate: "2020-06-29",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskStorage{
				db:         db,
				driverName: util.DriverSQLite3,
			}
			err := t.AddTask(context.Background(), tt.args.task)
			if (err != nil) != tt.wantErr {
				t1.Errorf("RetrieveTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTaskStorage_LimitReached(t1 *testing.T) {
	t1.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t1.Error(err)
		return
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(SQLLimitReached)).
		WithArgs("firstUser", "firstUser", "2020-06-28").
		WillReturnError(errors.New("mock error"))
	mock.ExpectQuery(regexp.QuoteMeta(SQLLimitReached)).
		WithArgs("firstUser", "firstUser", "2020-06-29").
		WillReturnRows(sqlmock.NewRows([]string{""}).AddRow(0))
	mock.ExpectQuery(regexp.QuoteMeta(SQLLimitReached)).
		WithArgs("firstUser", "firstUser", "2020-06-30").
		WillReturnRows(sqlmock.NewRows([]string{""}).AddRow(1))

	type args struct {
		userID      string
		createdDate string
	}
	tests := []struct {
		name    string
		args    args
		reached bool
		wantErr bool
	}{
		{
			name: "unable to check the limit",
			args: args{
				userID:      "firstUser",
				createdDate: "2020-06-28",
			},
			reached: false,
			wantErr: true,
		},
		{
			name: "the limit is not reached",
			args: args{
				userID:      "firstUser",
				createdDate: "2020-06-29",
			},
			reached: false,
			wantErr: false,
		},
		{
			name: "the limit is reached",
			args: args{
				userID:      "firstUser",
				createdDate: "2020-06-30",
			},
			reached: true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskStorage{
				db:         db,
				driverName: util.DriverPostgres,
			}
			got, err := t.LimitReached(context.Background(), tt.args.userID, tt.args.createdDate)
			if (err != nil) != tt.wantErr {
				t1.Errorf("RetrieveTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.reached {
				t1.Errorf("RetrieveTasks() gotReached = %v, wantReached %v", got, tt.reached)
				return
			}
		})
	}
}
