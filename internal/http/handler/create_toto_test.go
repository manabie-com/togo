package handler_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/chi07/todo/internal/http/handler"
	"github.com/chi07/todo/internal/model"
	"github.com/chi07/todo/tests/mocks"
)

func TestCreateTaskHandler_ServeHTTP(t *testing.T) {
	t.Run("without auth user", func(t *testing.T) {
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`abc-xyz`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		ctx := req.Context()
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":403, "message":"invalid user"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("when missing required param", func(t *testing.T) {
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`{"test": "abc"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("auth-user-id", "0a1a88c2-18f9-42dd-b4b0-99ee4dc77751")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":400, "message":"Key: 'CreateTaskRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("when missing required param", func(t *testing.T) {
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`{"test": "abc"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("auth-user-id", "0a1a88c2-18f9-42dd-b4b0-99ee4dc77751")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":400, "message":"Key: 'CreateTaskRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("when create task for user who is not setting limit", func(t *testing.T) {
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`{"title": "task1"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("auth-user-id", "0a1a88c2-18f9-42dd-b4b0-99ee4dc77721")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		createTaskMockService.On("CreateTask", mock.Anything, mock.Anything).Return(nil, model.ErrorNotFound)

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":400, "message":"no record found"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("when user reach max limitation tasks", func(t *testing.T) {
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`{"title": "task1"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("auth-user-id", "0a1a88c2-18f9-42dd-b4b0-99ee4dc77721")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		createTaskMockService.On("CreateTask", mock.Anything, mock.Anything).Return(nil, model.ErrorNotAllowed)

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":403, "message":"permission denied, you reach maximum posts a day"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("when create user got internal error", func(t *testing.T) {
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`{"title": "task1"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("auth-user-id", "0a1a88c2-18f9-42dd-b4b0-99ee4dc77721")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		createTaskMockService.On("CreateTask", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":500, "message":"internal error"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("success", func(t *testing.T) {
		taskID := uuid.New()
		createTaskMockService := new(mocks.CreateTaskService)
		h := handler.NewCreateTaskHandler(createTaskMockService)

		body := bytes.NewBufferString(`{"title": "task1"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("auth-user-id", "0a1a88c2-18f9-42dd-b4b0-99ee4dc77721")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		createTaskMockService.On("CreateTask", mock.Anything, mock.Anything).Return(taskID, nil)

		h.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		createTaskMockService.AssertExpectations(t)
		expectedResponse := fmt.Sprintf(`{"data":{"taskID":"%s"}}`, taskID)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
