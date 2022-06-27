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
	users, err := models.Repo.GetAllUser()
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
	user, ok := models.Repo.FindUserByID(id)
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
	if _, ok := models.Repo.CheckUserNameExist(user.Username); ok { // Check username exist or not
		http.Error(w, "this username already exist", http.StatusNotAcceptable)
		return
	}

	if strings.ToLower(user.Username) != "admin" { // check admin or not
		user.LimitTask = 10
	} else {
		user.LimitTask = 0
	}

	user.Password, _ = models.Hash(user.Password)
	if err := models.Repo.InsertUser(user); err != nil { // insert user to database
		http.Error(w, "insert user failed", http.StatusFailedDependency)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil { // return response
		http.Error(w, "encode failed", http.StatusCreated)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) { // Delete one user from database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url
	if _, ok := models.Repo.FindUserByID(id); !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	if err := models.Repo.DeleteAllTaskFromUser(id); err != nil {
		http.Error(w, "delete all task of user failed", http.StatusFailedDependency)
		return
	}
	if err := models.Repo.DeleteUser(id); err != nil {
		http.Error(w, "delete user failed", http.StatusFailedDependency)
		return
	}
	w.Write([]byte("message: delete success"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) { // Update one user already exist in database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url

	var newUser models.NewUser
	user, ok := models.Repo.FindUserByID(id)
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
	if err := models.Repo.UpdateUser(newUser, id); err != nil {
		http.Error(w, "update user failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}
