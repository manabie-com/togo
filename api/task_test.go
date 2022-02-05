package api

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	sqlc "github.com/roandayne/togo/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Init() *httptest.ResponseRecorder {
	database_url := "postgres://postgres:postgres@db:5432/todo_app?sslmode=disable"
	os.Setenv("DATABASE_URL", database_url)
	w := httptest.NewRecorder()
	var jsonString = []byte(`{"title": "Test","content": "API Test","is_complete": true,"username": "testaccount"}`)
	r := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(jsonString))
	r.Header.Set("Content-Type", "application/json")

	arg := sqlc.CreateUserParams{
		Username:       "testaccount",
		DailyTaskLimit: 3,
	}
	testQueries.CreateUser(context.Background(), arg)

	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(w, r)
	return w
}

func TestCreateTaskHandler(t *testing.T) {
	w := Init()

	expected := `{"title":"Test","content":"API Test","is_complete":true,"username":"testaccount"}`
	assert.Contains(t, w.Body.String(), expected, "it should contain expected")
}

func TestMaximumNumberOfTask(t *testing.T) {
	user, err := testQueries.GetUserByName(context.Background(), "testaccount")

	require.NoError(t, err)
	require.NotEmpty(t, user)
	assert.Equal(t, int32(3), user.DailyTaskLimit)

	ta := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}
	ta2 := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}
	ta3 := sqlc.CreateTaskParams{
		Title:      "test title",
		Content:    "test content",
		IsComplete: false,
		UserID:     int64(user.ID),
	}

	testQueries.CreateTask(context.Background(), ta)
	testQueries.CreateTask(context.Background(), ta2)
	testQueries.CreateTask(context.Background(), ta3)

	w := Init()
	expected := "You have reached the maximum limit of tasks today!"
	assert.Contains(t, w.Body.String(), expected, "it should contain expected")
}

func TestCreateTask(t *testing.T) {
	user, _ := testQueries.GetUserByName(context.Background(), "testaccount")

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
