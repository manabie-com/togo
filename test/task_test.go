package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/handler"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"github.com/huuthuan-nguyen/manabie/app/router"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"github.com/uptrace/bun"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1ODg0NTYwOX0.2dCHBNe0A1kbAriNUG-vevC71Sj6PKY9qthIh6DAR9Y"
var server *mux.Router
var db bun.IDB

// metaResponse /**
type metaResponse struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
}

// taskResponse /**
type taskResponse struct {
	metaResponse
	Data struct {
		ID            int       `json:"id"`
		Content       string    `json:"content"`
		PublishedDate string    `json:"published_date"`
		Status        int       `json:"status"`
		CreatedBy     int       `json:"created_by"`
		CreateAt      time.Time `json:"create_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	} `json:"data"`
}

// listResponse /**
type listResponse struct {
	metaResponse
	Data struct {
		Items []struct {
			ID            int       `json:"id"`
			Content       string    `json:"content"`
			PublishedDate string    `json:"published_date"`
			Status        int       `json:"status"`
			CreatedBy     int       `json:"created_by"`
			CreateAt      time.Time `json:"create_at"`
			UpdatedAt     time.Time `json:"updated_at"`
		} `json:"items"`
	} `json:"data"`
}

func truncateDB(ctx context.Context) error {
	_, err := db.ExecContext(ctx, "TRUNCATE TABLE users CASCADE")
	return err
}

func truncateTaskRecords(ctx context.Context) error {
	_, err := db.ExecContext(ctx, "TRUNCATE TABLE tasks")
	return err
}

func populateTasks(ctx context.Context) error {
	tasks := []struct {
		id            int
		content       string
		publishedDate time.Time
		status        int
		createdBy     int
		createdAt     time.Time
		updatedAt     time.Time
	}{
		{
			id:            1,
			content:       "Do homework",
			publishedDate: time.Now().UTC(),
			status:        0,
			createdBy:     1,
			createdAt:     time.Now().UTC(),
			updatedAt:     time.Now().UTC(),
		},
		{
			id:            2,
			content:       "Do laundry",
			publishedDate: time.Now().UTC(),
			status:        0,
			createdBy:     1,
			createdAt:     time.Now().UTC(),
			updatedAt:     time.Now().UTC(),
		},
		{
			id:            3,
			content:       "Do housework",
			publishedDate: time.Now().UTC(),
			status:        0,
			createdBy:     1,
			createdAt:     time.Now().UTC(),
			updatedAt:     time.Now().UTC(),
		},
	}
	for _, task := range tasks {
		if _, err := db.ExecContext(ctx, "INSERT INTO tasks(id, content, published_date, status, created_by, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)",
			task.id,
			task.content,
			task.publishedDate,
			task.status,
			task.createdBy,
			task.createdAt,
			task.updatedAt,
		); err != nil {
			return err
		}
	}
	return nil
}

func initUser(ctx context.Context) error {
	password, err := model.HashPassword("123@pass!word")
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx,
		"INSERT INTO users(id, email, password, is_active, daily_limit, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)",
		1,
		"huuthuan.nguyen@hotmail.com",
		password,
		true,
		3,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return err
}

func TestMain(m *testing.M) { // life cycle testing

	config := utils.ReadConfig() // read config from env
	c := context.Background()
	h := handler.New(c, config)
	server = router.NewRouter(config, h)

	ctx := context.Background()
	db = h.GetDB()
	// truncate
	if err := truncateDB(ctx); err != nil {
		log.Printf("truncate DB failed:%v\n", err)
		return
	}
	// init user
	if err := initUser(ctx); err != nil {
		log.Printf("init user failed:%v\n", err)
		return
	}

	code := m.Run()
	os.Exit(code)
}

// TestCreateTask /**
func TestCreateTask(t *testing.T) {

	var jsonStr = []byte(`{"content":"Do homework"}`)
	request, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusOK != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}

	var data taskResponse
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("Unable to unmarshal json response\n")
	}

	// assert Status on json body
	if data.Status != 1 {
		t.Errorf("Expected status is '%d'. Got '%d'\n", 1, data.Status)
	}

	// assert ID on json body
	if data.Data.ID == 0 {
		t.Errorf("Expected ID of task is greater than '0'. Got '%v'\n", data.Data.ID)
	}

	// assert Content on json body
	if data.Data.Content != "Do homework" {
		t.Errorf("Expected Content of Task to be 'Do homework'. Got '%v'\n", data.Data.Content)
	}
}

// TestCreateTaskReturnFailWhenUserExceededDailyLimit /**
func TestCreateTaskReturnFailWhenUserExceededDailyLimit(t *testing.T) {

	ctx := context.Background()
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records")
	}

	testData := []struct {
		input      string
		statusCode int
	}{
		{
			input:      `{"content":"Do homework"}`,
			statusCode: http.StatusOK,
		},
		{
			input:      `{"content":"Do laundry"}`,
			statusCode: http.StatusOK,
		},
		{
			input:      `{"content":"Do housework"}`,
			statusCode: http.StatusOK,
		},
		{
			input:      `{"content":"Get sleeping"}`,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, data := range testData {
		request, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer([]byte(data.input)))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)

		// assert status code
		if data.statusCode != responseRecorder.Code {
			t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
		}
	}
}

// TestListTasks /**
func TestListTasks(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("GET", "/api/tasks", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusOK != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}

	var data listResponse
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("Unable to unmarshal json response\n")
	}

	// count number of tasks
	if len(data.Data.Items) != 3 {
		t.Errorf("Expected number of tasks to be '3'. Got '%d'\n", len(data.Data.Items))
	}
}

// TestListTasks /**
func TestShowSpecificTask(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("GET", "/api/tasks/1", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusOK != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}

	var data taskResponse
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("Unable to unmarshal json response\n")
	}

	// count number of tasks
	if data.Data.Content != "Do homework" {
		t.Errorf("Expected content of task to be 'Do homework'. Got '%v'\n", data.Data.Content)
	}
}

// TestShowSpecificTaskReturnNotFoundWhenTaskDoesNotExist /**
func TestShowSpecificTaskReturnNotFoundWhenTaskDoesNotExist(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("GET", "/api/tasks/4", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusNotFound != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}
}

// TestUpdateSpecificTask /**
func TestUpdateSpecificTask(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("PUT", "/api/tasks/1", bytes.NewBuffer([]byte(`{"content":"Do homework later"}`)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusOK != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}

	var data taskResponse
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("Unable to unmarshal json response\n")
	}

	// count number of tasks
	if data.Data.Content != "Do homework later" {
		t.Errorf("Expected content of task to be 'Do homework later'. Got '%v'\n", data.Data.Content)
	}
}

// TestUpdateSpecificTaskReturnNotFoundWhenTaskDoesNotExist /**
func TestUpdateSpecificTaskReturnNotFoundWhenTaskDoesNotExist(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("PUT", "/api/tasks/4", bytes.NewBuffer([]byte(`{"content":"Do homework later"}`)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusNotFound != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}
}

// TestDeleteSpecificTask /**
func TestDeleteSpecificTask(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("DELETE", "/api/tasks/1", bytes.NewBuffer([]byte(`{"content":"Do homework later"}`)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusNoContent != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}
}

// TestDeleteSpecificTaskReturnNotFoundWhenTaskDoesNotExist /**
func TestDeleteSpecificTaskReturnNotFoundWhenTaskDoesNotExist(t *testing.T) {

	ctx := context.Background()
	// truncate all tasks
	if err := truncateTaskRecords(ctx); err != nil {
		t.Errorf("Unable to truncate task records\n")
	}
	// populate test tasks
	if err := populateTasks(ctx); err != nil {
		t.Errorf("Unable to populate task records\n")
	}

	request, _ := http.NewRequest("DELETE", "/api/tasks/4", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// assert status code
	if http.StatusNotFound != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}
}
