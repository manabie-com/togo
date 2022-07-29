package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"togo/constants"
	"togo/internal/response"
	"togo/internal/task/dto"
	"togo/internal/task/mocks"
	"togo/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/test-go/testify/assert"
)

func mockRequestCreateTask(userID int) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = validator.NewValidator()
	json := `{
		"description": "description",
		"ended_at": "2022-07-27T11:55:37+07:00"
	}`
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/users/%d/tasks", userID), strings.NewReader(json))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.Set(string(constants.ContextUserID), userID)

	return e, context, recorder
}

func TestTaskHandler_CreateTaskSuccess(t *testing.T) {
	userID := 1
	createTaskDto := &dto.CreateTaskDto{
		Description: "description",
		EndedAt:     time.Date(2022, 7, 27, 11, 55, 37, 0, time.Local),
	}
	response := &response.TaskResponse{
		ID:          1,
		UserID:      userID,
		Description: "description",
		EndedAt:     time.Date(2022, 7, 27, 11, 55, 37, 0, time.Local),
	}

	e, context, recorder := mockRequestCreateTask(userID)

	taskGroup := e.Group("/users/:id/tasks")
	service := mocks.NewTaskService(t)
	service.On("Create", createTaskDto, userID).Return(response, nil)

	handler := NewTaskHandler(taskGroup, service)

	err := handler.Create(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	response1 := `{"data":{"id":1,"user_id":1,"description":"description","ended_at":"2022-07-27T11:55:37+07:00"},"status":200}`
	response2 := `{"status":200,"data":{"id":1,"user_id":1,"description":"description","ended_at":"2022-07-27T11:55:37+07:00"}}`
	dataResponse := strings.Trim(recorder.Body.String(), "\n")
	assert.True(t, dataResponse == response1 || dataResponse == response2)
}

func TestUserHandler_HandleCreateUser_ServiceReturnError(t *testing.T) {
	userID := 1
	createTaskDto := &dto.CreateTaskDto{
		Description: "description",
		EndedAt:     time.Date(2022, 7, 27, 11, 55, 37, 0, time.Local),
	}
	e, context, recorder := mockRequestCreateTask(userID)

	taskGroup := e.Group("/users/:id/tasks")
	service := mocks.NewTaskService(t)
	mockError := errors.New("mock_err")
	service.On("Create", createTaskDto, userID).Return(nil, mockError)

	handler := NewTaskHandler(taskGroup, service)

	err := handler.Create(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	response1 := `{"data":"mock_err","status":400}`
	response2 := `{"status":400,"data":"mock_err"}`
	dataResponse := strings.Trim(recorder.Body.String(), "\n")
	assert.True(t, dataResponse == response1 || dataResponse == response2)
}
