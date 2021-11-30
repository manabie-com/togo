package transport_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	mocks "github.com/manabie-com/togo/internal/mock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAuthToken(t *testing.T) {
	mockUcase := new(mocks.MockedTaskUsecase)
	mockUcase.On("ValidateUser", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true, nil)
	user := "user"
	password := "password"
	endpoint := fmt.Sprintf("/login?user_id=%s&password=%s", user, password)
	req, err := http.NewRequest("GET", endpoint, nil)
	handler := transport.NewTaskHandler(mockUcase)
	rec := httptest.NewRecorder()
	err = handler.GetAuthToken(rec, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUcase.AssertExpectations(t)
}

func TestListTasks(t *testing.T) {
	var mockTask storages.Task
	err := faker.FakeData(&mockTask)
	assert.NoError(t, err)
	mockUcase := new(mocks.MockedTaskUsecase)
	mockListTask := make([]storages.Task, 0)
	mockListTask = append(mockListTask, mockTask)
	mockUcase.On("ListTasks", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(mockListTask, nil)
	createdDate := time.Now().Format("2006-01-02")
	req, err := http.NewRequest("GET", "/tasks?created_date="+createdDate, nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	handler := transport.NewTaskHandler(mockUcase)
	err = handler.ListTasks(rec, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUcase.AssertExpectations(t)
}
func TestAddTask(t *testing.T) {
	mockTask := storages.Task{
		Content:     "Reading book",
		CreatedDate: time.Now(),
		UserID:      "first user",
	}
	tempMockTask := mockTask
	mockUcase := new(mocks.MockedTaskUsecase)

	js, err := json.Marshal(tempMockTask)
	assert.NoError(t, err)

	mockUcase.On("AddTask", mock.Anything, mock.AnythingOfType("*storages.Task")).Return(nil)
	mockUcase.On("CountTaskPerDay", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(uint8(0), nil)
	req, err := http.NewRequest("POST", "/tasks", strings.NewReader(string(js)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := transport.NewTaskHandler(mockUcase)

	err = handler.AddTask(rec, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUcase.AssertExpectations(t)
}
