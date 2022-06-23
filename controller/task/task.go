package task

import (
	"fmt"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"

	"github.com/gorilla/context"
)

type ctxData string

const ctxUsername ctxData = "username"

func CreateTask(w http.ResponseWriter, r *http.Request) {

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
