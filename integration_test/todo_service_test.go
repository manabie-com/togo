package integration_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/utils/test_constants"
	"github.com/manabie-com/togo/utils/test_utils"
	_ "github.com/mattn/go-sqlite3"
)

func setupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todo_service_test.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT NOT NULL, password TEXT NOT NULL, max_todo INTEGER DEFAULT 5 NOT NULL, CONSTRAINT users_PK PRIMARY KEY (id))")
	if err != nil {
		log.Fatal("error creating users table", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id TEXT NOT NULL, content TEXT NOT NULL, user_id TEXT NOT NULL, created_date TEXT NOT NULL, CONSTRAINT tasks_PK PRIMARY KEY (id), CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id))")
	if err != nil {
		log.Fatal("error creating tasks table", err)
	}

	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal("error deleting all rows from users", err)
	}

	_, err = db.Exec("DELETE FROM tasks")
	if err != nil {
		log.Fatal("error deleting all rows from tasks", err)
	}

	_, err = db.Exec("INSERT INTO users (id, password, max_todo) VALUES(?, ?, ?)",
		test_constants.UserName, test_constants.Password, test_constants.DefaultMaxToDo)
	if err != nil {
		log.Fatal("error inserting user", err)
	}

	return db, err
}

func getValidToken(t *testing.T, service *services.ToDoService) string {
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

func validateRetrievedTasks(t *testing.T, actualTask, expectedTask *storages.Task) {
	if _, err := uuid.Parse(actualTask.ID); err != nil {
		t.Errorf("%v", err)
	}

	if actualTask.Content != expectedTask.Content {
		t.Errorf("handler returned unexpected content: actual: %v, expected: %v",
			actualTask.Content, expectedTask.Content)
	}

	if actualTask.UserID != expectedTask.UserID {
		t.Errorf("handler returned unexpected userID: actual: %v, expected: %v",
			actualTask.UserID, expectedTask.UserID)
	}

	if actualTask.CreatedDate != expectedTask.CreatedDate {
		t.Errorf("handler returned unexpected createdDate: actual: %v, expected: %v",
			actualTask.CreatedDate, expectedTask.CreatedDate)
	}
}

func TestCorrectCredentialReturnsAuthToken(t *testing.T) {
	db, err := setupDatabase()

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
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
	db, err := setupDatabase()

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/login?user_id=wrongUser&password=wrongPass", nil)
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
	db, err := setupDatabase()

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
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
	db, err := setupDatabase()

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
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
	db, err := setupDatabase()

	expectedTask := storages.Task{
		ID:          "123",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: test_constants.CreatedDate,
	}

	_, err = db.Exec("INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)",
		expectedTask.ID, expectedTask.Content, expectedTask.UserID, expectedTask.CreatedDate)
	if err != nil {
		log.Fatal("error inserting task", err)
	}

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
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

func TestRetrieveNoTasks(t *testing.T) {
	db, err := setupDatabase()

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
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

	if len(retrievedTasks.Data) != 0 {
		t.Errorf("handler returned unexpected body: actual: %v expected: []",
			retrievedTasks.Data)
	}
}

func TestCreatingTasksAfterLimit(t *testing.T) {
	db, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	service := services.ToDoService{
		JWTKey: test_constants.TestJWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}

	expectedTasks := make([]storages.Task, test_constants.DefaultMaxToDo+1)
	timeNow := time.Now().Format("2006-01-02")

	for i := range expectedTasks {
		expectedTasks[i] = storages.Task{
			ID:          strconv.Itoa(i),
			Content:     "Test Content " + strconv.Itoa(i),
			UserID:      test_constants.UserName,
			CreatedDate: timeNow,
		}
	}

	token := getValidToken(t, &service)

	for _, expectedTask := range expectedTasks {
		jsonStr := []byte(`{"content":"` + expectedTask.Content + `"}`)
		req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set(test_constants.HeaderAuthorizationKey, token)

		resp := httptest.NewRecorder()
		service.ServeHTTP(resp, req)

		var actualTask test_utils.CreatedTask
		json.Unmarshal([]byte(resp.Body.String()), &actualTask)

		if actualTask.Data == nil {
			continue
		}

		validateRetrievedTasks(t, actualTask.Data, &expectedTask)
	}

	req, err := http.NewRequest(http.MethodGet, "/tasks?created_date="+timeNow, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set(test_constants.HeaderAuthorizationKey, token)

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
			expectedTasks)
	}

	for i, actualTask := range retrievedTasks.Data {
		validateRetrievedTasks(t, actualTask, &expectedTasks[i])
	}
}
