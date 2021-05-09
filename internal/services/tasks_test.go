package services

import (
	"bytes"
	"context"
	"encoding/json"
	mocks2 "github.com/manabie-com/togo/internal/mocks/services"
	mocks "github.com/manabie-com/togo/internal/mocks/storages"
	"github.com/manabie-com/togo/internal/models"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestToDoService_AddTask_firstTime(t *testing.T) {
	task := &models.Task{
		Content:     "task",
		UserID:      "firstUser",
	}
	bData, err := json.Marshal(task)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, TASKS, bytes.NewBuffer(bData))
	resp := httptest.NewRecorder()
	
	service := &ToDoService{
		JWTKey:  "test",
		Store:   &mocks.IDatabase{},
		cache:   &mocks2.ICache{},
		maxTodo: 5,
	}
	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), task.UserID))
	// declare called mock's methods
	service.cache.(*mocks2.ICache).On("GetMaxTodo", task.UserID).Return(int32(-1), nil)
	service.cache.(*mocks2.ICache).On("SetMaxTodo", mock.Anything, service.maxTodo).Return(nil)
	service.cache.(*mocks2.ICache).On("GetNumberOfTasks", task.UserID, mock.Anything).Return(int32(-1), nil)
	service.cache.(*mocks2.ICache).On("SetNumberOfTasks", task.UserID, mock.Anything, int32(0)).Return(nil)

	service.Store.(*mocks.IDatabase).On("GetMaxTodo", task.UserID).Return(service.maxTodo, nil)
	service.Store.(*mocks.IDatabase).On("CountTasks", task.UserID, mock.Anything).Return(int32(0), nil)
	service.Store.(*mocks.IDatabase).On("AddTask", mock.Anything, mock.Anything).Return(nil)

	service.AddTask(resp, req)
	result := resp.Result()
	require.Equal(t, http.StatusOK, result.StatusCode)
}

func TestToDoService_AddTask_ReachLimit(t *testing.T) {
	task := &models.Task{
		Content:     "task",
		UserID:      "firstUser",
	}
	bData, err := json.Marshal(task)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, TASKS, bytes.NewBuffer(bData))
	resp := httptest.NewRecorder()

	service := &ToDoService{
		JWTKey:  "test",
		Store:   &mocks.IDatabase{},
		cache:   &mocks2.ICache{},
		maxTodo: 5,
	}
	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), task.UserID))
	// declare called mock's methods
	service.cache.(*mocks2.ICache).On("GetMaxTodo", task.UserID).Return(service.maxTodo, nil)
	service.cache.(*mocks2.ICache).On("GetNumberOfTasks", task.UserID, mock.Anything).Return(int32(5), nil)
	service.cache.(*mocks2.ICache).On("SetNumberOfTasks", task.UserID, mock.Anything, int32(0)).Return(nil)

	service.AddTask(resp, req)
	result := resp.Result()
	require.Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestToDoService_SignUp(t *testing.T) {
	user := &models.User{
		ID:       "firstUser",
		Password: "example",
	}
	bData, err := json.Marshal(user)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, SIGNUP, bytes.NewBuffer(bData))
	resp := httptest.NewRecorder()

	service := &ToDoService{
		JWTKey:  "test",
		Store:   &mocks.IDatabase{},
		cache:   &mocks2.ICache{},
		maxTodo: 5,
	}
	service.Store.(*mocks.IDatabase).On("AddUser", user.ID, user.Password, service.maxTodo).Return(nil)
	service.SignUp(resp, req)
	result := resp.Result()
	require.Equal(t, http.StatusOK, result.StatusCode)
	responseData := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(responseData)
	require.NotEmpty(t, responseData.Data)
}

func TestToDoService_GetAuthToken(t *testing.T) {
	user := &models.User{
		ID:       "firstUser",
		Password: "example",
	}
	bData, err := json.Marshal(user)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, LOGIN, bytes.NewBuffer(bData))
	resp := httptest.NewRecorder()

	service := &ToDoService{
		JWTKey:  "test",
		Store:   &mocks.IDatabase{},
		cache:   &mocks2.ICache{},
		maxTodo: 5,
	}
	service.Store.(*mocks.IDatabase).On("ValidateUser", user.ID, user.Password).Return(true)
	service.GetAuthToken(resp, req)

	result := resp.Result()
	require.Equal(t, http.StatusOK, result.StatusCode)
	responseData := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(responseData)
	require.NotEmpty(t, responseData.Data)
}

func TestToDoService_GetAuthToken_InvalidUsernamePwd(t *testing.T) {
	user := &models.User{
		ID:       "firstUser",
		Password: "example",
	}
	bData, err := json.Marshal(user)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, LOGIN, bytes.NewBuffer(bData))
	resp := httptest.NewRecorder()

	service := &ToDoService{
		JWTKey:  "test",
		Store:   &mocks.IDatabase{},
		cache:   &mocks2.ICache{},
		maxTodo: 5,
	}
	service.Store.(*mocks.IDatabase).On("ValidateUser", user.ID, user.Password).Return(false)
	service.GetAuthToken(resp, req)
	result := resp.Result()

	require.Equal(t, http.StatusBadRequest, result.StatusCode)
	responseData := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(responseData)
	require.Equal(t, InvalidUserPwd.Error(), responseData.Data.(string))
}

func TestToDoService_ListTasks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, TASKS, nil)
	resp := httptest.NewRecorder()

	service := &ToDoService{
		JWTKey:  "test",
		Store:   &mocks.IDatabase{},
		cache:   &mocks2.ICache{},
		maxTodo: 5,
	}
	userId := "firstUser"
	createdDate := "2021-05-08"
	tasks := []*models.Task{
		{
			ID:          "1",
			Content:     "task 1",
			UserID:      userId,
			CreatedDate: createdDate,
		},
		{
			ID:          "2",
			Content:     "task 2",
			UserID:      userId,
			CreatedDate: createdDate,
		},
		{
			ID:          "3",
			Content:     "task 3",
			UserID:      userId,
			CreatedDate: createdDate,
		},
	}
	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), userId))
	req.Form = url.Values{}
	req.Form.Set("created_date", createdDate)

	service.Store.(*mocks.IDatabase).On("RetrieveTasks", userId, createdDate).Return(tasks, nil)
	service.ListTasks(resp, req)

	result := resp.Result()
	require.Equal(t, http.StatusOK, result.StatusCode)

	responseData := &ResponseData{}
	err := json.NewDecoder(result.Body).Decode(responseData)
	require.NoError(t, err)

	var tasksResults []*models.Task
	err = mapstructure.Decode(responseData.Data, &tasksResults)
	require.NoError(t, err)
	require.Len(t, tasksResults, len(tasks))

	for i, task := range tasksResults {
		require.Equal(t, tasks[i].ID, task.ID)
		require.Equal(t, tasks[i].Content, task.Content)
		require.Equal(t, tasks[i].UserID, task.UserID)
		require.Equal(t, tasks[i].CreatedDate, task.CreatedDate)
	}
}
