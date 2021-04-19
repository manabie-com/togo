package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/services"

)

type TaskController struct {
	services.ITaskService
}

func (controller *TaskController) ListTasks(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)

	var ok bool
	req, ok = services.ValidToken(req)
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}	

	id, _ := services.UserIDFromCtx(req.Context())
	createTime := req.FormValue("created_date")

	tasks, err := controller.FindByIdAndTimeFromService(id, createTime)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]storages.Task{
		"data": tasks,
	})
}

func (controller *TaskController) AddTask(resp http.ResponseWriter, req *http.Request) {

	log.Println(req.Method, req.URL.Path)

	var ok bool
	req, ok = services.ValidToken(req)
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}		

	t := storages.Task{}
	err := json.NewDecoder(req.Body).Decode(&t)

	defer req.Body.Close()
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := services.UserIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	err, error_code := controller.StoreFromService(t)
	if err != nil {
		resp.WriteHeader(error_code)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]storages.Task{
		"data": t,
	})
}
