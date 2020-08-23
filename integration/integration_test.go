package integration

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/phuwn/togo/internal/services"
	"github.com/phuwn/togo/internal/storages/database"
	"github.com/phuwn/togo/util"

	_ "github.com/lib/pq"
)

var (
	pgSourceTest = "postgres://admin:password@localhost:15423/togo_test?sslmode=disable"
	jwtKey       = "wqGyEBBfPK9w3Lxw"

	// token generated at 2020-08-20 20:34:58
	validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTc5NTY1OTgsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.jfKWtsLnc-nq0E4sC7WKBSGwI5MpgdaeuxF7zw_vbZQ"
)

// TestMain - preparation for the integration test
func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", pgSourceTest)
	if err != nil {
		log.Fatal("error opening test db", err)
	}
	dbConn = db
	err = prepareSeedData()
	if err != nil {
		log.Fatal("error preparing seed data", err)
	}
	util.MockRuntimeFunc()

	os.Exit(m.Run())
}

func TestGetAuthToken(t *testing.T) {
	err := refreshDB()
	if err != nil {
		t.Errorf("error while refresh db data: %s", err.Error())
		return
	}

	s := &services.ToDoService{
		JWTKey: jwtKey,
		Store:  &database.Storage{dbConn},
	}

	type args struct {
		userID   string
		password string
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			"happy case",
			args{"firstUser", "example"},
			http.StatusOK,
			fmt.Sprintf(`{"data":"%v"}`+"\n", validToken),
		},
		{
			"wrong userID/password",
			args{"", ""},
			http.StatusUnauthorized,
			`{"error":"incorrect user_id/pwd"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/login", nil)
			if err != nil {
				t.Errorf("failed to create request, err: %s\n", err.Error())
				return
			}
			q := req.URL.Query()
			q.Add("user_id", tt.args.userID)
			q.Add("password", tt.args.password)
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()
			s.ServeHTTP(rr, req)
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

func TestListTasks(t *testing.T) {
	err := refreshDB()
	if err != nil {
		t.Errorf("error while refresh db data: %s", err.Error())
		return
	}

	s := &services.ToDoService{
		JWTKey: jwtKey,
		Store:  &database.Storage{dbConn},
	}

	type args struct {
		createdDate string
		token       string
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			"happy case",
			args{"2020-08-20", validToken},
			http.StatusOK,
			`{"data":[{"id":"3f44bbd3-8550-4a7c-a654-4495c060c36d","content":"content_1","user_id":"firstUser","created_date":"2020-08-20"},{"id":"498f276a-9145-4892-a480-f106b4708240","content":"content_2","user_id":"firstUser","created_date":"2020-08-20"},{"id":"a5a2ad9f-9472-4b02-a57e-659d0d561a0f","content":"content_3","user_id":"firstUser","created_date":"2020-08-20"},{"id":"af2760a4-da6f-402b-9339-856c287b66a1","content":"content_4","user_id":"firstUser","created_date":"2020-08-20"}]}` + "\n",
		},
		{
			"empty_task day case",
			args{"2020-08-19", validToken},
			http.StatusOK,
			`{"data":null}` + "\n",
		},
		{
			"invalid token case",
			args{"", ""},
			http.StatusUnauthorized,
			`{"error":"invalid token"}` + "\n",
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
			req.Header.Set("Authorization", tt.args.token)

			rr := httptest.NewRecorder()
			s.ServeHTTP(rr, req)
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

func TestAddTask(t *testing.T) {
	err := refreshDB()
	if err != nil {
		t.Errorf("error while refresh db data: %s", err.Error())
		return
	}

	s := &services.ToDoService{
		JWTKey: jwtKey,
		Store:  &database.Storage{dbConn},
	}

	type args struct {
		body  []byte
		token string
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			"happy case",
			args{[]byte(`{"content":"example_content"}`), validToken},
			http.StatusCreated,
			`{"data":{"id":"af1c772f-9abd-4e3c-94af-80d57d262028","content":"example_content","user_id":"firstUser","created_date":"2020-08-20"}}` + "\n",
		},
		{
			"task_create limit reach error",
			args{[]byte(`{"content":"example_content_1"}`), validToken},
			http.StatusTooManyRequests,
			fmt.Sprintf(`{"error":"%v"}`+"\n", services.CreateTaskLimitErrResp),
		},
		{
			"no content provided",
			args{nil, validToken},
			http.StatusBadRequest,
			fmt.Sprintf(`{"error":"%v"}`+"\n", services.InvalidBodyErrResp),
		},
		{
			"invalid token case",
			args{nil, ""},
			http.StatusUnauthorized,
			`{"error":"invalid token"}` + "\n",
		},
	}

	for _, tt := range tests {
		body := bytes.Buffer{}
		body.Write(tt.args.body)
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/tasks", &body)
			if err != nil {
				t.Errorf("failed to create request, err: %s\n", err.Error())
				return
			}
			req.Header.Set("Authorization", tt.args.token)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			s.ServeHTTP(rr, req)
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
