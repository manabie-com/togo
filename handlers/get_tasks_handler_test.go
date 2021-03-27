package handlers

import (
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/pkg/core"
	"github.com/manabie-com/togo/usecases"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTasksHandler_Success(t *testing.T) {
	req := createGetTasksRequest("")
	expectedUserId := int64(9999)
	expectedResp := `{"data":[{"id":0,"content":"","created_date":"0001-01-01T00:00:00Z"}],"status":"success"}`

	mockRepo := new(usecases.DBMock)
	mockRepo.On("GetTasks", req.Context(), &domains.TaskRequest{UserId: expectedUserId}).Return([]*domains.Task{{}}, nil)
	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(expectedUserId, nil)

	resp := getTasksResponse(req, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func TestGetTasksHandler_SuccessWithCreatedDateParams(t *testing.T) {
	req := createGetTasksRequest("2021-01-02")
	expectedUserId := int64(9999)
	expectedResp := `{"data":[{"id":0,"content":"","created_date":"0001-01-01T00:00:00Z"}],"status":"success"}`

	mockRepo := new(usecases.DBMock)
	cDate, _ := time.Parse("2006-01-02", "2021-01-02")
	mockRepo.
		On("GetTasks", req.Context(), &domains.TaskRequest{UserId: expectedUserId, CreatedDate: cDate}).
		Return([]*domains.Task{{}}, nil)
	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(expectedUserId, nil)

	resp := getTasksResponse(req, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func TestGetTasksHandler_SuccessWithEmptyResult(t *testing.T) {
	req := createGetTasksRequest("2021-01-02")
	expectedUserId := int64(9999)
	expectedResp := `{"data":[],"status":"success"}`

	mockRepo := new(usecases.DBMock)
	cDate, _ := time.Parse("2006-01-02", "2021-01-02")
	mockRepo.
		On("GetTasks", req.Context(), &domains.TaskRequest{UserId: expectedUserId, CreatedDate: cDate}).
		Return([]*domains.Task{}, nil)
	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(expectedUserId, nil)

	resp := getTasksResponse(req, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func TestGetTasksHandler_FailUnauthorized(t *testing.T) {
	req := createGetTasksRequest("")
	expectedResp := `{"message":"un-authorized","status":"error"}`

	mockRepo := new(usecases.DBMock)
	mockAuth := new(usecases.AuthMock)
	mockAuth.On("ValidateToken", req).Return(int64(0), core.AuthTokenIsInvalid)

	resp := getTasksResponse(req, mockRepo, mockAuth)
	defer resp.Body.Close()

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	res, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, expectedResp, string(res))
}

func getTasksResponse(req *http.Request, mockRepo domains.TaskRepository, mockAuth core.AppAuthenticator) *http.Response {
	uc := usecases.NewGetTasksUseCase(mockRepo)
	recorder := httptest.NewRecorder()
	handler := &GetTasksHandler{Uc: uc, Auth: mockAuth}
	handler.ServeHTTP(recorder, req)

	return recorder.Result()
}

func createGetTasksRequest(createdDate string) *http.Request {
	url := "localhost:5050/tasks"
	if createdDate != "" {
		url += "?created_date=" + createdDate
	}
	req := httptest.NewRequest("GET", url, nil)
	return req
}