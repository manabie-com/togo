package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlc "github.com/roandayne/togo/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initialize() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var jsonString = []byte(`{"title": "Test","content": "API Test","is_complete": true,"fullname": "Roan Dino"}`)
	r := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(jsonString))
	r.Header.Set("Content-Type", "application/json")
	fmt.Println(w.Body.String())

	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(w, r)

	return w
}

func TestCreateTaskHandler(t *testing.T) {
	w := initialize()
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"title":"Test","content":"API Test","is_complete":true,"fullname":"Roan Dino"}`
	assert.Contains(t, w.Body.String(), expected, "it should contain expected")
}

func TestMaximumNumberOfTask(t *testing.T) {
	arg := sqlc.CreateUserParams{
		FullName: "Test Account",
		Maximum:  3,
	}

	testQueries.CreateUser(context.Background(), arg)

	user, err := testQueries.GetUserByName(context.Background(), "Test Account")

	require.NoError(t, err)
	require.NotEmpty(t, user)
	assert.Equal(t, int32(3), user.Maximum)

	tArg := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}
	tArg2 := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}
	tArg3 := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}

	testQueries.CreateTask(context.Background(), tArg)
	testQueries.CreateTask(context.Background(), tArg2)
	testQueries.CreateTask(context.Background(), tArg3)

	w := initialize()
	expected := "You have reached the maximum allowed tasks for today!"
	assert.Contains(t, w.Body.String(), expected, "it should contain expected")
}

func TestCreateTask(t *testing.T) {
	user, _ := testQueries.GetUserByName(context.Background(), "Test Account")

	tArg := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}
	task, err := testQueries.CreateTask(context.Background(), tArg)

	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, tArg.Title, task.Title)
	require.Equal(t, tArg.Content, task.Content)
	require.Equal(t, tArg.IsComplete, task.IsComplete)
	require.Equal(t, tArg.UserID, task.UserID)

	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)
}
