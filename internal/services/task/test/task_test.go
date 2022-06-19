package test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"togo/internal/services/task/api"
	"togo/internal/services/task/domain"
	"togo/pkg/random"
)

func createUser(t *testing.T) *domain.User {
	id, _ := uuid.NewUUID()
	dailyTaskLimit := random.RandomInt(1, 10)
	user := domain.NewUser(id, dailyTaskLimit)
	err := userRepo.Save(user)
	assert.NoError(t, err)
	return user
}
func TestCreateTask(t *testing.T) {
	t.Run("should create task", func(t *testing.T) {
		user := createUser(t)
		w := httptest.NewRecorder()
		title := random.RandomQuote()
		description := random.RandomQuote()
		dueDate := time.Now().Add(time.Hour * 24)
		server.ServeHTTP(w, createTaskHttp(user.ID.String(), title, description, dueDate))
		assert.Equal(t, 200, w.Code)
		var response CreateTaskResponse
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
		assert.Equal(t, user.ID.String(), response.Data.UserID.String())
		assert.Equal(t, title, response.Data.Title)
		assert.Equal(t, description, response.Data.Description)
	})
}

func createTaskHttp(userID, title, description string, dueDate time.Time) *http.Request {
	data := api.CreateTaskRequest{
		UserID:      userID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}
	requestBody, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
	return req
}

type CreateTaskResponse struct {
	Data domain.Task `json:"data"`
}
