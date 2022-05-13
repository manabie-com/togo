package controllers

import (
	"backend_test/task/models"
	"backend_test/task/services"
	"backend_test/utils"
	"encoding/json"
	"net/http"
)

func GetAllPaymentAPI(w http.ResponseWriter, r *http.Request) {

	task, err := services.GetAllTasksDetails()
	if err != nil {
		utils.WriteResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteResponse(w, task, http.StatusOK)
}

func GetPaymentAPI(w http.ResponseWriter, r *http.Request) {

	transaction_id := r.URL.Query().Get("id")

	tasks, err := services.GetTasksDetails(transaction_id)
	if err != nil {
		utils.WriteResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteResponse(w, tasks, http.StatusOK)
}

func SavePaymentAPI(w http.ResponseWriter, r *http.Request) {

	var tasks *models.Tasks
	_ = json.NewDecoder(r.Body).Decode(&tasks)

	err := services.ProcessTasks(tasks)
	if err != nil {
		utils.WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteResponse(w, tasks, http.StatusOK)
}
