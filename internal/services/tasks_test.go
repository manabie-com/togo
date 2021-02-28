package services

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	taskmodel "github.com/manabie-com/togo/internal/storages/task/model"
	_ "github.com/manabie-com/togo/internal/storages/user/model"
	cmsqlmock "github.com/manabie-com/togo/pkg/common/cmsql/mock"
	"github.com/manabie-com/togo/pkg/common/crypto"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
)

type AddTaskResponse struct {
	Data struct {
		taskmodel.Task
	} `json:"data"`
}

type ListTaskResponse struct {
	Data []*taskmodel.Task `json:"data"`
}

func TestToDoService_AddTask(t *testing.T) {
	mock := integrationTest.mock

	passwordHashed, _ := crypto.HashPassword("example")
	mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
		WithArgs("000001").
		WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
			AddRow("000001", passwordHashed, 1))

	t.Run("TestToDoService_AddTask_Login", func(t *testing.T) {
		loginBody, err := json.Marshal(map[string]string{
			"user_id":  "000001",
			"password": "example",
		})
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyStr, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var loginResp LoginResponse

		err = json.Unmarshal(bodyStr, &loginResp)
		assert.Nil(t, err)

		t.Run("TestToDoService_AddTask_AddTask", func(t *testing.T) {
			mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
				WithArgs("000001").
				WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
					AddRow("000001", passwordHashed, 1))

			mock.ExpectQuery("SELECT count(id) as count FROM tasks WHERE user_id = $1 AND created_date = $2").
				WithArgs("000001", cmsqlmock.AnyString{}).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			mock.ExpectExec("INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)").
				WithArgs(cmsqlmock.AnyString{}, "content task 1", "000001", cmsqlmock.AnyString{}).
				WillReturnResult(sqlmock.NewResult(1, 1))

			addTaskBody, err := json.Marshal(map[string]string{
				"content": "content task 1",
			})
			assert.Nil(t, err)

			addTaskReq := httptest.NewRequest("POST", "/tasks", bytes.NewReader(addTaskBody))
			addTaskReq.Header.Set("authorization", loginResp.Data)
			w := httptest.NewRecorder()
			integrationTest.hander.ServeHTTP(w, addTaskReq)

			addTaskResponse := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			addTaskBodyStr, err := ioutil.ReadAll(addTaskResponse.Body)
			assert.Nil(t, err)

			var addTaskResp AddTaskResponse

			assert.Nil(t, json.Unmarshal(addTaskBodyStr, &addTaskResp))

			assert.NotNil(t, addTaskResp.Data)
			assert.Equal(t, "content task 1", addTaskResp.Data.Content)
		})

		t.Run("TestToDoService_AddTask_AddTask_TodosLimitIsReached", func(t *testing.T) {
			addTaskBody, err := json.Marshal(map[string]string{
				"content": "content task 1",
			})
			assert.Nil(t, err)

			addTaskReq := httptest.NewRequest("POST", "/tasks", bytes.NewReader(addTaskBody))
			addTaskReq.Header.Set("authorization", loginResp.Data)
			w := httptest.NewRecorder()
			integrationTest.hander.ServeHTTP(w, addTaskReq)

			addTaskResponse := w.Result()
			assert.Equal(t, http.StatusBadRequest, addTaskResponse.StatusCode)

			addTaskBodyStr, err := ioutil.ReadAll(addTaskResponse.Body)
			assert.Nil(t, err)

			var errResponse ErrorResponse

			assert.Nil(t, json.Unmarshal(addTaskBodyStr, &errResponse))

			assert.Equal(t, "the number of todos daily limit is reached", errResponse.Error)
		})
	})
}

func TestToDoService_ListTasks(t *testing.T) {
	mock := integrationTest.mock

	passwordHashed, _ := crypto.HashPassword("example")
	mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
		WithArgs("000001").
		WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
			AddRow("000001", passwordHashed, 10))

	t.Run("TestToDoService_ListTasks_Login", func(t *testing.T) {
		loginBody, err := json.Marshal(map[string]string{
			"user_id":  "000001",
			"password": "example",
		})
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		bodyStr, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var loginResp LoginResponse

		err = json.Unmarshal(bodyStr, &loginResp)
		assert.Nil(t, err)

		t.Run("TestToDoService_ListTasks_ListTasks", func(t *testing.T) {
			mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2").
				WithArgs("000001", "2020-02-21").
				WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
					AddRow("task_00001", "content task 1", "000001", "2020-02-21").
					AddRow("task_00002", "content task 2", "000001", "2020-02-21"))

			listTasksReq := httptest.NewRequest("GET", "/tasks", nil)

			q := listTasksReq.URL.Query()
			q.Set("created_date", "2020-02-21")
			listTasksReq.URL.RawQuery = q.Encode()

			listTasksReq.Header.Set("authorization", loginResp.Data)
			w := httptest.NewRecorder()
			integrationTest.hander.ServeHTTP(w, listTasksReq)

			assert.Nil(t, err)

			listTasksResponse := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var listTasksBodyResp ListTaskResponse

			listTasksBodyStr, err := ioutil.ReadAll(listTasksResponse.Body)
			assert.Nil(t, err)
			assert.Nil(t, json.Unmarshal(listTasksBodyStr, &listTasksBodyResp))

			assert.NotNil(t, listTasksBodyResp.Data)
			assert.Equal(t, len(listTasksBodyResp.Data), 2)
			assert.Equal(t, listTasksBodyResp.Data[0].CreatedDate, "2020-02-21")
			assert.Equal(t, listTasksBodyResp.Data[1].CreatedDate, "2020-02-21")
			assert.Equal(t, listTasksBodyResp.Data[0].UserID, "000001")
		})

		t.Run("TestToDoService_ListTasks_ListTasks_NoResults", func(t *testing.T) {
			mock.ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2").
				WithArgs("000001", "2020-02-22").
				WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}))

			listTasksReq := httptest.NewRequest("GET", "/tasks", nil)

			q := listTasksReq.URL.Query()
			q.Set("created_date", "2020-02-22")
			listTasksReq.URL.RawQuery = q.Encode()

			listTasksReq.Header.Set("authorization", loginResp.Data)
			w := httptest.NewRecorder()
			integrationTest.hander.ServeHTTP(w, listTasksReq)

			assert.Nil(t, err)

			listTasksResponse := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var listTasksBodyResp ListTaskResponse

			listTasksBodyStr, err := ioutil.ReadAll(listTasksResponse.Body)
			assert.Nil(t, err)
			assert.Nil(t, json.Unmarshal(listTasksBodyStr, &listTasksBodyResp))

			assert.Nil(t, listTasksBodyResp.Data)
		})
	})
}
