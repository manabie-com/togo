package task

import (
	"encoding/json"
	"fmt"
	e "lntvan166/togo/entities"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	username := context.Get(r, "username").(string)
	id, err := model.GetUserIDByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	isLimit, err := model.CheckLimitTaskToday(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	if isLimit {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("you have reached the limit of task today"))
		return
	}

	task := e.Task{}

	json.NewDecoder(r.Body).Decode(&task)

	task.CreatedAt = utils.GetCurrentTime()
	task.UserID = id
	err = model.AddTask(&task)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	numberTask, err := model.GetNumberOfTaskToday(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	utils.JSON(w, http.StatusOK, map[string]interface{}{"number_task_today": numberTask, "message": "create task success"})

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := model.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSON(w, http.StatusOK, tasks)
}

func GetAllTaskOfUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	tasks, err := model.GetTaskByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}
	utils.JSON(w, http.StatusOK, tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf(err.Error()))
		return
	}

	username := context.Get(r, "username").(string)

	task, err := model.GetTaskByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	err = utils.CheckAccessPermission(w, username, task.UserID)
	if err != nil {
		return
	}

	utils.JSON(w, http.StatusOK, task)
}

func CheckTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf(err.Error()))
		return
	}

	username := context.Get(r, "username").(string)

	user_id, err := model.GetUserIDByTaskID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	err = utils.CheckAccessPermission(w, username, user_id)
	if err != nil {
		return
	}

	err = model.CheckTask(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}
	utils.JSON(w, http.StatusOK, "message: check task success")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf(err.Error()))
		return
	}

	username := context.Get(r, "username").(string)
	user_id, err := model.GetUserIDByTaskID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	err = utils.CheckAccessPermission(w, username, user_id)
	if err != nil {
		return
	}

	err = model.DeleteTask(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	utils.JSON(w, http.StatusOK, "message: delete task success")
}
