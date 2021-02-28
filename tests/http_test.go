package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/manabie-com/togo/internal/db"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/services"
)

func getToken(userID string, password string) (*httptest.ResponseRecorder, string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/login?user_id=%s&password=%s", userID, password), nil)
	if err != nil {
		return nil, "", err
	}

	w := httptest.NewRecorder()
	httpHandler.ServeHTTP(w, req)

	var response struct {
		Data string `json:"data"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		return nil, "", err
	}

	return w, response.Data, nil
}

func addTask(token, content string) (*httptest.ResponseRecorder, error) {
	requestBody := storages.Task{
		Content: content,
	}
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(requestBody); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/tasks", &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", token)

	defer func() {
		if err := req.Body.Close(); err != nil {
			return
		}
	}()

	w := httptest.NewRecorder()
	httpHandler.ServeHTTP(w, req)

	return w, nil
}

func TestLogin(t *testing.T) {
	defer func() {
		if err := db.Truncate(database); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	users, err := db.SeedUsers(database)
	if err != nil {
		t.Fatalf("error seeding users: %v", err)
	}

	w, token, err := getToken(users[0].ID, users[0].Password)
	if err != nil {
		t.Fatalf("error gettting token: %v", err)
	}

	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(http.StatusOK, w.Code)
	assert.NotEmpty(token)
}

func TestLoginFail(t *testing.T) {
	defer func() {
		if err := db.Truncate(database); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	users, err := db.SeedUsers(database)
	if err != nil {
		t.Fatalf("error seeding users: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/login?user_id=%s&password=%s", users[0].ID, "wrongPass"), nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	w := httptest.NewRecorder()
	httpHandler.ServeHTTP(w, req)

	var response struct {
		Error string `json:"error"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)

	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(http.StatusUnauthorized, w.Code)
	assert.Equal("incorrect user_id/pwd", response.Error)
}

func TestListTasks(t *testing.T) {
	defer func() {
		if err := db.Truncate(database); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	users, err := db.SeedUsers(database)
	if err != nil {
		t.Fatalf("error seeding users: %v", err)
	}

	tasks, err := db.SeedTasks(database, users)
	if err != nil {
		t.Fatalf("error seeding tasks: %v", err)
	}

	_, token, err := getToken(users[0].ID, users[0].Password)
	if err != nil {
		t.Fatalf("error gettting token: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks?created_date=%s", time.Now().Format("2006-01-02")), nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	httpHandler.ServeHTTP(w, req)

	type responseBody struct {
		Data []storages.Task `json:"data"`
	}
	response := responseBody{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(len(response.Data), 1)
	assert.Equal(response.Data[0], tasks[0])
}

func TestAddTask(t *testing.T) {
	defer func() {
		if err := db.Truncate(database); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}

		httpHandler.TaskService.Redis.FlushAll(context.Background())
	}()

	users, err := db.SeedUsers(database)
	if err != nil {
		t.Fatalf("error seeding users: %v", err)
	}

	_, token, err := getToken(users[0].ID, users[0].Password)
	if err != nil {
		t.Fatalf("error gettting token: %v", err)
	}

	content := "task content"
	w, err := addTask(token, content)
	if err != nil {
		t.Fatalf("error adding task: %v", err)
	}

	type responseBody struct {
		Data storages.Task `json:"data"`
	}
	response := responseBody{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(http.StatusOK, w.Code)
	assert.NotEmpty(response.Data.ID)
	assert.Equal(users[0].ID, response.Data.UserID)
	assert.Equal(content, response.Data.Content)
	assert.Equal(time.Now().Format("2006-01-02"), response.Data.CreatedDate)
}

func TestAddTaskLimitExceeded(t *testing.T) {
	defer func() {
		if err := db.Truncate(database); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}

		httpHandler.TaskService.Redis.FlushAll(context.Background())
	}()

	users, err := db.SeedUsers(database)
	if err != nil {
		t.Fatalf("error seeding users: %v", err)
	}

	_, token, err := getToken(users[0].ID, users[0].Password)
	if err != nil {
		t.Fatalf("error gettting token: %v", err)
	}

	content := "task content"
	for i := 0; i < services.LIMIT_PER_DAY; i++ {
		_, err := addTask(token, content)
		if err != nil {
			t.Fatalf("error adding task: %v", err)
		}
	}

	w, err := addTask(token, content)
	if err != nil {
		t.Fatalf("error adding task: %v", err)
	}

	type responseBody struct {
		Error string `json:"error"`
	}
	response := responseBody{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(w.Code, http.StatusTooManyRequests)
	assert.Equal(response.Error, "daily tasks limit exceeded")
}
