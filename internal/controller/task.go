package controller

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

type TaskController struct {
	TaskUsecase domain.TaskUsecase
	UserUsecase domain.UserUsecase
}

func NewTaskController(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) *TaskController {
	return &TaskController{
		TaskUsecase: taskUsecase,
		UserUsecase: userUsecase,
	}
}

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {

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

func (t *TaskController) GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.TaskUsecase.GetAllTask()
	if err != nil {
		// pkg.ERROR(w, http.StatusInternalServerError, err, "get all task failed")
		return
	}
	pkg.JSON(w, http.StatusOK, tasks)
}

func (t *TaskController) GetAllTaskOfUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	tasks, err := t.TaskUsecase.GetTasksByUsername(username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get all task of user failed!")
		return
	}
	pkg.JSON(w, http.StatusOK, tasks)
}

func (t *TaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
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

func (t *TaskController) CompleteTask(w http.ResponseWriter, r *http.Request) {
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

func (t *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	err = t.TaskUsecase.UpdateTask(id, username, r)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "update task failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: update task success")
}

func (t *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
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
