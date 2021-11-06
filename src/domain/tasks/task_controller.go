package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

// TaskController ...
type TaskController struct {
	TaskCreateAction TaskCreateAction
	TaskListAction   TaskListAction
}

// name...
func (ctrl TaskController) name() string {
	return "task.TaskController"
}

// Create ...
func (ctrl TaskController) Create(w http.ResponseWriter, r *http.Request) {
	payload := &TaskCreatePayload{}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid payload")
		return
	}

	if validErrs := payload.Validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validErrs)
		return
	}

	taskDetail, err := ctrl.TaskCreateAction.Execute(payload.Content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("While creating task error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskDetail)
}

// List ...
func (ctrl TaskController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	tasks, err := ctrl.TaskListAction.Execute(token.AccessUser.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("While getting list task error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
