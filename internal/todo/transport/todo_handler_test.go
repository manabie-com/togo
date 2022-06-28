package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"manabieAssignment/internal/mocks"
	"manabieAssignment/internal/todo/transport/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func MockJsonPost(c *gin.Context, content []byte) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	c.Request.Body = io.NopCloser(bytes.NewBuffer(content))
}

func TestTodoHandler_CreateTodo_Success(t *testing.T) {
	mockTodoUseCase := new(mocks.TodoUseCase)
	mockTodoUseCase.On("CreateTodo", mock.Anything).Return(uint(1), nil)

	todoModel := model.TodoModel{
		UserID:  1,
		Name:    "todo_name",
		Content: "todo_content",
	}
	todoModelRequestBody, err := json.Marshal(todoModel)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	request, err := http.NewRequest(http.MethodPost, "/todo", strings.NewReader(string(todoModelRequestBody)))
	require.NoError(t, err)

	NewTodoHandler(r, mockTodoUseCase)
	r.ServeHTTP(w, request)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestTodoHandler_CreateTodo_CannotBindRequestBody(t *testing.T) {
	mockTodoUseCase := new(mocks.TodoUseCase)
	mockTodoUseCase.On("CreateTodo", mock.Anything).Return(uint(1), nil)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	request, err := http.NewRequest(http.MethodPost, "/todo", nil)
	require.NoError(t, err)

	NewTodoHandler(r, mockTodoUseCase)
	r.ServeHTTP(w, request)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTodoHandler_CreateTodo_TodoUseCaseReturnError(t *testing.T) {
	mockTodoUseCase := new(mocks.TodoUseCase)
	mockTodoUseCase.On("CreateTodo", mock.Anything).Return(uint(0), errors.New("something"))

	todoModel := model.TodoModel{
		UserID:  1,
		Name:    "todo_name",
		Content: "todo_content",
	}
	todoModelRequestBody, err := json.Marshal(todoModel)
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, "/todo", strings.NewReader(string(todoModelRequestBody)))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	NewTodoHandler(r, mockTodoUseCase)
	r.ServeHTTP(w, request)
	require.Equal(t, http.StatusBadRequest, w.Code)
}
