package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/db"
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/routes"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var database *gorm.DB

var Router *mux.Router

func setup() {
	config.LoadEnv()

	conn := db.ConnectDB(true)

	database = conn

	db.Migrate(conn)

	db.Seed(conn)

	Router = routes.NewRouter(conn)
}

func tearDown() {
	defer db.DisconnectDB(database)
	database.Where("1 = 1").Delete(&models.Task{})
	database.Where("1 = 1").Delete(&models.User{})
}

func requestLogin(buffer []byte) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(buffer))

	response := httptest.NewRecorder()

	Router.ServeHTTP(response, request)

	return response
}

func requestAddTask(token string, buffer []byte) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(buffer))

	request.Header = map[string][]string{
		"Authorization": {token},
	}

	response := httptest.NewRecorder()

	Router.ServeHTTP(response, request)

	return response
}

func requestGetTasks(token string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", "/api/tasks", nil)

	request.Header = map[string][]string{
		"Authorization": {token},
	}

	response := httptest.NewRecorder()

	Router.ServeHTTP(response, request)

	return response
}

func getToken(buffer []byte) string {
	var m map[string]string

	json.Unmarshal(buffer, &m)

	return m["token"]
}

func getTasks(buffer []byte) []models.Task {
	var m []models.Task

	json.Unmarshal(buffer, &m)

	return m
}

func getTask(buffer []byte) models.Task {
	var m models.Task

	json.Unmarshal(buffer, &m)

	return m
}

func TestMain(m *testing.M) {
	setup()

	exitVal := m.Run()

	tearDown()

	os.Exit(exitVal)

}

func TestAPILogin_ok(t *testing.T) {
	user := map[string]string{"username": "huyha", "password": "123456"}

	b, _ := json.Marshal(user)

	response := requestLogin(b)

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var m map[string]string

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] != "" {
		t.Errorf("Got message '%s'", m["message"])
	}
}

func TestAPILogin_fail(t *testing.T) {
	user := map[string]string{"username": "huyha", "password": "wrongpass"}

	b, _ := json.Marshal(user)

	response := requestLogin(b)

	if status := response.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	var m map[string]string

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		t.Errorf("Got message '%s'", m["message"])
	}
}

func TestAPIAddTasksAndGetTasks(t *testing.T) {
	user := map[string]string{"username": "huyha", "password": "123456"}

	b, _ := json.Marshal(user)

	response := requestLogin(b)

	token := getToken(response.Body.Bytes())

	testCases := []struct {
		name         string
		task         map[string]string
		expectStatus int
	}{
		{name: "ADD TASK 1 is ok", task: map[string]string{"content": "ADD TASK 1"}, expectStatus: http.StatusCreated},
		{name: "ADD TASK 2 is ok", task: map[string]string{"content": "ADD TASK 2"}, expectStatus: http.StatusCreated},
		{name: "ADD TASK 3 is ok", task: map[string]string{"content": "ADD TASK 3"}, expectStatus: http.StatusCreated},
		{name: "ADD TASK 4 is fail", task: map[string]string{"content": "ADD TASK 4"}, expectStatus: http.StatusUnprocessableEntity},
	}

	for index, testCase := range testCases {
		b, _ = json.Marshal(testCase.task)

		response = requestAddTask(token, b)

		if status := response.Code; status != testCase.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, testCase.expectStatus)
		}

		newTask := getTask(response.Body.Bytes())

		if index != 3 {
			if newTask.Content != testCase.task["content"] {
				t.Errorf("Got message %q", newTask)
			}
		}

	}

	response = requestGetTasks(token)

	tasks := getTasks(response.Body.Bytes())

	if len(tasks) != 3 {
		t.Errorf("Got message %q", tasks)
	}
}
