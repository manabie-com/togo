package integration

import (
	"bytes"
	"encoding/json"
	"github.com/manabie-com/togo/internal/iservices"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	t.Run("Create task success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/login?user_id=2&password=example2", nil)
		require.Nil(t, err)
		w := httptest.NewRecorder()
		todoApi.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		var loginRes iservices.LoginResponse
		err = json.NewDecoder(w.Body).Decode(&loginRes)
		require.Nil(t, err)
		require.NotNil(t, loginRes.Data)

		reqBody := iservices.AddTaskRequest{
			Content: "Test something",
		}
		body, err := json.Marshal(reqBody)
		require.Nil(t, err)
		reqAddTask, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		require.Nil(t, err)
		reqAddTask.Header.Set("Authorization", loginRes.Data)
		wAddTask := httptest.NewRecorder()
		todoApi.ServeHTTP(wAddTask, reqAddTask)
		require.Equal(t, http.StatusOK, wAddTask.Code)
		var addTaskRes iservices.AddTaskResponse
		err = json.NewDecoder(wAddTask.Body).Decode(&addTaskRes)
		require.Nil(t, err)
		require.NotNil(t, addTaskRes.Data)
		require.Equal(t, reqBody.Content, addTaskRes.Data.Content)
		require.Equal(t, "2", addTaskRes.Data.UserID)
	})

	t.Run("Create task success and stick permission", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/login?user_id=3&password=example3", nil)
		require.Nil(t, err)
		w := httptest.NewRecorder()
		todoApi.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		var loginRes iservices.LoginResponse
		err = json.NewDecoder(w.Body).Decode(&loginRes)
		require.Nil(t, err)
		require.NotNil(t, loginRes.Data)

		reqBody := iservices.AddTaskRequest{
			Content: "Test something",
		}
		body, err := json.Marshal(reqBody)
		require.Nil(t, err)
		reqAddTask, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		require.Nil(t, err)
		reqAddTask.Header.Set("Authorization", loginRes.Data)
		wAddTask := httptest.NewRecorder()
		todoApi.ServeHTTP(wAddTask, reqAddTask)
		require.Equal(t, http.StatusOK, wAddTask.Code)
		var addTaskRes iservices.AddTaskResponse
		err = json.NewDecoder(wAddTask.Body).Decode(&addTaskRes)
		require.Nil(t, err)
		require.NotNil(t, addTaskRes.Data)
		require.Equal(t, reqBody.Content, addTaskRes.Data.Content)
		require.Equal(t, "3", addTaskRes.Data.UserID)

		wAddTask1 := httptest.NewRecorder()
		todoApi.ServeHTTP(wAddTask1, reqAddTask)
		require.Equal(t, http.StatusMethodNotAllowed, wAddTask1.Code)
	})

	t.Run("Create task fail by invalid token", func(t *testing.T) {
		reqBody := iservices.AddTaskRequest{
			Content: "Test something",
		}
		body, err := json.Marshal(reqBody)
		require.Nil(t, err)
		reqAddTask, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		require.Nil(t, err)
		reqAddTask.Header.Set("Authorization", "invalid token")
		wAddTask := httptest.NewRecorder()
		todoApi.ServeHTTP(wAddTask, reqAddTask)
		require.Equal(t, http.StatusInternalServerError, wAddTask.Code)
	})
}
