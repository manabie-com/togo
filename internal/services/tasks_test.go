package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	mock_sqllite "github.com/manabie-com/togo/internal/storages/sqlite/db_mock"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

type LoginJSONResponse struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type ListTasksJSONResponse struct {
	Data  []*storages.Task `json:"data"`
	Error string           `json:"error"`
}

type AddTasksJSONResponse struct {
	Data  *storages.Task `json:"data"`
	Error string         `json:"error"`
}

var (
	JWTKeyDummy     string = "1234567890abcdxyz"
	testUser               = storages.User{ID: "userid", Password: "password"}
	testCreatedDate        = "2020-11-22"
	testTaskData           = []*storages.Task{
		{ID: "123", Content: "content1", UserID: testUser.ID, CreatedDate: testCreatedDate},
		{ID: "456", Content: "content2", UserID: testUser.ID, CreatedDate: testCreatedDate},
	}
)

func TestGetAuthToken_Success(t *testing.T) {
	resp := mockGetAuthToken(t, testUser, true)
	defer resp.Body.Close()

	require := require.New(t)
	require.Equal(http.StatusOK, resp.StatusCode)
	var jsonResp LoginJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	require.NotEmpty(jsonResp.Data)
	require.Empty(jsonResp.Error)
}

func TestGetAuthToken_InvalidCredentials(t *testing.T) {
	resp := mockGetAuthToken(t, testUser, false)
	defer resp.Body.Close()

	require := require.New(t)
	require.Equal(http.StatusUnauthorized, resp.StatusCode)
	var jsonResp LoginJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	require.Empty(jsonResp.Data)
	require.NotEmpty(jsonResp.Error)
}

func TestListTasks_Success(t *testing.T) {
	testCases := [][]*storages.Task{
		testTaskData, // Case user has some tasks available on such date
		{},           // Case user has no tasks avaiable on such date
	}

	for _, taskData := range testCases {
		resp := mockListTasks(t, testUser.ID, testCreatedDate, taskData, nil)
		defer resp.Body.Close()

		require := require.New(t)
		require.Equal(http.StatusOK, resp.StatusCode)

		var jsonResp ListTasksJSONResponse
		require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
		require.Len(jsonResp.Data, len(taskData))
		require.Empty(jsonResp.Error)
		for _, v := range jsonResp.Data {
			require.NotEmpty(v.ID)
			require.NotEmpty(v.Content)
			require.Equal(testUser.ID, v.UserID)
			require.Equal(testCreatedDate, v.CreatedDate)
		}
	}
}

func TestListTasks_InternalError(t *testing.T) {
	resp := mockListTasks(t, testUser.ID, testCreatedDate, testTaskData, errors.New("dummy error"))
	defer resp.Body.Close()

	require := require.New(t)
	require.Equal(http.StatusInternalServerError, resp.StatusCode)

	var jsonResp ListTasksJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	require.Empty(jsonResp.Data)
	require.NotEmpty(jsonResp.Error)
}

func TestAddTask_Success(t *testing.T) {
	taskData := testTaskData[0]
	resp := mockAddTasks(t, taskData, 1, nil)
	defer resp.Body.Close()

	require := require.New(t)
	require.Equal(http.StatusOK, resp.StatusCode)

	var jsonResp AddTasksJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	require.Empty(jsonResp.Error)
	require.Equal(taskData.UserID, jsonResp.Data.UserID)
	require.Equal(taskData.Content, jsonResp.Data.Content)
}

func TestAddTask_DatabaseError(t *testing.T) {
	taskData := testTaskData[0]
	resp := mockAddTasks(t, taskData, 0, errors.New("dummy database zeroerror"))
	defer resp.Body.Close()

	require := require.New(t)
	require.Equal(http.StatusInternalServerError, resp.StatusCode)

	var jsonResp AddTasksJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	require.NotEmpty(jsonResp.Error)
	require.Empty(jsonResp.Data)
}

func TestAddTask_UserQuotaExceeded(t *testing.T) {
	taskData := testTaskData[0]
	resp := mockAddTasks(t, taskData, 0, nil)
	defer resp.Body.Close()

	require := require.New(t)
	require.Equal(http.StatusTooManyRequests, resp.StatusCode)

	var jsonResp AddTasksJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	require.NotEmpty(jsonResp.Error)
	require.Empty(jsonResp.Data)
}

func ns(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func newLoginRequest(userID, password string) *http.Request {
	req := httptest.NewRequest("GET", "localhost:5050/login", nil)
	q := req.URL.Query()
	q.Add("user_id", userID)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()
	return req
}

func newListTasksRequest(userID, createdDate string) *http.Request {
	req := httptest.NewRequest("GET", "localhost:5050/tasks", nil)
	q := req.URL.Query()
	q.Add("user_id", userID)
	q.Add("created_date", createdDate)
	req.URL.RawQuery = q.Encode()
	return req
}

func newAddTaskRequest(content string) *http.Request {
	payload, err := json.Marshal(struct {
		Content string `json:"content"`
	}{
		Content: content,
	})
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest("POST", "localhost:5050/tasks", bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	return req
}

func newContextWithUserID(userID string) context.Context {
	return context.WithValue(context.Background(), userAuthKey(0), userID)
}

func mockGetAuthToken(t *testing.T, u storages.User, validationResult bool) *http.Response {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_sqllite.NewMockDatabaser(mockCtrl)
	store.EXPECT().
		ValidateUser(context.Background(), ns(u.ID), ns(u.Password)).
		Return(validationResult)

	service := ToDoService{
		JWTKey: JWTKeyDummy,
		Store:  store,
	}
	req := newLoginRequest(u.ID, u.Password)
	w := httptest.NewRecorder()
	service.getAuthToken(w, req)
	return w.Result()
}

func mockListTasks(t *testing.T, uid, cdate string, taskData []*storages.Task, err error) *http.Response {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := newContextWithUserID(uid)
	store := mock_sqllite.NewMockDatabaser(mockCtrl)
	store.EXPECT().
		RetrieveTasks(ctx, ns(uid), ns(cdate)).
		Return(taskData, err)

	service := ToDoService{
		JWTKey: JWTKeyDummy,
		Store:  store,
	}
	req := newListTasksRequest(uid, cdate).WithContext(ctx)
	w := httptest.NewRecorder()
	service.listTasks(w, req)
	return w.Result()
}

// Only UserID and Content fields from taskData are used.
func mockAddTasks(t *testing.T, taskData *storages.Task, res int64, err error) *http.Response {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := newContextWithUserID(taskData.UserID)
	store := mock_sqllite.NewMockDatabaser(mockCtrl)
	store.EXPECT().
		AddTask(ctx, gomock.Any()).
		Return(res, err)
	service := ToDoService{
		JWTKey: JWTKeyDummy,
		Store:  store,
	}
	req := newAddTaskRequest(taskData.Content).WithContext(ctx)
	w := httptest.NewRecorder()
	service.addTask(w, req)
	return w.Result()
}
