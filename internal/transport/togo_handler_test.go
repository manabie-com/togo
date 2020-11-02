package transport_test

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const jwt_key = "wqGyEBBfPK9w3Lxw"

var (
	mockTogoHandler transport.TogoHandler
	task            []entities.Task
	now             = time.Now().Format("2006-01-02")
)

func TestLogin(t *testing.T) {
	var tests = []struct {
		url            string
		jsonRequest    []byte
		expectedStatus int
	}{
		{
			"/login",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login/",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login?notexsitusername",
			[]byte(`{"user_id": "manabie","password": "example"}`),
			401,
		},
		{
			"/login?notexsitusername",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login?",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login",
			[]byte(`{"user_id": "firstUser","password": "wrong pass"}`),
			401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer(tt.jsonRequest))
			rec := httptest.NewRecorder()
			http.HandlerFunc(mockTogoHandler.GetAuthToken).ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

//TestListContent only test the ListTasks, not included authMiddleware(check token)
func TestListContent(t *testing.T) {
	var tests = []struct {
		url            string
		userID         string
		expectedStatus int
	}{
		{
			fmt.Sprintf("/tasks?created_date=%s", now),
			"firstUser",
			200,
		},
		{
			fmt.Sprintf("/tasks"),
			"firstUser",
			400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			req.WithContext(context.WithValue(req.Context(), int8(0), tt.userID))
			rec := httptest.NewRecorder()
			http.HandlerFunc(mockTogoHandler.ListTasks).ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
func TestAddTask(t *testing.T) {
	var tests = []struct {
		url            string
		jsonRequest    []byte
		userID         string
		expectedStatus int
	}{
		{
			"/tasks",
			[]byte(`{"content": "this is task content"}`),
			"firstUser",
			200,
		},
		{
			"/tasks",
			[]byte(`{"contentt": "this is task content"}`),
			"firstUser",
			400,
		},
		//test rate limit
		{
			"/tasks",
			[]byte(`{"content": "this is task content"}`),
			"firstUser",
			403,
		},
	}
	for idx, tt := range tests {
		if idx == len(tests)-1 {
			// add max_todo task to make sure exceeding the rate litmit
			task = append(task, generateSliceTask(5, tt.userID)...)
		}
		t.Run(tt.url, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer([]byte(tt.jsonRequest)))
			req.WithContext(context.WithValue(req.Context(), int8(0), tt.userID))
			rec := httptest.NewRecorder()
			http.HandlerFunc(mockTogoHandler.AddTask).ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
func TestCreateToken(t *testing.T) {
	token, err := mockTogoHandler.CreateToken("firstUser")
	require.NoError(t, err)
	require.NotEmpty(t, token)
}
func TestValidToken(t *testing.T) {
	// generate valid token for Authorization header
	ValidToken, err := mockTogoHandler.CreateToken("firstUser")
	require.NoError(t, err)
	require.NotEmpty(t, ValidToken)

	var tests = []struct {
		token         string
		expectedValid bool
	}{
		{
			ValidToken,
			true,
		},
		{
			"123456",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			// set value for Authorization header
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tt.token)
			_, isValid := mockTogoHandler.ValidToken(req)
			require.Equal(t, tt.expectedValid, isValid)
		})
	}

}
func TestMain(m *testing.M) {
	mockTogoHandler = transport.TogoHandler{TogoUsecase: togoUsecaseMock{}, JWTKey: jwt_key}
	code := m.Run()
	os.Exit(code)
}

type togoUsecaseMock struct{}

func (tmock togoUsecaseMock) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]entities.Task, error) {
	if createdDate.String != now {
		return []entities.Task{}, errors.New("Wrong date format")
	}
	return task, nil
}
func (tmock togoUsecaseMock) AddTask(ctx context.Context, t entities.Task) error {
	task = append(task, t)
	return nil
}
func (tmock togoUsecaseMock) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return userID.String == "firstUser" && pwd.String == "example"
}
func (tmock togoUsecaseMock) GetMaxTaskTodo(ctx context.Context, userID string) (int, error) {
	return 5, nil
}
func generateSliceTask(num int, userID string) []entities.Task {
	var tasks []entities.Task
	for i := 0; i < num; i++ {
		tasks = append(tasks, entities.Task{
			ID:          "e1da0b9b-7ecc-44f9-82ff-4623cc504401",
			Content:     "content",
			UserID:      userID,
			CreatedDate: now,
		})
	}
	return tasks
}
