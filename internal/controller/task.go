package controller

import (
	"encoding/json"
	"fmt"
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/internal/repository"
	"lntvan166/togo/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	username := context.Get(r, "username").(string)
	id, err := repository.Repository.GetUserIDByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "get user id failed")
		return
	}

	isLimit, err := repository.Repository.GetLimitTaskToday(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "check limit task today failed")
		return
	}

	if isLimit {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("you have reached the limit of task today"), "")
		return
	}

	task := e.Task{}

	json.NewDecoder(r.Body).Decode(&task)

	task.CreatedAt = utils.GetCurrentTime()
	task.UserID = id
	err = repository.Repository.AddTask(&task)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "add task failed")
		return
	}

	numberTask, err := repository.Repository.GetNumberOfTaskToday(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "get number of task today failed")
		return
	}

	utils.JSON(w, http.StatusOK, map[string]interface{}{"number_task_today": numberTask, "message": "create task success"})

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.Repository.GetAllTask()
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "get all task failed")
		return
	}
	utils.JSON(w, http.StatusOK, tasks)
}

func GetAllTaskOfUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	tasks, err := repository.Repository.GetTaskByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "get all task of user failed!")
		return
	}
	utils.JSON(w, http.StatusOK, tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	task, err := repository.Repository.GetTaskByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "get task by id failed!")
		return
	}

	err = utils.CheckAccessPermission(w, username, task.UserID)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "check access permission failed: ")
		return
	}

	utils.JSON(w, http.StatusOK, task)
}

func CheckTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	user_id, err := repository.Repository.GetUserIDByTaskID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "task does not exist!")
		return
	}

	err = utils.CheckAccessPermission(w, username, user_id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "check access permission failed: ")
		return
	}

	err = repository.Repository.CheckTask(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "complete task failed!")
		return
	}
	utils.JSON(w, http.StatusOK, "message: check task success")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)
	user_id, err := repository.Repository.GetUserIDByTaskID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "task does not exist!")
		return
	}

	err = utils.CheckAccessPermission(w, username, user_id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "check access permission failed: ")
		return
	}

	task, err := repository.Repository.GetTaskByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "get task by id failed!")
		return
	}

	json.NewDecoder(r.Body).Decode(&task)

	err = repository.Repository.UpdateTask(task)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "update task failed!")
		return
	}

	utils.JSON(w, http.StatusOK, "message: update task success")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)
	user_id, err := repository.Repository.GetUserIDByTaskID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "task does not exist!")
		return
	}

	err = utils.CheckAccessPermission(w, username, user_id)
	if err != nil {
		return
	}

	err = repository.Repository.DeleteTask(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "delete task failed!")
		return
	}

	utils.JSON(w, http.StatusOK, "message: delete task success")
}
