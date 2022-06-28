package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/huynhhuuloc129/todo/models"
)

// Get all user from database
func (bh *BaseHandler) ResponseAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := bh.BaseCtrl.GetAllUser()
	if err != nil {
		http.Error(w, "get all user failed", http.StatusFailedDependency)
		return
	}

	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "encode failed", 500)
		return
	}
}

// Get one user from database
func (bh *BaseHandler) ResponseOneUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url
	user, ok := bh.BaseCtrl.FindUserByID(id)
	if !ok {
		http.Error(w, "id invalid", http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "encode failed", http.StatusFailedDependency)
		return
	}
}

// Create a new user
func (bh *BaseHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.NewUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}

	if strings.ToLower(user.Username) != "admin" { // check admin or not
		user.LimitTask = 10
	} else {
		user.LimitTask = 0
	}

	if err := bh.BaseCtrl.InsertUser(user); err != nil { // insert user to database
		http.Error(w, "insert user failed, err: "+err.Error(), http.StatusFailedDependency)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil { // return response
		http.Error(w, "encode failed, err: "+err.Error(), http.StatusCreated)
		return
	}
}

// Delete user from database
func (bh *BaseHandler) DeleteFromUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url
	if _, ok := bh.BaseCtrl.FindUserByID(id); !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	if err := bh.BaseCtrl.DeleteAllTaskFromUser(id); err != nil {
		http.Error(w, "delete all task of user failed", http.StatusFailedDependency)
		return
	}
	if err := bh.BaseCtrl.DeleteUser(id); err != nil {
		http.Error(w, "delete user failed", http.StatusFailedDependency)
		return
	}
	w.Write([]byte("message: delete success"))
}

// Update one user already exist in database
func (bh *BaseHandler) UpdateToUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url

	var newUser models.NewUser
	user, ok := bh.BaseCtrl.FindUserByID(id)
	newUser.Username = user.Username
	newUser.Password = user.Password
	newUser.LimitTask = user.LimitTask

	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "decode failed, input invalid", http.StatusBadRequest)
		return
	}

	if newUser.Password != user.Password {
		newUser.Password, _ = models.Hash(newUser.Password)
	}
	if err := bh.BaseCtrl.UpdateUser(newUser, id); err != nil {
		http.Error(w, "update user failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}
