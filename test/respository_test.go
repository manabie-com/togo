package test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/respository"
)

func TestUserRespository(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' when opening a stub database connection", err)
	}
	defer db.Close()

	userID, pwd := "test", "123456"
	_pwd := "123123"
	s := respository.NewUserLiteDBRespository(db)
	testcases := []struct {
		name string
		s    model.UserRespository
		mock func()
		rs   bool
		err  bool
	}{
		{
			name: "Login success",
			s:    s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("test")
				mock.ExpectQuery("SELECT id FROM users WHERE id = \\? AND password = \\?").WithArgs(userID, pwd).WillReturnRows(rows)
			},
			rs:  true,
			err: false,
		},

		{
			name: "UserId not found in database",
			s:    s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT id FROM users WHERE id = \\? AND password = \\? ").WithArgs(userID, pwd).WillReturnRows(rows)
			},
			err: true,
		},

		{
			name: "Login failed with wrong password",
			s:    s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("test")
				mock.ExpectQuery("SELECT id FROM users WHERE id = \\? AND password = \\? ").WithArgs(userID, _pwd).WillReturnRows(rows)
			},
			err: true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.s.ValidateUser(context.TODO(), userID, pwd)
			if (err != nil) != tt.err {
				t.Errorf("Get() error new = %v, err %v", err, tt.err)
				return
			}
			if err == nil && !(got == tt.rs) {
				t.Errorf("Get() = %v, rs %v", got, tt.rs)
			}
		})
	}
}

func TestTaskRespository(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' when opening a stub database connection", err)
	}
	defer db.Close()

	id, content, user_id, created_date := "1", "test", "test", "2021-03-04"

	//id := 1
	s := respository.NewTaskLiteDBRespository(db)
	testcases := []struct {
		name string
		s    model.TaskRespository
		mock func()
		rs   bool
		err  bool
	}{
		{
			name: "Add task success",
			s:    s,
			mock: func() {
				mock.ExpectExec("INSERT INTO tasks").WithArgs(id, content, user_id, created_date).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: false,
		},
	}

	task := &model.Task{
		ID:          "1",
		Content:     "test",
		UserID:      "test",
		CreatedDate: "2021-03-04",
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := tt.s.AddTask(context.TODO(), task)
			if (err != nil) || (err == nil && tt.err) {
				t.Errorf("Get() error new = %v, err %v", err, tt.err)
				return
			}
		})
	}
}
