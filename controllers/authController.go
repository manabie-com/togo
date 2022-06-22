package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/huynhhuuloc129/todo/models"
)

func Register(w http.ResponseWriter, r *http.Request) { // Handle register with method post
	var user, user1 models.NewUser
	var passwordCrypt string
	_ = json.NewDecoder(r.Body).Decode(&user)
	if _, ok := models.CheckUsername(user.Username); ok {
		http.Error(w, "This username already exist", http.StatusNotAcceptable)
		return
	}

	passwordCrypt, _ = models.Hash(user.Password)
	user1 = models.NewUser{
		Username: user.Username,
		Password: passwordCrypt,
	}

	err := models.InsertUser(user1)
	ErrorHandle(w, err, http.StatusBadRequest)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	ErrorHandle(w, err, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) { // handle login with method post
	var user models.NewUser
	var passwordCrypt string
	_ = json.NewDecoder(r.Body).Decode(&user)
	if _, ok := models.CheckUsername(user.Username); !ok {
		http.Error(w, "Account doesn't exist", http.StatusNotFound)
		return
	}
	if ok := models.CheckUser(user); !ok {
		http.Error(w, "Account input invalid", http.StatusNotFound)
		return
	}

	passwordCrypt, _ = models.Hash(user.Password)
	
}
