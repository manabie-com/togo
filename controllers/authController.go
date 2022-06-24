package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/huynhhuuloc129/todo/jwt"
	"github.com/huynhhuuloc129/todo/models"
)

type responseToken struct { //response token
	Message string
	Token   string
}

func Register(w http.ResponseWriter, r *http.Request) { // Handle register with method post
	var user, user1 models.NewUser
	var passwordCrypt string

	_ = json.NewDecoder(r.Body).Decode(&user)
	if _, ok := models.CheckUserInput(user.Username); ok { // Check username exist or not
		http.Error(w, "this username already exist", http.StatusNotAcceptable)
		return
	}

	passwordCrypt, _ = models.Hash(user.Password) // hash password
	user1 = models.NewUser{
		Username: user.Username,
		Password: passwordCrypt,
	}

	if strings.ToLower(user1.Username) != "admin" {
		user1.LimitTask = 10
	} else {
		user1.LimitTask = 0
	}

	if err := models.InsertUser(user1); err != nil { // insert new user to database
		http.Error(w, "insert user failed.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil { // response message and token back to view
		http.Error(w, "encode failed.", http.StatusCreated)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) { // handle login with method post
	var user models.NewUser

	_ = json.NewDecoder(r.Body).Decode(&user)
	user1, ok := models.CheckUserInput(user.Username)
	if !ok { // check username exist or not
		http.Error(w, "account doesn't exist", http.StatusNotFound)
		return
	}

	if ok := models.CheckUser(user); !ok { // check if user input valid or not
		http.Error(w, "account input invalid", http.StatusNotFound)
		return
	}

	if err := models.CheckPasswordHash(user1.Password, user.Password); err != nil { // check password correct or not
		http.Error(w, "password incorrect", http.StatusAccepted)
		return
	}

	token, err := jwt.Create(w, user.Username, int(user1.Id)) // Create token
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resToken := responseToken{
		Message: "login success",
		Token:   token,
	}

	if err = json.NewEncoder(w).Encode(resToken); err != nil { // response token back to client
		http.Error(w, "encode failed", http.StatusFailedDependency)
	}
}
