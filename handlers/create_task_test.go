package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/pkg/core"
	"github.com/manabie-com/togo/usecases"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTaskHandler_Success(t *testing.T) {
	expectedUserId := int64(9999)
	expectedContent := "AAAAA"
	req := createTaskRequest(&usecases.TaskInput{Content: expectedContent})
	expectedResp := `{"data":{"id":0,"content":"AAAAA","created_date":"0001-01-01T00:00:00Z"},"status":"success"}`

	mockRepo := new(usecases.DBMock)
	mockRepo.On("GetUserById", req.Context(), expectedUserId).Return(&domains.User{MaxTodo: 5, Id: expectedUserId}, nil)
	mockRepo.On("GetCountCreatedTaskTodayByUser", req.Context(), expectedUserId).Return(int64(4), nil)
	mockRepo.On("CreateTask", req.Context(), &domains.Task{Content: expectedContent, UserId: expectedUserId}).
		Return(&domains.Task{Content: expectedContent}, nil)

	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(expectedUserId, nil)

	resp := createTaskResponse(req, mockRepo, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func TestCreateTaskHandler_FailReachedLimit(t *testing.T) {
	expectedUserId := int64(9999)
	expectedContent := "AAAAA"
	req := createTaskRequest(&usecases.TaskInput{Content: expectedContent})
	expectedResp := `{"message":"reached limit create tasks per day","status":"error"}`

	mockRepo := new(usecases.DBMock)
	mockRepo.On("GetUserById", req.Context(), expectedUserId).Return(&domains.User{MaxTodo: 5, Id: expectedUserId}, nil)
	mockRepo.On("GetCountCreatedTaskTodayByUser", req.Context(), expectedUserId).Return(int64(5), nil)

	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(expectedUserId, nil)

	resp := createTaskResponse(req, mockRepo, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func TestCreateTaskHandler_FailUnauthorized(t *testing.T) {
	expectedContent := "AAAAA"
	req := createTaskRequest(&usecases.TaskInput{Content: expectedContent})
	expectedResp := `{"message":"un-authorized","status":"error"}`

	mockRepo := new(usecases.DBMock)
	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(int64(0), core.AuthTokenIsInvalid)

	resp := createTaskResponse(req, mockRepo, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func createTaskResponse(req *http.Request, mockRepo domains.TaskRepository, userRepo domains.UserRepository, mockAuth core.AppAuthenticator) *http.Response {
	uc := usecases.NewCreateTaskUseCase(mockRepo, userRepo)
	recorder := httptest.NewRecorder()
	handler := &CreateTaskHandler{Uc: uc, Auth: mockAuth}
	handler.ServeHTTP(recorder, req)
	return recorder.Result()
}

func createTaskRequest(input *usecases.TaskInput) *http.Request {
	url := "localhost:5050/tasks"
	b, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(b))
	return req
}
