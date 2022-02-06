package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kozloz/togo"
	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/tasks"
)

type Task struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

// CreateTaskResult is a JSON object needed to create a task
type CreateTaskRequest struct {
	Name string `json:"name"`
}

// CreateTaskResult is a JSON object we'll return in our API
type CreateTaskResult struct {
	Error `json:"error"`
	*Task `json:"task,omitempty"`
}

// TaskHandler is a handler that can be attached to a router. This handler specifically manages the Task resource.
type TaskHandler struct {
	op *tasks.Operation
}

// Converts the error to our JSON Error type
func TaskToJSON(task *togo.Task) *Task {
	return &Task{
		ID:     task.ID,
		UserID: task.UserID,
		Name:   task.Name,
	}
}

// CreateTask is a handler that creates the task for the user
func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	// Parse request
	res := CreateTaskResult{}
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.ParseInt(userIDStr, 10, 0)
	if err != nil {
		resErr := CustomErrorToJSON(err)
		res.Error = resErr
		resBody, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resBody)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resErr := CustomErrorToJSON(err)
		res.Error = resErr
		resBody, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resBody)
		return
	}

	log.Println(string(body))
	createTaskReq := CreateTaskRequest{}
	err = json.Unmarshal(body, &createTaskReq)
	if err != nil {
		resErr := CustomErrorToJSON(err)
		res.Error = resErr
		resBody, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resBody)
		return
	}

	// Request validation
	if createTaskReq.Name == "" {
		log.Printf("Error: Empty task name provided")
		resErr := CustomErrorToJSON(errors.InvalidTaskName)
		res.Error = resErr
		w.WriteHeader(http.StatusBadRequest)
		resBody, _ := json.Marshal(res)
		w.Write(resBody)
		return
	}

	// Create the task via operation class
	task, err := t.op.Create(userID, createTaskReq.Name)
	if err != nil {
		resErr := CustomErrorToJSON(err)
		res.Error = resErr
		resBody, _ := json.Marshal(res)
		w.Write(resBody)
		return
	}

	res.Task = TaskToJSON(task)
	resErr := CustomErrorToJSON(errors.Success)
	res.Error = resErr
	resBody, _ := json.Marshal(res)
	w.Write(resBody)
}
