package services

import (
	"github.com/manabie-com/togo/internal/storages"

	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

var authUser = storages.User{
	ID:       "firstUser",
	Password: "example",
}

const CreatedDate string = "2021-01-27"

func createRandomTask(t *testing.T, createdDate string) {
	param := &storages.Task{
		ID:          faker.FirstName(),
		Content:     faker.Sentence(2),
		UserID:      authUser.ID,
		CreatedDate: createdDate,
	}

	err := TestTodoService.TaskUsecase.AddTask(context.Background(), param)
	assert.NoError(t, err)
}

func login(t *testing.T) string {
	res := struct {
		Data string `json:"data"`
	}{}
	req := clientRequest{
		method: http.MethodGet,
		path:   "/login?user_id=firstUser&password=example",
		body:   nil,
		token:  "",
	}
	responseRecorder := apiClient(req)
	err := json.NewDecoder(responseRecorder.Body).Decode(&res)

	assert.NoError(t, err)
	assert.Equal(t, 200, responseRecorder.Code)
	assert.NotEmpty(t, res.Data)

	return res.Data
}

func TestLogin(t *testing.T) {
	Convey("Test Login", t, func() {
		login(t)
	})
}

// Wrong user_id and password
func TestLoginErr(t *testing.T) {
	Convey("Test Login Error", t, func() {
		req := clientRequest{
			method: http.MethodGet,
			path:   "/login?user_id=wrongID&password=wrongPassword",
			body:   nil,
			token:  "",
		}
		responseRecorder := apiClient(req)
		assert.Equal(t, 401, responseRecorder.Code)
	})
}

func TestAddTask(t *testing.T) {
	Convey("Test Add New Task", t, func() {
		token := login(t)
		mapBody := map[string]interface{}{
			"content": faker.Sentence(2),
		}
		body, _ := json.Marshal(mapBody)
		req := clientRequest{
			method: http.MethodPost,
			path:   "/tasks",
			body:   bytes.NewReader(body),
			token:  token,
		}
		responseRecorder := apiClient(req)
		res := struct {
			Data storages.Task `json:"data"`
		}{
			Data: storages.Task{},
		}
		err := json.NewDecoder(responseRecorder.Body).Decode(&res)

		assert.NoError(t, err)
		assert.Equal(t, mapBody["content"], res.Data.Content)
		assert.Equal(t, authUser.ID, res.Data.UserID)
		assert.NotEmpty(t, res.Data.ID)
		assert.NotEmpty(t, res.Data.CreatedDate)
	})
}

func TestGetTasks(t *testing.T) {
	Convey("Test Get Tasks", t, func() {
		// Seed tasks
		for i := 0; i < 10; i++ {
			createRandomTask(t, CreatedDate)
		}
		// Login
		token := login(t)

		path := fmt.Sprintf("/tasks?created_date=%s", CreatedDate)
		req := clientRequest{
			method: http.MethodGet,
			path:   path,
			token:  token,
		}
		responseRecorder := apiClient(req)

		res := struct {
			Data []*storages.Task `json:"data"`
		}{
			Data: []*storages.Task{},
		}

		err := json.NewDecoder(responseRecorder.Body).Decode(&res)
		assert.NoError(t, err)

		for _, task := range res.Data {
			assert.Equal(t, authUser.ID, task.UserID)
			assert.Equal(t, CreatedDate, task.CreatedDate)
			assert.NotEmpty(t, task.Content)
			assert.NotEmpty(t, task.ID)
		}
	})
}

// User's tasks exceed max task per day
func TestAddTaskErr(t *testing.T) {
	Convey("Test Add Task Error", t, func() {
		// Seed tasks
		now := time.Now()
		user, _ := TestTodoService.UserUsecase.GetUser(context.Background(), sql.NullString{
			String: authUser.ID,
			Valid:  true,
		})
		for i := 0; i < user.MaxTodo; i++ {
			createRandomTask(t, now.Format("2006-01-02"))
		}

		token := login(t)
		mapBody := map[string]interface{}{
			"content": faker.Sentence(2),
		}
		body, _ := json.Marshal(mapBody)
		req := clientRequest{
			method: http.MethodPost,
			path:   "/tasks",
			body:   bytes.NewReader(body),
			token:  token,
		}
		responseRecorder := apiClient(req)
		assert.Equal(t, 400, responseRecorder.Code)
	})
}

// Delete all task of user
func TestDeleteTasks(t *testing.T) {
	Convey("Test Delete Tasks", t, func() {
		token := login(t)
		req := clientRequest{
			method: http.MethodDelete,
			path:   "/tasks",
			token:  token,
		}
		responseRecorder := apiClient(req)
		res := struct {
			Success bool `json:"success"`
		}{
			Success: true,
		}
		err := json.NewDecoder(responseRecorder.Body).Decode(&res)

		assert.NoError(t, err)
		assert.Equal(t, true, res.Success)
	})
}
