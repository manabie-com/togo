package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/manabie-com/togo/model"
	intCtx "github.com/manabie-com/togo/pkg/ctx"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/go-chi/chi"

	"github.com/manabie-com/togo/pkg/httpx"

	"github.com/manabie-com/togo/internal/task/service"
)

type TaskHandler struct {
	taskService service.TaskService
}

func New(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	req := new(CreateTaskRequest)

	if err = json.Unmarshal(b, req); err != nil {
		httpx.WriteError(w, err)
		return
	}

	if err = req.Validate(); err != nil {
		httpx.WriteError(w, err)
		return
	}

	currentUser := intCtx.Get(r.Context(), intCtx.UserKey).(*model.User)
	err = h.taskService.CreateTask(r.Context(), &service.CreateTaskArgs{
		Content: req.Content,
		UserID:  currentUser.ID,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, nil)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpx.WriteError(w, errorx.ErrInvalidParameter(err))
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpx.WriteError(w, errorx.ErrInvalidParameter(err))
	}

	serviceTasks, err := h.taskService.GetTasks(r.Context(), &service.GetTasksArgs{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	res := convertServiceTasksToHandlerTasks(serviceTasks)
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, res)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "task_id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		httpx.WriteError(w, errorx.ErrInvalidParameter(err))
	}
	serviceTask, err := h.taskService.GetTask(r.Context(), &service.GetTaskArgs{
		ID: id,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	res := convertServiceTaskToHandlerTask(serviceTask)
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, res)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "task_id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		httpx.WriteError(w, errorx.ErrInvalidParameter(err))
	}

	err = h.taskService.DeleteTask(r.Context(), &service.DeleteTaskArgs{
		ID: id,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, nil)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	req := new(UpdateTaskRequest)
	if err = json.Unmarshal(b, req); err != nil {
		httpx.WriteError(w, err)
		return
	}

	taskID := chi.URLParam(r, "task_id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		httpx.WriteError(w, errorx.ErrInvalidParameter(err))
	}

	err = h.taskService.UpdateTask(r.Context(), &service.UpdateTaskArgs{
		TaskID:  id,
		Content: req.Content,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, nil)
}
