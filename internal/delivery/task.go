package delivery

import (
	"encoding/json"
	"lntvan166/togo/internal/domain"
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type TaskDelivery struct {
	TaskUsecase domain.TaskUsecase
	UserUsecase domain.UserUsecase
}

func NewTaskDelivery(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) *TaskDelivery {
	return &TaskDelivery{
		TaskUsecase: taskUsecase,
		UserUsecase: userUsecase,
	}
}

func (t *TaskDelivery) CreateTask(w http.ResponseWriter, r *http.Request) {

	username := context.Get(r, "username").(string)

	task := e.Task{}

	json.NewDecoder(r.Body).Decode(&task)

	numberTask, err := t.TaskUsecase.CreateTask(&task, username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "create task failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, map[string]interface{}{"number_task_today": numberTask, "message": "create task success"})

}

func (t *TaskDelivery) GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.TaskUsecase.GetAllTask()
	if err != nil {
		// pkg.ERROR(w, http.StatusInternalServerError, err, "get all task failed")
		return
	}
	pkg.JSON(w, http.StatusOK, tasks)
}

func (t *TaskDelivery) GetAllTaskOfUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	tasks, err := t.TaskUsecase.GetTasksByUsername(username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get all task of user failed!")
		return
	}
	pkg.JSON(w, http.StatusOK, tasks)
}

func (t *TaskDelivery) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	task, err := t.TaskUsecase.GetTaskByID(id, username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get task by id failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, task)
}

func (t *TaskDelivery) CompleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	err = t.TaskUsecase.CompleteTask(id, username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "complete task failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: check task success")
}

func (t *TaskDelivery) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	err = t.TaskUsecase.DeleteTask(id, username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "delete task failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: delete task success")
}
