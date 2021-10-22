package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"

	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils/test_constants"
	"github.com/manabie-com/togo/utils/test_utils"
)

type TestDAL struct {
	Tasks       []*storages.Task
	Error       error
	IsValidUser bool
	MaxToDo     int
}

func (t *TestDAL) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	return t.Tasks, t.Error
}

func (t *TestDAL) RetrieveTasksCount(ctx context.Context, userID, createdDate sql.NullString) int {
	return len(t.Tasks)
}

func (t *TestDAL) AddTask(ctx context.Context, task *storages.Task) error {
	return t.Error
}

func (t *TestDAL) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return t.IsValidUser
}

func (t *TestDAL) GetMaxToDo(ctx context.Context, userID sql.NullString) int {
	return t.MaxToDo
}

func getValidToken(t *testing.T, service *ToDoService) string {

	loginReq, loginErr := http.NewRequest(http.MethodGet, test_constants.LoginUrl, nil)
	if loginErr != nil {
		t.Fatal(loginErr)
	}

	loginRr := httptest.NewRecorder()
	service.ServeHTTP(loginRr, loginReq)

	loginResponse := make(map[string]string)
	jsonErr := json.Unmarshal([]byte(loginRr.Body.String()), &loginResponse)

	if jsonErr != nil {
		t.Fatal(jsonErr)
	}

	return loginResponse["data"]
}

// NEW TEST CASE: other http methods other than GET in /login should fail
// NEW TEST CASE: json decoder error

func TestCorrectCredentialReturnsAuthToken(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			IsValidUser: true,
		},
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.LoginUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusOK
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	expected := "data"
	if !strings.Contains(resp.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: actual: %v expected: %v",
			resp.Body.String(), expected)
	}
}

func TestIncorrectCredentialsReturnsError(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			IsValidUser: false,
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/login?user_id="+test_constants.WrongUserName+"&password="+test_constants.WrongPassword, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusUnauthorized
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	expected := "error"
	if !strings.Contains(resp.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: actual: %v expected: %v",
			resp.Body.String(), expected)
	}
}

func TestUnauthorizedTasksRetrievalReturnsError(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			IsValidUser: false,
		},
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.GetTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(test_constants.HeaderAuthorizationKey, test_constants.InvalidToken)

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusUnauthorized
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}
}

func TestUnauthorizedTasksCreationReturnsError(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			IsValidUser: false,
		},
	}

	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(test_constants.HeaderAuthorizationKey, test_constants.InvalidToken)

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusUnauthorized
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}
}

func TestRetrieveTasks(t *testing.T) {
	expectedTask := storages.Task{
		ID:          "281",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: test_constants.CreatedDate,
	}

	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			Tasks: []*storages.Task{
				&expectedTask,
			},
			IsValidUser: true,
		},
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.GetTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set(test_constants.HeaderAuthorizationKey, getValidToken(t, &service))

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusOK
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	var retrievedTasks test_utils.RetrievedTasks
	json.Unmarshal([]byte(resp.Body.String()), &retrievedTasks)

	if retrievedTasks.Data == nil {
		t.Errorf("handler returned unexpected body: actual: nil expected: %v",
			expectedTask)
	}

	for _, actualTask := range retrievedTasks.Data {
		if expectedTask != *actualTask {
			t.Errorf("handler returned unexpected body: actual: %v expected: %v",
				*actualTask, expectedTask)
		}
	}
}

func TestInternalServerErrorWhenRetrievingTasks(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			Error:       errors.New("Error retrieving tasks"),
			IsValidUser: true,
		},
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.GetTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(test_constants.HeaderAuthorizationKey, getValidToken(t, &service))

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusInternalServerError
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	expected := "error"
	if !strings.Contains(resp.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: actual: %v expected: %v",
			resp.Body.String(), expected)
	}
}

func TestCreatingTaskUnderLimit(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			IsValidUser: true,
			MaxToDo:     2,
		},
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(test_constants.HeaderAuthorizationKey, getValidToken(t, &service))

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusOK
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	expectedTask := storages.Task{
		ID:          "001",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: time.Now().Format("2006-01-02"),
	}

	var actualTask test_utils.CreatedTask
	json.Unmarshal([]byte(resp.Body.String()), &actualTask)

	if _, err := uuid.Parse(actualTask.Data.ID); err != nil {
		t.Errorf("%v", err)
	}

	if actualTask.Data.Content != expectedTask.Content {
		t.Errorf("handler returned unexpected content: actual: %v, expected: %v",
			actualTask.Data.Content, expectedTask.Content)
	}

	if actualTask.Data.UserID != expectedTask.UserID {
		t.Errorf("handler returned unexpected userID: actual: %v, expected: %v",
			actualTask.Data.UserID, expectedTask.UserID)
	}

	if actualTask.Data.CreatedDate != expectedTask.CreatedDate {
		t.Errorf("handler returned unexpected createdDate: actual: %v, expected: %v",
			actualTask.Data.CreatedDate, expectedTask.CreatedDate)
	}
}

func TestInternalServerErrorWhenCreatingTask(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			Error:       errors.New("Error creating task"),
			IsValidUser: true,
			MaxToDo:     2,
		},
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(test_constants.HeaderAuthorizationKey, getValidToken(t, &service))

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusInternalServerError
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	expected := "error"
	if !strings.Contains(resp.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: actual: %v expected: %v",
			resp.Body.String(), expected)
	}
}

func TestCreatingTaskAboveLimit(t *testing.T) {
	service := ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &TestDAL{
			IsValidUser: true,
			MaxToDo:     0,
		},
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(test_constants.HeaderAuthorizationKey, getValidToken(t, &service))

	resp := httptest.NewRecorder()
	service.ServeHTTP(resp, req)

	expectedStatus := http.StatusBadRequest
	if status := resp.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: actual: %v expected: %v",
			status, expectedStatus)
	}

	expectedKeyError := "error"
	if !strings.Contains(resp.Body.String(), expectedKeyError) {
		t.Errorf("handler returned unexpected body: actual: %v expected: %v",
			resp.Body.String(), expectedKeyError)
	}
}
