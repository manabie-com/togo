package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var GetTasks = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get user id here
	decoded := r.Context().Value("user").(*models.Token)
	task := &models.Task{
		UserId: decoded.UserId,
	}
	tasks, err := task.GetTasksByUserId(db)
	if err != nil {
		u.SuccessRespond(w, http.StatusOK, "Success", nil)
		return
	}
	//Everything OK
	u.SuccessRespond(w, http.StatusOK, "Success", tasks)
}

var GetTask = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// decode token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	// convert string id to uint32
	id, err := u.Str2Uint32(mux.Vars(r)["id"])
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "ID must be a number type.")
		return
	}

	task := &models.Task{
		ID:     id,
		UserId: decoded.UserId,
	}
	err = task.GetTaskByUserId(db)
	if err != nil {
		u.SuccessRespond(w, http.StatusOK, "Not found task", nil)
		return
	}

	u.SuccessRespond(w, http.StatusOK, "Success", map[string]interface{}{
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})
}

var Add = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// decode token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	task := &models.Task{
		UserId: decoded.UserId,
	}
	// json body -> task object
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "invalid request")
		return
	}
	// validate task object
	validate := validator.New()
	err = validate.Struct(task)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, err.Error())
		return
	}
	// check today tasks limit
	if task.IsLimit(db, decoded.LimitDayTasks) {
		u.FailureRespond(w, http.StatusBadRequest, "Today tasks had limited, Please Comeback tomorrow.")
		return
	}
	// insert database
	err = task.InsertTask(db)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, err.Error())
		return
	}
	//everything OK
	u.SuccessRespond(w, http.StatusCreated, "Success create task", map[string]interface{}{
		"id":         fmt.Sprint(task.ID),
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})
}

var Edit = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoded := r.Context().Value("user").(*models.Token)
	// convert string id to uint32
	id, err := u.Str2Uint32(mux.Vars(r)["id"])
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "ID must be a number type.")
		return
	}

	task := &models.Task{
		ID:     id,
		UserId: decoded.UserId,
	}

	// err := GetTaskByUserId(db)
	err = json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, err.Error())
		return
	}
	var (
		newName    string = task.Name
		newContent string = task.Content
	)
	err = task.GetTaskByUserId(db)
	if err != nil {
		u.FailureRespond(w, http.StatusInternalServerError, "Somethings went wrong. Please try again"+err.Error())
		return
	}
	if newName != "" {
		task.Name = newName
	}
	if newContent != "" {
		task.Content = newContent
	}
	err = task.UpdateTaskById(db)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, err.Error())
		return
	}
	u.SuccessRespond(w, http.StatusOK, "Success update task", map[string]interface{}{
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})
}

var Delete = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoded := r.Context().Value("user").(*models.Token)
	// convert string id to uint32
	id, err := u.Str2Uint32(mux.Vars(r)["id"])
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "ID must be a number type.")
		return
	}

	task := &models.Task{
		ID:     id,
		UserId: decoded.UserId,
	}

	err = task.DeleteTaskById(db)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, err.Error())
		return
	}

	u.SuccessRespond(w, http.StatusNoContent, "Success delete task", nil)
}
