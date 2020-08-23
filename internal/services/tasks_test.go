package services

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phuwn/togo/internal/storages"
	"github.com/phuwn/togo/internal/storages/mocks"
	"github.com/phuwn/togo/util"
	"github.com/stretchr/testify/mock"
)

var (
	exampleUserID = "firstUser"
	validTaskDate = "2020-08-20"
)

func TestToDoService_listTasks(t *testing.T) {
	store := &mocks.Store{}
	store.On("RetrieveTasks",
		mock.Anything,
		sql.NullString{exampleUserID, true},
		sql.NullString{validTaskDate, true},
	).Return([]*storages.Task{
		&storages.Task{"1", "content_1", exampleUserID, validTaskDate},
		&storages.Task{"2", "content_2", exampleUserID, validTaskDate},
		&storages.Task{"3", "content_3", exampleUserID, validTaskDate},
	}, nil)
	store.On("RetrieveTasks", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	s := &ToDoService{
		Store: store,
	}

	type args struct {
		userID      string
		createdDate string
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			"happy case",
			args{exampleUserID, validTaskDate},
			http.StatusOK,
			`{"data":[{"id":"1","content":"content_1","user_id":"firstUser","created_date":"2020-08-20"},{"id":"2","content":"content_2","user_id":"firstUser","created_date":"2020-08-20"},{"id":"3","content":"content_3","user_id":"firstUser","created_date":"2020-08-20"}]}` + "\n",
		},
		{
			"empty_task date case",
			args{exampleUserID, "2020-08-19"},
			http.StatusOK,
			`{"data":null}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
			if err != nil {
				t.Errorf("failed to create request, err: %s\n", err.Error())
				return
			}

			q := req.URL.Query()
			q.Add("created_date", tt.args.createdDate)
			req.URL.RawQuery = q.Encode()
			req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), tt.args.userID))

			rr := httptest.NewRecorder()
			s.listTasks(rr, req)
			if rr.Code != tt.code {
				t.Errorf("unexpected response: status code want %v, got %v, error message: %s", tt.code, rr.Code, rr.Body.String())
				return
			}

			if rr.Body.String() != tt.want {
				t.Errorf("unexpected output, got %s, want %s", rr.Body.String(), tt.want)
				return
			}
		})
	}
}

func TestToDoService_addTask(t *testing.T) {
	util.MockRuntimeFunc()

	type args struct {
		body    []byte
		userID  string
		errResp error
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			"happy case",
			args{[]byte(`{"content":"example_content"}`), exampleUserID, nil},
			http.StatusCreated,
			`{"data":{"id":"af1c772f-9abd-4e3c-94af-80d57d262028","content":"example_content","user_id":"firstUser","created_date":"2020-08-20"}}` + "\n",
		},
		{
			"task_create limit reach error",
			args{[]byte(`{"content":"example_content"}`), exampleUserID, fmt.Errorf(CreateTaskLimitErrMsg)},
			http.StatusTooManyRequests,
			fmt.Sprintf(`{"error":"%v"}`+"\n", CreateTaskLimitErrResp),
		},
		{
			"no content provided",
			args{nil, exampleUserID, nil},
			http.StatusBadRequest,
			fmt.Sprintf(`{"error":"%v"}`+"\n", InvalidBodyErrResp),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &mocks.Store{}
			store.On("AddTask", mock.Anything, mock.Anything).Return(tt.args.errResp)

			s := &ToDoService{
				Store: store,
			}

			body := bytes.Buffer{}
			body.Write(tt.args.body)
			req, err := http.NewRequest(http.MethodPost, "/tasks", &body)
			if err != nil {
				t.Errorf("failed to create request, err: %s\n", err.Error())
				return
			}

			req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), tt.args.userID))

			rr := httptest.NewRecorder()
			s.addTask(rr, req)
			if rr.Code != tt.code {
				t.Errorf("unexpected response: status code want %v, got %v, error message: %s", tt.code, rr.Code, rr.Body.String())
				return
			}

			if rr.Body.String() != tt.want {
				t.Errorf("unexpected output, got %s, want %s", rr.Body.String(), tt.want)
				return
			}
		})
	}
}
