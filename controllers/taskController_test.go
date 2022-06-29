package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/context"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/huynhhuuloc129/todo/util"
	"github.com/stretchr/testify/assert"
)

func CreateMockingDB() (sqlmock.Sqlmock, *BaseHandler) {
	db, mock := models.NewMock()
	dbConn := models.NewdbConn(db)
	h := NewBaseHandler(dbConn)
	return mock, h
}

func TestResponseAllTask(t *testing.T) {
	mock, h := CreateMockingDB()
	userId := util.RandomInt(0, 100)

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < 10; i++ {
		task := models.RandomTask()
		rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, userId)
	}

	//query
	mock.ExpectQuery(regexp.QuoteMeta(models.QueryAllTaskText)).WithArgs(userId).WillReturnRows(rows)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/tasks", nil)
	context.Set(req, "userid", userId)
	h.ResponseAllTask(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var tasks []models.Task
	err = json.Unmarshal(respBody, &tasks)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, respBody)
	assert.Len(t, tasks, 10)
}

func TestResponseOneTask(t *testing.T) {
	mock, h := CreateMockingDB()
	task := models.RandomTask()

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)

	//query
	mock.ExpectQuery(regexp.QuoteMeta(models.FindTaskByIDText)).WithArgs(task.Id, task.UserId).WillReturnRows(rows)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/tasks/"+fmt.Sprintf("%v", task.UserId), nil)
	context.Set(req, "id", task.Id)
	context.Set(req, "userid", task.UserId)
	h.ResponseOneTask(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var taskfromdb models.Task
	err = json.Unmarshal(respBody, &taskfromdb)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, taskfromdb)
	assert.Equal(t, taskfromdb.Content, task.Content)
	assert.Equal(t, taskfromdb.Status, task.Status)
	assert.Equal(t, taskfromdb.UserId, task.UserId)
}

func TestCreateTask(t *testing.T) {
	mock, h := CreateMockingDB()
	task := models.RandomTask()
	taskJSON, err := json.Marshal(task)
	if err != nil {
		t.Errorf("Can't marshal task, err: " + err.Error())
	}

	//exec
	mock.ExpectExec(regexp.QuoteMeta(models.InsertTaskText)).WithArgs(task.Content, task.Status, task.Time, task.TimeDone, task.UserId).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder() // set custom writer and response
	req := httptest.NewRequest("POST", "localhost:8000/users", bytes.NewReader(taskJSON))
	context.Set(req, "userid", task.UserId)
	h.CreateTask(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var taskfromdb models.NewTask
	err = json.Unmarshal(respBody, &taskfromdb)
	if err != nil {
		fmt.Println(taskfromdb, string(respBody))
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, taskfromdb)
	assert.Equal(t, taskfromdb.Content, task.Content)
	assert.Equal(t, taskfromdb.Status, task.Status)
	assert.Equal(t, taskfromdb.UserId, task.UserId)
}

func TestDeleteFromTask(t *testing.T) {
	mock, h := CreateMockingDB()
	task := models.RandomTask()

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)

	mock.ExpectQuery(regexp.QuoteMeta(models.FindTaskByIDText)).WithArgs(task.Id, task.UserId).WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(models.DeleteTaskText)).WithArgs(task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("localhost:8000/%v", task.Id), nil)
	context.Set(req, "userid", task.UserId)
	context.Set(req, "id", task.Id)
	h.DeleteFromTask(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	respBodyString := string(respBody)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, respBodyString, "message: delete success")
}

func TestUpdateToTask(t *testing.T) {
	mock, h := CreateMockingDB()

	task := models.RandomTask()
	newTask := models.RandomNewTask()
	newTaskJSON, err := json.Marshal(newTask)
	if err != nil {
		t.Errorf("Can't marshal task, err: " + err.Error())
	}

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)

	mock.ExpectQuery(regexp.QuoteMeta(models.FindTaskByIDText)).WithArgs(task.Id, task.UserId).WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(models.UpdateTaskText)).WithArgs(newTask.Content, newTask.Status, newTask.TimeDone, task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", fmt.Sprintf("localhost:8000/%v", task.Id), bytes.NewReader(newTaskJSON))
	context.Set(req, "id", task.Id)
	context.Set(req, "userid", task.UserId)
	h.UpdateToTask(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response, err: " + err.Error())
	}

	var taskfromdb models.NewTask
	err = json.Unmarshal(respBody, &taskfromdb)
	if err != nil {
		fmt.Println(taskfromdb, string(respBody))
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, taskfromdb.Content, newTask.Content)
	assert.Equal(t, taskfromdb.Status, newTask.Status)
	assert.Equal(t, taskfromdb.UserId, newTask.UserId)
}
