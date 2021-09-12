package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/mocks"
)

const (
	jwtKey = "wqGyEBBfPK9w3Lxw"
	userID = "firstUser"
	password = "example"
)

var token string

func TestHandler(t *testing.T) {
	t.Run("login", testLoginHandler)

	task := &storages.Task{
		Content:     "hash password",
		UserID:      userID,
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	t.Run("addTask", func(t *testing.T) {
		testAddTaskHandler(t, task)
	})

	t.Run("listTasks", func(t *testing.T) {
		testListTasksHandler(t, task)
	})
}

func testLoginHandler(t *testing.T) {
	formData := url.Values{}
	formData.Add("user_id", userID)
	formData.Add("password", password)

	testCases := map[string]struct{
		method string
		body io.Reader
		code int
	}{
		"get": {
			http.MethodGet,
			nil,
			http.StatusMethodNotAllowed,
		},
		"post": {
			http.MethodPost,
			strings.NewReader(formData.Encode()),
			http.StatusOK,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			request, err := http.NewRequest(tc.method, "/login", tc.body)
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			taskService := new(mocks.TaskService)
			taskService.On("ValidateUser", context.Background(), value(request, "user_id"), value(request, "password")).Return(true, nil)

			s := NewToDoService(jwtKey, taskService, nil)
			s.ServeHTTP(rr, request)
			assert.Equal(t, tc.code, rr.Code)

			if rr.Code == http.StatusOK {
				var loginResp map[string]string
				require.NoError(t, json.NewDecoder(rr.Body).Decode(&loginResp))
				token = loginResp["data"]
			}
		})
	}
}

func testAddTaskHandler(t *testing.T, task *storages.Task) {
	taskJson, err := json.Marshal(task)
	require.NoError(t, err)

	taskService := new(mocks.TaskService)
	taskService.On("AddTask", mock.MatchedBy(func (_ context.Context) bool { return true }), mock.AnythingOfType("*storages.Task")).Return(nil)

	limiter := rate.NewLimiter(1, 1)
	userLimiter := new(mocks.UserLimiter)
	userLimiter.On("GetLimiter", userID).Return(limiter)

	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(taskJson))
	require.NoError(t, err)
	request.Header.Add("Authorization", token)

	s := NewToDoService(jwtKey, taskService, userLimiter)
	s.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp map[string]*storages.Task
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&resp))
	assert.Equal(t, task.Content, resp["data"].Content)
	assert.Equal(t, task.UserID, resp["data"].UserID)
	assert.Equal(t, task.CreatedDate, resp["data"].CreatedDate)
}

func testListTasksHandler(t *testing.T, task *storages.Task) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks?created_date=%s", time.Now().Format("2006-01-02")), nil)
	require.NoError(t, err)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println("user_id", request.FormValue("user_id"))

	taskService := new(mocks.TaskService)
	taskService.On(
		"RetrieveTasks",
		mock.MatchedBy(func (_ context.Context) bool { return true }),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: time.Now().Format("2006-01-02"),
			Valid:  true,
		},
	).Return([]*storages.Task{task}, nil)

	rr := httptest.NewRecorder()
	request.Header.Add("Authorization", token)

	s := NewToDoService(jwtKey, taskService, nil)
	s.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp map[string][]*storages.Task
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&resp))
	assert.Equal(t, 1, len(resp["data"]))
	assert.Equal(t, task.Content, resp["data"][0].Content)
	assert.Equal(t, task.UserID, resp["data"][0].UserID)
	assert.Equal(t, task.CreatedDate, resp["data"][0].CreatedDate)
}
