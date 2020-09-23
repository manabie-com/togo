package services

import (
	"github.com/gavv/httpexpect/v2"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"log"
	"net/http/httptest"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"reflect"
	"testing"
)

func TestToDoService_ServeHTTP(t *testing.T) {
	db, err := sql.Open("sqlite3", "./../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	handler := &ToDoService{
		JWTKey: "",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	e.GET("/login").WithQuery("user_id", "firstUser").WithQuery("password", "example").
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("data")
	e.GET("/tasks").WithQuery("created_date", "2020-06-29").
		Expect().
		Status(http.StatusUnauthorized)
}

func TestToDoService_getAuthToken(t *testing.T) {
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *ToDoService
		inspect func(r *ToDoService, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			receiver.getAuthToken(tArgs.resp, tArgs.req)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

		})
	}
}

func TestToDoService_listTasks(t *testing.T) {
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *ToDoService
		inspect func(r *ToDoService, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			receiver.listTasks(tArgs.resp, tArgs.req)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

		})
	}
}

func TestToDoService_addTask(t *testing.T) {
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *ToDoService
		inspect func(r *ToDoService, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			receiver.addTask(tArgs.resp, tArgs.req)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

		})
	}
}

func Test_value(t *testing.T) {
	type args struct {
		req *http.Request
		p   string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 sql.NullString
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := value(tArgs.req, tArgs.p)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("value got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestToDoService_createToken(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *ToDoService
		inspect func(r *ToDoService, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1      string
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			name: "success",
			init: func(*testing.T) *ToDoService {
				return &ToDoService{}
			},
			args: func(*testing.T) args {
				return args{
					id: "juliette",
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1, err := receiver.createToken(tArgs.id)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !(len(got1) > 0) {
				t.Errorf("ToDoService.createToken got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ToDoService.createToken error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestToDoService_validToken(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *ToDoService
		inspect func(r *ToDoService, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1 *http.Request
		want2 bool
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1, got2 := receiver.validToken(tArgs.req)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToDoService.validToken got1 = %v, want1: %v", got1, tt.want1)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ToDoService.validToken got2 = %v, want2: %v", got2, tt.want2)
			}
		})
	}
}
