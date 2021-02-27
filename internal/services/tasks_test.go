package services

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var testTasks           = []*storages.Task{
	{Id: 1, Content: "content 1", UsrId: 1, CreateAt: time.Now()},
	{Id: 2, Content: "content 2", UsrId: 1, CreateAt: time.Now()},
}

func TestListTask(t *testing.T) {
	testCases := [][]*storages.Task{
		testTasks, // Case user has some tasks available on such date
		{},           // Case user has no tasks avaiable on such date
	}

	for _, taskData := range testCases {
		resp := mockListTasks(t, 1, "2006-01-02", taskData, nil)

		requireTest := require.New(t)
		requireTest.Equal(http.StatusOK, resp.StatusCode)

		expectedRespData := &ApiDataResp{Data: taskData}
		assertDataResp(t, expectedRespData, resp)
	}
}

func TestListTaskErr(t *testing.T) {
	resp := mockListTasks(t, 1, "2006-01-02", testTasks, errInternal)
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusInternalServerError, resp.StatusCode)

	apiErrResp := &ApiErrResp{Error: errInternal.Error()}
	assertErrResp(t, apiErrResp, resp)
}

func TestAddTaskSuccess(t *testing.T) {
	testTask := &storages.Task{Content: "test content", UsrId: 1}
	resp := mockAddTasks(t, testTask, 1, nil)
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusOK, resp.StatusCode)

	expectedRespData := &ApiDataResp{Data: testTask}
	assertDataResp(t, expectedRespData, resp)
}

func TestAddTaskInternalErr(t *testing.T) {
	testTask := &storages.Task{Content: "test content", UsrId: 1}
	resp := mockAddTasks(t, testTask, 1, errInternal)
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusInternalServerError, resp.StatusCode)

	apiErrResp := &ApiErrResp{Error: errInternal.Error()}
	assertErrResp(t, apiErrResp, resp)
}

func TestAddTaskQuotaExceed(t *testing.T) {
	testTask := &storages.Task{Content: "test content", UsrId: 1}
	resp := mockAddTasks(t, testTask, 1, postgres.ErrUserMaxTodoReached)
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusTooManyRequests, resp.StatusCode)

	apiErrResp := &ApiErrResp{Error: postgres.ErrUserMaxTodoReached.Error()}
	assertErrResp(t, apiErrResp, resp)
}

func newListTasksRequest(userID int, createdDate string) *http.Request {
	req := httptest.NewRequest("GET", "localhost:5050/tasks", nil)
	q := req.URL.Query()
	q.Add("created_date", createdDate)
	req.URL.RawQuery = q.Encode()
	return req
}

func mockListTasks(t *testing.T, usrId int, createDate string, taskData []*storages.Task, taskErr error) *http.Response {
	ctx := context.WithValue(context.Background(), authSubKey, usrId)
	req := newListTasksRequest(usrId, createDate).WithContext(ctx)

	createdAt, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		t.Fail()
		return nil
	}

	db := new(postgres.DatabaseMock)
	db.On("GetTasks", req.Context(), usrId, createdAt).Return(taskData, taskErr)

	s := NewToDoService(testJWTKey, ":6000", db)
	w := httptest.NewRecorder()

	s.listTasksHandler(w, req)
	db.AssertExpectations(t)

	return w.Result()
}

func mockAddTasks(t *testing.T, taskData *storages.Task, usrId int, err error) *http.Response {
	ctx := context.WithValue(context.Background(), authSubKey, usrId)
	req := newAddTaskRequest(t, taskData.Content).WithContext(ctx)

	db := new(postgres.DatabaseMock)
	db.On("InsertTask", req.Context(), taskData).Return(err)

	s := NewToDoService(testJWTKey, ":6000", db)
	w := httptest.NewRecorder()

	s.addTaskHandler(w, req)
	db.AssertExpectations(t)

	return w.Result()
}

func newAddTaskRequest(t *testing.T, content string) *http.Request {
	payload, err := json.Marshal(struct {
		Content string `json:"content"`
	}{
		Content: content,
	})
	if err != nil {
		t.Fail()
	}
	req := httptest.NewRequest("POST", "localhost:5050/tasks", bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	return req
}
