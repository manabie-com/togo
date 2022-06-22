package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/huynhhuuloc129/todo/models"
)

func ResponeAllUser(w http.ResponseWriter, r *http.Request) { // Get all user from database
	w.Header().Set("Content-Type", "application/json")
	users, err := models.GetAllUser()
	ErrorHandle(w, err, http.StatusFailedDependency)

	err = json.NewEncoder(w).Encode(users)
	ErrorHandle(w, err, 500)
}

func ResponeOneUser(w http.ResponseWriter, r *http.Request, id int) { // Get one user from database
	user, err := models.GetOneUser(id)
	ErrorHandle(w, err, http.StatusFailedDependency)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(user)
	ErrorHandle(w, err, http.StatusFailedDependency)
}

func CreateUser(w http.ResponseWriter, r *http.Request) { // Create a new user
	var user models.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	ErrorHandle(w, err, http.StatusFailedDependency)

	err = models.InsertUser(user)
	ErrorHandle(w, err, http.StatusFailedDependency)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	ErrorHandle(w, err, http.StatusCreated)

}

func DeleteUser(w http.ResponseWriter, r *http.Request, id int) { // Delete one user from database
	_, ok := models.CheckID(id)
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	err := models.DeleteUser(id)
	ErrorHandle(w, err, http.StatusFailedDependency)
	w.Write([]byte("message: delete success"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request, id int) { // Update one user already exist in database
	var newUser models.NewUser;
	_, ok := models.CheckID(id)
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	ErrorHandle(w, err, http.StatusBadRequest)

	err = models.UpdateUser(newUser, id)
	ErrorHandle(w, err, http.StatusBadRequest)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)
	ErrorHandle(w, err, http.StatusBadRequest)
}

func ErrorHandle(w http.ResponseWriter, err error, status int) {
	if err != nil {
		http.Error(w, err.Error(), status)
		os.Exit(1)
	}
}
