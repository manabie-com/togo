package handler

import (
	"encoding/json"
	"net/http"
	"time"

	mm "github.com/manabie-com/togo/internal/pkg/middleware"
	d "github.com/manabie-com/togo/internal/todo/domain"
	s "github.com/manabie-com/togo/internal/todo/service"
)

type TaskHandler struct {
	AppHandler
	taskService *s.TaskService
}

func NewTaskHandler(taskService *s.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	rLog := mm.GetLogEntry(r)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	currentUserID, err := h.getUserIDFromCtx(r)
	if err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	taskCreateParam := d.TaskCreateParam{}
	if err := decoder.Decode(&taskCreateParam); err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	task, err := h.taskService.CreateTaskForUser(currentUserID, taskCreateParam)
	if err == d.ErrTaskLimitReached {
		h.responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	data := map[string]*d.Task{
		"data": task,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Error parsing response")
	}
}

func (h *TaskHandler) ListTask(w http.ResponseWriter, r *http.Request) {
	rLog := mm.GetLogEntry(r)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	currentUserID, err := h.getUserIDFromCtx(r)
	if err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	dateParam := r.URL.Query().Get("created_date")
	if _, err := time.Parse("2006-01-02", dateParam); err != nil {
		dateParam = time.Now().Format("2006-01-02")
	}

	tasks, err := h.taskService.ListTaskForUser(currentUserID, dateParam)
	if err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	data := map[string][]*d.Task{
		"data": tasks,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Error parsing response")
	}
}
