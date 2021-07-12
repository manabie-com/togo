package handler

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/services/task"
)

type (
	taskHandler struct {
		taskSvc task.Service
	}

	CreateTaskRequest struct {
		Content string `json:"content"`
	}
)

func NewTaskHandler(taskSvc task.Service) *taskHandler {
	return &taskHandler{
		taskSvc: taskSvc,
	}
}

func (s *taskHandler) List(resp http.ResponseWriter, req *http.Request) {
	createDate := req.FormValue("created_date")

	tasks, err := s.taskSvc.List(req.Context(), createDate)
	if err != nil {
		respondWithError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(resp, http.StatusOK, tasks)
}

func (s *taskHandler) Create(resp http.ResponseWriter, req *http.Request) {
	createTaskRequest, err := getCreateTaskRequest(req)
	if err != nil {
		respondWithError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	t, err := s.taskSvc.Create(req.Context(), createTaskRequest.Content)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if err == task.ErrReachedOutTaskTodoPerDay {
			httpStatus = http.StatusTooManyRequests
		}
		respondWithError(resp, httpStatus, err.Error())
		return
	}
	respondWithJSON(resp, http.StatusOK, t)
}

func getCreateTaskRequest(r *http.Request) (*CreateTaskRequest, error) {
	var createTaskRequest CreateTaskRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&createTaskRequest); err != nil {
		return nil, err
	}
	return &createTaskRequest, nil
}
