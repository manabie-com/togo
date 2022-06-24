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

func ResponeAllUser(w http.ResponseWriter, r *http.Request) { // Get all user from database
	w.Header().Set("Content-Type", "application/json")
	users, err := models.GetAllUser()
	if err != nil {
		http.Error(w, "get all user failed", http.StatusFailedDependency)
		return
	}

	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "encode failed", 500)
		return
	}
}

func ResponeOneUser(w http.ResponseWriter, r *http.Request) { // Get one user from database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url
	user, ok := models.CheckIDUserAndReturn(id)
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

func CreateUser(w http.ResponseWriter, r *http.Request) { // Create a new user
	var user models.NewUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}
	if _, ok := models.CheckUserInput(user.Username); ok { // Check username exist or not
		http.Error(w, "this username already exist", http.StatusNotAcceptable)
		return
	}

	if strings.ToLower(user.Username) != "admin" {
		user.LimitTask = 10
	} else {
		user.LimitTask = 0
	}

	user.Password, _ = models.Hash(user.Password)
	if err := models.InsertUser(user); err != nil {
		http.Error(w, "insert user failed", http.StatusFailedDependency)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "encode failed", http.StatusCreated)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) { // Delete one user from database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url
	if _, ok := models.CheckIDUserAndReturn(id); !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	if err := models.DeleteAllTaskFromUser(id); err != nil {
		http.Error(w, "delete all task of user failed", http.StatusFailedDependency)
		return
	}
	if err := models.DeleteUser(id); err != nil {
		http.Error(w, "delete user failed", http.StatusFailedDependency)
		return
	}
	w.Write([]byte("message: delete success"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) { // Update one user already exist in database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url

	var newUser models.NewUser
	user, ok := models.CheckIDUserAndReturn(id)
	newUser.Username = user.Username
	newUser.Password = user.Password

	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "decode failed", http.StatusBadRequest)
		return
	}

	if newUser.Password != user.Password {
		newUser.Password, _ = models.Hash(newUser.Password)
	}
	if err := models.UpdateUser(newUser, id); err != nil {
		http.Error(w, "update user failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}
