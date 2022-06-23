package task

import (
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := model.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.JSON(w, http.StatusOK, tasks)
}
