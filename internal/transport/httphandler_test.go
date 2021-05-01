package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
)

var (
	validateUserMock func(ctx context.Context, userID, password string) bool
	listTasksMock    func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error)
	addTaskMock      func(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error)
)

type serviceMock struct{}

func (s *serviceMock) ValidateUser(ctx context.Context, userID, password string) bool {
	return validateUserMock(ctx, userID, password)
}
func (s *serviceMock) ListTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	return listTasksMock(ctx, userID, createdDate)
}
func (s *serviceMock) AddTask(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error) {
	return addTaskMock(ctx, userID, reqBody)
}

type tokenResponse struct {
	Token string `json:"data"`
}

type getTasksResponse struct {
	Tasks []*storages.Task `json:"data"`
}

type addTasksRequest struct {
	Content string `json:"content"`
}

type addTasksResponse struct {
	Task *storages.Task `json:"data"`
}

func TestHttpHandler_Login(t *testing.T) {
	jwtToken := "wqGyEBBfPK9w3Lxw"
	httpHandler := NewHttpHandler(jwtToken, &serviceMock{})

	validateUserMock = func(ctx context.Context, userID, password string) bool {
		return true
	}
	token, err := login(httpHandler, "someUser", "somePassword")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	validateUserMock = func(ctx context.Context, userID, password string) bool {
		return false
	}

	token, err = login(httpHandler, "someUser", "somePassword")
	assert.Nil(t, err)
	assert.Empty(t, token)
}

func TestHttpHandler_GetTasks(t *testing.T) {
	jwtToken := "wqGyEBBfPK9w3Lxw"
	httpHandler := NewHttpHandler(jwtToken, &serviceMock{})

	statusCode, tasks, err := getTasks(httpHandler, "")
	assert.EqualValues(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, tasks)
	assert.NotNil(t, err)

	validateUserMock = func(ctx context.Context, userID, password string) bool {
		return true
	}
	token, err := login(httpHandler, "someUser", "somePassword")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	listTasksMock = func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
		return []*storages.Task{
			{
				ID:          "1",
				Content:     "c1",
				UserID:      "someUser",
				CreatedDate: "2021-04-30",
			},
			{
				ID:          "2",
				Content:     "c2",
				UserID:      "someUser",
				CreatedDate: "2021-04-30",
			},
		}, nil

	}
	statusCode, tasks, err = getTasks(httpHandler, token)
	assert.EqualValues(t, http.StatusOK, statusCode)
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
	assert.EqualValues(t, 2, len(tasks))
	assert.EqualValues(t, "1", tasks[0].ID)
	assert.EqualValues(t, "c1", tasks[0].Content)
	assert.EqualValues(t, "someUser", tasks[0].UserID)
	assert.EqualValues(t, "2021-04-30", tasks[0].CreatedDate)
	assert.EqualValues(t, "2", tasks[1].ID)
	assert.EqualValues(t, "c2", tasks[1].Content)
	assert.EqualValues(t, "someUser", tasks[1].UserID)
	assert.EqualValues(t, "2021-04-30", tasks[1].CreatedDate)

	listTasksMock = func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
		return nil, errors.New("some error")
	}
	statusCode, tasks, err = getTasks(httpHandler, token)
	assert.EqualValues(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, tasks)
}

func TestHttpHandler_AddTask(t *testing.T) {
	jwtToken := "wqGyEBBfPK9w3Lxw"
	httpHandler := NewHttpHandler(jwtToken, &serviceMock{})

	addTaskReq := addTasksRequest{Content: "example"}
	req, err := json.Marshal(addTaskReq)
	assert.Nil(t, err)

	statusCode, task, err := addTask(httpHandler, "", req)
	assert.EqualValues(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, task)
	assert.NotNil(t, err)

	validateUserMock = func(ctx context.Context, userID, password string) bool {
		return true
	}
	token, err := login(httpHandler, "someUser", "somePassword")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	addTaskMock = func(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error) {
		return &storages.Task{
			ID:          "1",
			Content:     "example",
			UserID:      "someUser",
			CreatedDate: "2021-04-30",
		}, nil

	}
	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusOK, statusCode)
	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.EqualValues(t, "someUser", task.UserID)
	assert.EqualValues(t, "example", task.Content)

	addTaskMock = func(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error) {
		return nil, errors.New("some error")

	}
	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, task)

	addTaskMock = func(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error) {
		return nil, services.ErrUserReachDailyRequestLimit

	}
	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusForbidden, statusCode)
	assert.Nil(t, task)
}

func login(httpHandler http.Handler, userID, password string) (string, error) {
	request := httptest.NewRequest(http.MethodGet, "/login?user_id="+userID+"&password="+password, nil)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	tokenResp := tokenResponse{}
	err := json.NewDecoder(response.Body).Decode(&tokenResp)
	if err != nil {
		return "", err
	}
	return tokenResp.Token, nil
}

func getTasks(httpHandler http.Handler, token string) (int, []*storages.Task, error) {
	request := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	getTasksResp := getTasksResponse{}
	err := json.NewDecoder(response.Body).Decode(&getTasksResp)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return response.Result().StatusCode, getTasksResp.Tasks, nil
}

func addTask(httpHandler http.Handler, token string, req []byte) (int, *storages.Task, error) {
	request := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(req))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	addTasksResp := addTasksResponse{}
	err := json.NewDecoder(response.Body).Decode(&addTasksResp)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return response.Result().StatusCode, addTasksResp.Task, nil
}
