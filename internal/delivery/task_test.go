package delivery

import (
	"bytes"
	"encoding/json"
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/pkg/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var taskToCreate = &e.Task{
	Name:        "task 1",
	Description: "description 1",
}

var task = &e.Task{
	ID:          1,
	Name:        "task 1",
	Description: "description 1",
	CreatedAt:   "2020-01-01",
	Completed:   false,
	UserID:      1,
}
var task2 = &e.Task{
	ID:          2,
	Name:        "task 2",
	Description: "description 2",
	CreatedAt:   "2020-01-01",
	Completed:   false,
	UserID:      1,
}

var tasks = &[]e.Task{
	*task,
	*task2,
}

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	taskUsecase.EXPECT().CreateTask(taskToCreate, user.Username).
		Return(task.ID, 1, nil).AnyTimes()

	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/task", strings.NewReader(`
	{
		"name": "task 1",
		"description": "description 1"
	}`))
	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	taskDelivery.CreateTask(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusCreated, w.Code)
	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	var element map[string]interface{}
	json.Unmarshal([]byte(body), &element)

	assert.Equal(t, float64(task.ID), element["taskID"])
}

func TestGetAllTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	taskUsecase.EXPECT().GetAllTask().
		Return(tasks, nil).AnyTimes()

	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/task", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	taskDelivery.GetAllTask(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.Equal(t, `[{"id":1,"name":"task 1","description":"description 1","created_at":"2020-01-01","completed":false,"user_id":1},{"id":2,"name":"task 2","description":"description 2","created_at":"2020-01-01","completed":false,"user_id":1}]`, body)
}

func TestGetAllTaskOfUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	taskUsecase.EXPECT().GetTasksByUsername(user.Username).
		Return(tasks, nil).AnyTimes()

	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/task", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	taskDelivery.GetAllTaskOfUser(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.Equal(t, `[{"id":1,"name":"task 1","description":"description 1","created_at":"2020-01-01","completed":false,"user_id":1},{"id":2,"name":"task 2","description":"description 2","created_at":"2020-01-01","completed":false,"user_id":1}]`, body)
}

func TestGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	taskUsecase.EXPECT().GetTaskByID(task.ID, user.Username).
		Return(task, nil).AnyTimes()

	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/task/1", nil)

	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	taskDelivery.GetTaskByID(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.Equal(t, `{"id":1,"name":"task 1","description":"description 1","created_at":"2020-01-01","completed":false,"user_id":1}`, body)
}

func TestCompleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	taskUsecase.EXPECT().CompleteTask(task.ID, user.Username).
		Return(nil).AnyTimes()

	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/task/1/complete", nil)

	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	taskDelivery.CompleteTask(w, r)

	// res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	taskUsecase.EXPECT().DeleteTask(task.ID, user.Username).
		Return(nil).AnyTimes()

	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/task/1", nil)

	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	taskDelivery.DeleteTask(w, r)

	// res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)
}
