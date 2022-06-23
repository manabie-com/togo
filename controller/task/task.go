package task

import (
	"encoding/json"
	"fmt"
	e "lntvan166/togo/entities"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"

	"github.com/gorilla/context"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	id, err := model.GetUserIDByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	task := e.Task{}

	json.NewDecoder(r.Body).Decode(&task)

	task.CreatedAt = utils.GetCurrentTime()
	task.UserID = id
	fmt.Println(task)
	err = model.AddTask(&task)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}
	utils.JSON(w, http.StatusOK, task)

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := model.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSON(w, http.StatusOK, tasks)
}

func GetTaskForUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	tasks, err := model.GetTaskByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}
	utils.JSON(w, http.StatusOK, tasks)
}
