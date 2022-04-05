package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	todoReq := dtos.CreateTodoRequest{
		Task: "task1",
	}
	j, err := json.Marshal(todoReq)
	assert.NoError(t, err)
	mockUCase := mocks.NewMockTodoUsecase()
	handler := NewTodoHandler(mockUCase)

	mockUCase.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	req, err := http.NewRequest(http.MethodPost, "/api/todo/create", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/todo/create", handler.Create)
	router.ServeHTTP(rec, req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
