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
func (h *BaseHandler)ResponseAllTask(w http.ResponseWriter, r *http.Request) { 
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login

	tasks, err := models.GetAllTasks(h.DB, userid)
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
func (h *BaseHandler)ResponseOneTask(w http.ResponseWriter, r *http.Request) {
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id")))         // get id from url

	task, ok := models.FindTaskByID(h.DB, id, userid)
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
func (h *BaseHandler)CreateTask(w http.ResponseWriter, r *http.Request) { 
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	if ok, err := models.CheckLimitTaskUser(h.DB, userid); !ok {
		if err != nil{
			http.Error(w, err.Error(), http.StatusFailedDependency)
			return
		}
		http.Error(w, "The limit of today is full or error occur during check limit", http.StatusFailedDependency)
		return
	}

	var task models.NewTask
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}

	task.UserId = userid
	task.Status = "pending"
	task.Time = time.Now()
	if ok := models.CheckTaskInput(task); !ok {
		http.Error(w, "task field invalid", http.StatusBadRequest)
	}
	if err := models.InsertTask(h.DB, task); err != nil {
		http.Error(w, "insert task failed", http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "encode failed", http.StatusCreated)
		return
	}
}

// Delete one user from database
func (h *BaseHandler)DeleteTask(w http.ResponseWriter, r *http.Request) { 
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id")))         // get id from url

	if _, ok := models.FindTaskByID(h.DB, id, userid); !ok { // Check task id exist or not and return that task
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	if err := models.DeleteTask(h.DB, id, userid); err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)
		return
	}
	w.Write([]byte("message: delete success"))
}

// Update one user already exist in database
func (h *BaseHandler)UpdateEntireTask(w http.ResponseWriter, r *http.Request) { 
	userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id")))         // get id from url

	oldTask, ok := models.FindTaskByID(h.DB, id, userid) // Check task id exist or not and return that task
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&oldTask); err != nil { // write on old task
		http.Error(w, "decode failed, input invalid", http.StatusBadRequest)
		return
	}

	if strings.ToLower(oldTask.Status) == "done" { // check status and insert time to it
		oldTask.TimeDone = time.Now()
	} else if strings.ToLower(oldTask.Status) == "pending" {
		oldTask.TimeDone = time.Date(0001, 01, 01, 0, 0, 0, 0, time.Local)
	} else {
		http.Error(w, "status can only be done or pending", http.StatusBadRequest)
		return
	}

	if err := models.UpdateTask(h.DB, oldTask, id, userid); err != nil {
		http.Error(w, "update task failed", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(oldTask); err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}


