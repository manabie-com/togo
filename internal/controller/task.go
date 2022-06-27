package controller

import (
	"encoding/json"
	"fmt"
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/internal/usecase"
	"lntvan166/togo/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	username := context.Get(r, "username").(string)
	id, err := usecase.GetUserIDByUsername(username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get user id failed")
		return
	}

	isLimit, err := usecase.CheckLimitTaskToday(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "check limit task today failed")
		return
	}

	if isLimit {
		pkg.ERROR(w, http.StatusBadRequest, fmt.Errorf("you have reached the limit of task today"), "")
		return
	}

	task := e.Task{}

	json.NewDecoder(r.Body).Decode(&task)

	task.CreatedAt = pkg.GetCurrentTime()
	task.UserID = id
	err = usecase.AddTask(&task)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "add task failed")
		return
	}

	numberTask, err := usecase.GetNumberOfTaskTodayByUserID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get number of task today failed")
		return
	}

	pkg.JSON(w, http.StatusOK, map[string]interface{}{"number_task_today": numberTask, "message": "create task success"})

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := usecase.GetAllTask()
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get all task failed")
		return
	}
	pkg.JSON(w, http.StatusOK, tasks)
}

func GetAllTaskOfUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	tasks, err := usecase.GetTasksByUsername(username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get all task of user failed!")
		return
	}
	pkg.JSON(w, http.StatusOK, tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	task, err := usecase.GetTaskByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get task by id failed!")
		return
	}

	err = usecase.CheckAccessPermission(w, username, task.UserID)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "check access permission failed: ")
		return
	}

	pkg.JSON(w, http.StatusOK, task)
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)

	user_id, err := usecase.GetUserIDByTaskID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "task does not exist!")
		return
	}

	err = usecase.CheckAccessPermission(w, username, user_id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "check access permission failed: ")
		return
	}

	err = usecase.CompleteTask(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "complete task failed!")
		return
	}
	pkg.JSON(w, http.StatusOK, "message: check task success")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)
	user_id, err := usecase.GetUserIDByTaskID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "task does not exist!")
		return
	}

	err = usecase.CheckAccessPermission(w, username, user_id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "check access permission failed: ")
		return
	}

	task, err := usecase.GetTaskByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "get task by id failed!")
		return
	}

	json.NewDecoder(r.Body).Decode(&task)

	err = usecase.UpdateTask(task)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "update task failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: update task success")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "id is not a number!")
		return
	}

	username := context.Get(r, "username").(string)
	user_id, err := usecase.GetUserIDByTaskID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "task does not exist!")
		return
	}

	err = usecase.CheckAccessPermission(w, username, user_id)
	if err != nil {
		return
	}

	err = usecase.DeleteTask(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "delete task failed!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: delete task success")
}
