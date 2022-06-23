package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/huynhhuuloc129/todo/models"
)

const (
	dateformat = "01-02-2006"
)


func ResponeAllTask(w http.ResponseWriter, r *http.Request, userid int) { // Get all user from database
	tasks, err := models.GetAllTasks(userid)
	if err != nil {
		http.Error(w, "get all task failed", http.StatusFailedDependency)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "encode tasks failed", 500)
		return
	}
}

func ResponeOneTask(w http.ResponseWriter, r *http.Request, id int, userid int) { // Get one user from database
	task, ok := models.CheckIDTask(w, id, userid)
	if !ok {
		http.Error(w, "id invalid", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "encode task failed", http.StatusFailedDependency)
		return
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request, userid int) { // Create a new user
	var task models.NewTask
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}
	task.UserId = userid
	task.Status = "pending"
	task.Time = time.Now().Format(dateformat)
	ok := models.CheckTask(task)
	if !ok {
		http.Error(w, "task field invalid", http.StatusBadRequest)
	}

	err = models.InsertTask(task)
	if err != nil {
		http.Error(w, "insert task failed", http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "encode failed", http.StatusCreated)
		return
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request, id int, userid int) { // Delete one user from database
	_, ok := models.CheckIDTask(w, id, userid)
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	err := models.DeleteTask(id, userid)
	if err != nil {
		http.Error(w, "delete task failed", http.StatusFailedDependency)
		return
	}
	w.Write([]byte("message: delete success"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request, id int, userid int) { // Update one user already exist in database
	var newTask models.NewTask
	_, ok := models.CheckIDTask(w, id, userid)
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "decode failed", http.StatusBadRequest)
		return
	}
	if strings.ToLower(newTask.Status) == "done" {
		newTask.TimeDone = time.Now().Format(dateformat)
	} else {
		newTask.Status = "pending"
	}

	err = models.UpdateTask(newTask, id, userid)
	if err != nil {
		http.Error(w, "update task failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}
