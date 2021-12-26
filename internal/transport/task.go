package transport

import (
	"encoding/json"
	"net/http"

	"github.com/perfectbuii/togo/internal/domain"
	"github.com/perfectbuii/togo/internal/storages"
)

type TaskHandler interface {
	GetList(w http.ResponseWriter, r *http.Request)
	AddTask(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	taskDomain domain.TaskDomain
}

func (h *taskHandler) GetList(w http.ResponseWriter, r *http.Request) {
	createdDate, ok := r.URL.Query()["created_date"]
	if !ok || len(createdDate[0]) < 1 {
		http.Error(w, "Url Param 'created_date' is missing", http.StatusInternalServerError)
		return
	}
	tasks, err := h.taskDomain.GetList(r.Context(), createdDate[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWithJson(w, http.StatusOK, tasks)
}

func (h *taskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	req := &storages.Task{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	task, err := h.taskDomain.Create(r.Context(), req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWithJson(w, http.StatusOK, task)
}

func NewTaskHandler(taskDomain domain.TaskDomain) TaskHandler {
	return &taskHandler{
		taskDomain: taskDomain,
	}
}
