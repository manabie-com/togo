package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/huynhhuuloc129/todo/models"
)

// Get all user from database
func (bh *BaseHandler) ResponseAllTask(w http.ResponseWriter, r *http.Request) {
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login

	tasks, err := bh.BaseCtrl.GetAllTasks(userid)
	if err != nil {
		http.Error(w, "get all task failed", http.StatusFailedDependency)
		return
	}

	// tasks = ChangeStatusAllTasksAfterDay(tasks)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "encode tasks failed", 500)
		return
	}
}

// Get one user from database
func (bh *BaseHandler) ResponseOneTask(w http.ResponseWriter, r *http.Request) {
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id")))         // get id from url

	task, ok := bh.BaseCtrl.FindTaskByID(id, userid)
	if !ok {
		http.Error(w, "id invalid", http.StatusBadRequest)
		return
	}

	// task = ChangeStatusOneTaskAfterDay(task)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "encode task failed", http.StatusFailedDependency)
		return
	}
}

// Create a new user
func (bh *BaseHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login

	var task models.NewTask
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "decode failed, err: "+err.Error(), http.StatusFailedDependency)
		return
	}

	task.UserId = userid
	task.Status = "pending"
	task.TimeDone = time.Date(0001, 01, 01, 0, 0, 0, 0, time.Local).Round(0)
	year, month, day := task.Time.Date()
	if year != time.Now().Year() || month != time.Now().Month() || day != time.Now().Day() {
		task.Time = time.Now()
	}
	if ok := models.CheckTaskInput(task); !ok {
		http.Error(w, "task field invalid", http.StatusBadRequest)
	}
	if err := bh.BaseCtrl.InsertTask(task); err != nil {
		http.Error(w, "insert task failed, err: "+err.Error(), http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "encode failed, err: "+err.Error(), http.StatusCreated)
		return
	}
}

// Delete one user from database
func (bh *BaseHandler) DeleteFromTask(w http.ResponseWriter, r *http.Request) {
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id")))         // get id from url

	if _, ok := bh.BaseCtrl.FindTaskByID(id, userid); !ok { // Check task id exist or not and return that task
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	if err := bh.BaseCtrl.DeleteTask(id, userid); err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)
		return
	}
	w.Write([]byte("message: delete success"))
}

// Update one user already exist in database
func (bh *BaseHandler) UpdateToTask(w http.ResponseWriter, r *http.Request) {
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id")))         // get id from url

	oldTask, ok := bh.BaseCtrl.FindTaskByID(id, userid) // Check task id exist or not and return that task
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&oldTask); err != nil { // write on old task
		http.Error(w, "decode failed, input invalid, err: " +err.Error(), http.StatusBadRequest)
		return
	}

	if strings.ToLower(oldTask.Status) == "done" { // check status and insert time to it
		oldTask.TimeDone = time.Now()
	} else if strings.ToLower(oldTask.Status) == "pending" {
		oldTask.TimeDone = time.Date(0001, 01, 01, 0, 0, 0, 0, time.Local).Round(0)
	} else {
		http.Error(w, "status can only be done or pending", http.StatusBadRequest)
		return
	}

	if err := bh.BaseCtrl.UpdateTask(oldTask, id, userid); err != nil {
		http.Error(w, "update task failed, err: "+ err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(oldTask); err != nil {
		http.Error(w, "encode failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}
}
