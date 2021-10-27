package services

import (
	"bytes"
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

	"github.com/julienschmidt/httprouter"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	jwtKey   = "wqGyEBBfPK9w3Lxw"
	email    = "me@here.com"
	password = "password"
	userID   = 1
)

var token string

func TestHandler(t *testing.T) {
	//for Login API tests
	t.Run("login", testLoginHandlerIncorrectHTTPMethod)

	// //for Task API tests
	// taskList := &storages.Task{
	// 	Content:   "hash password",
	// 	UserID:    userID,
	// 	CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	// }
	// t.Run("listTasks", func(t *testing.T) {
	// 	testListTasksHandler(t, taskList)
	// })
	// task := &storages.Task{
	// 	Content:   "hash password",
	// 	UserID:    userID,
	// 	CreatedAt: time.Now(),
	// }
	// t.Run("addTask", func(t *testing.T) {
	// 	testAddTaskHandler(t, task)
	// })
	// tasktoUpdate := &storages.Task{
	// 	ID:      1,
	// 	Content: "update content",
	// }
	// t.Run("updateTask", func(t *testing.T) {
	// 	testUpdateTaskHandler(t, tasktoUpdate)
	// })
	//test id
	tasktoDelete := &storages.Task{
		ID: 1,
	}
	t.Run("deleteTask", func(t *testing.T) {
		testDeleteTaskHandler(t, tasktoDelete.ID)
	})
}

//Testing Login API with HTTP Methods
func testLoginHandlerIncorrectHTTPMethod(t *testing.T) {

	formData := url.Values{}
	formData.Add("email", email)
	formData.Add("password", password)

	testCases := map[string]struct {
		method string
		body   io.Reader
		code   int
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
		"put": {
			http.MethodPut,
			strings.NewReader(formData.Encode()),
			http.StatusMethodNotAllowed,
		},
		"delete": {
			http.MethodDelete,
			strings.NewReader(formData.Encode()),
			http.StatusMethodNotAllowed,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			request, err := http.NewRequest(tc.method, "/login", tc.body)
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			taskService := new(mocks.TaskService)
			taskService.On("ValidateUser", value(request, "email"), value(request, "password")).Return(true, nil)
			s := NewToDoService(jwtKey, taskService)
			router := httprouter.New()
			router.HandlerFunc(tc.method, "/login", s.getAuthToken)
			corsRouter := s.enableCORS(router)
			corsRouter.ServeHTTP(rr, request)
			assert.Equal(t, tc.code, rr.Code)
			if rr.Code == http.StatusOK {
				var loginResp map[string]string
				require.NoError(t, json.NewDecoder(rr.Body).Decode(&loginResp))
				token = loginResp["data"]
			}
		})
	}
}
func testListTasksHandler(t *testing.T, task *storages.Task) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks?created_date=%s", time.Now().Format("2006-01-02")), nil)
	require.NoError(t, err)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	taskService := new(mocks.TaskService)
	taskService.On(
		"RetrieveTasks",
		sql.NullString{
			String: email,
			Valid:  true,
		},
		sql.NullString{
			String: time.Now().Format("2006-01-02"),
			Valid:  true,
		},
	).Return([]*storages.Task{task}, nil)

	rr := httptest.NewRecorder()
	request.Header.Add("Authorization", token)

	s := NewToDoService(jwtKey, taskService)

	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/tasks", s.listTasks)
	corsRouter := s.enableCORS(router)
	corsRouter.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp map[string][]*storages.Task
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&resp))
	assert.Equal(t, 1, len(resp["data"]))
	assert.Equal(t, task.Content, resp["data"][0].Content)
	assert.Equal(t, task.UserID, resp["data"][0].UserID)
	assert.Equal(t, task.CreatedAt, resp["data"][0].CreatedAt)
}
func testAddTaskHandler(t *testing.T, task *storages.Task) {
	taskJson, err := json.Marshal(task)
	require.NoError(t, err)

	taskService := new(mocks.TaskService)
	taskService.On("AddTask", mock.AnythingOfType("*storages.Task"), email).Return(nil)

	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(taskJson))
	require.NoError(t, err)
	request.Header.Add("Authorization", token)

	s := NewToDoService(jwtKey, taskService)
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/tasks", s.addTask)
	corsRouter := s.enableCORS(router)
	corsRouter.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp map[string]*storages.Task
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&resp))
	assert.Equal(t, task.Content, resp["data"].Content)
	assert.Equal(t, task.UserID, resp["data"].UserID)
	//to get the date only
	dataCreated := resp["data"].CreatedAt
	assert.Equal(t, task.CreatedAt.Format("09-07-2017"), dataCreated.Format("09-07-2017"))
}

func testUpdateTaskHandler(t *testing.T, task *storages.Task) {
	taskJson, err := json.Marshal(task)
	require.NoError(t, err)

	taskService := new(mocks.TaskService)
	taskService.On("UpdateTask", mock.AnythingOfType("*storages.Task")).Return(nil)
	s := NewToDoService(jwtKey, taskService)

	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/task/update", s.updateTask)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/task/update", bytes.NewReader(taskJson))
	require.NoError(t, err)
	request.Header.Add("Authorization", token)
	corsRouter := s.enableCORS(router)
	corsRouter.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	assert.Equal(t, http.StatusOK, rr.Code)
	var updateResp map[string]string
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&updateResp))
	dataMessage := updateResp["data"]
	assert.Equal(t, "Successfully Updated", dataMessage)
}

func testDeleteTaskHandler(t *testing.T, id int) {
	taskJson, err := json.Marshal(&storages.Task{
		ID: id,
	})
	require.NoError(t, err)

	taskService := new(mocks.TaskService)
	taskService.On("DeleteTask", id).Return(nil)
	s := NewToDoService(jwtKey, taskService)

	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/task/delete", s.deleteTask)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/task/delete", bytes.NewReader(taskJson))
	require.NoError(t, err)
	request.Header.Add("Authorization", token)
	corsRouter := s.enableCORS(router)
	corsRouter.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	assert.Equal(t, http.StatusOK, rr.Code)
	var updateResp map[string]string
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&updateResp))
	dataMessage := updateResp["data"]
	assert.Equal(t, "Successfully Deleted", dataMessage)
}
