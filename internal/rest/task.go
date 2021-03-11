package rest

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/services"
)

// I refer this package as a transport layer which handles in coming request
// at transport level to achieve the reusability of service layer

// TaskHandler like the controller for the routes that related to task domain
// It also have task service as a dependency
type TaskHandler struct {
	taskSvc services.TaskService
	authSvc services.AuthService
}

// NewTaskHandler ...
func NewTaskHandler(authSvc services.AuthService, taskSvc services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskSvc: taskSvc,
		authSvc: authSvc,
	}
}

// Register ...
func (th *TaskHandler) Register(mux *http.ServeMux) {
	mux.Handle("/task", MethodMiddleware(AuthMiddleware(LoggingMiddleWare(http.HandlerFunc(th.AddTask)), th.authSvc), http.MethodPost))
	mux.Handle("/tasks", MethodMiddleware(AuthMiddleware(LoggingMiddleWare(http.HandlerFunc(th.ListTasks)), th.authSvc), http.MethodGet))
}

type addTaskRequest struct {
	Content string `json:"content"`
}

// AddTask handle addtask
func (th *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var reqTask entities.Task
	if err := json.NewDecoder(r.Body).Decode(&reqTask); err != nil {
		returnErrorJSONResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	reqTask.UserID = r.Context().Value(ContextUserKey).(string)
	task, err := th.taskSvc.AddTask(r.Context(), reqTask)
	if err != nil {
		switch err {
		case services.ErrTaskLimitOfDayReached:
			returnErrorJSONResponse(w, err.Error(), http.StatusBadRequest)
			return
		default:
			returnErrorJSONResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	returnSuccessJSONResponse(w, task, http.StatusCreated)
	return
}

// ListTasks handle /tasks which list all tasks of request user
func (th *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserKey).(string)
	createdDate, reqParamValid := r.URL.Query()["created_date"]
	if !reqParamValid || len(createdDate[0]) < 1 {
		returnSuccessJSONResponse(w, []*entities.Task{}, 200)
		return
	}

	tasks, err := th.taskSvc.ListTasksByUserAndDate(r.Context(), userID, createdDate[0])
	if err != nil {
		if _, ok := err.(entities.TaskValidationError); ok {
			returnErrorJSONResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		returnErrorJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnSuccessJSONResponse(w, tasks, http.StatusOK)
	return
}
