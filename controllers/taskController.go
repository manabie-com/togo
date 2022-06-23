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

func ResponeAllTask(w http.ResponseWriter, r *http.Request) { // Get all user from database
	userid, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "userid"))) // get userid from login 

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

func ResponeOneTask(w http.ResponseWriter, r *http.Request) { // Get one user from database
	userid, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "userid"))) // get userid from login 
	id, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "id"))) // get id from url 

	task, ok := models.CheckIDTaskAndReturn(w, id, userid)
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

func CreateTask(w http.ResponseWriter, r *http.Request) { // Create a new user
	userid, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "userid"))) // get id from url 

	var task models.NewTask
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}
	task.UserId = userid
	task.Status = "pending"
	task.Time = time.Now()
	ok := models.CheckTaskInput(task)
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

func DeleteTask(w http.ResponseWriter, r *http.Request ){ // Delete one user from database
	userid, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "userid")))// get userid from login 
	id, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "id"))) // get id from url 

	_, ok := models.CheckIDTaskAndReturn(w, id, userid)
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

func UpdateTask(w http.ResponseWriter, r *http.Request) { // Update one user already exist in database
	userid, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "userid"))) // get userid from login 
	id, _ := strconv.Atoi(fmt.Sprintf("%v",context.Get(r, "id"))) // get id from url 

	var newTask models.NewTask
	oldTask, ok := models.CheckIDTaskAndReturn(w, id, userid) // check id task have on database or not
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "decode failed, input invalid", http.StatusBadRequest)
		return
	}
	if strings.ToLower(newTask.Status) == "done" {
		newTask.TimeDone = time.Now()
	} else {
		http.Error(w, "status can only be done or pending", http.StatusBadRequest)
		return
	}

	err = models.UpdateTask(newTask, id, userid)
	oldTask.Status = newTask.Status
	oldTask.TimeDone = newTask.TimeDone
	if err != nil {
		http.Error(w, "update task failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(oldTask)
	if err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}
