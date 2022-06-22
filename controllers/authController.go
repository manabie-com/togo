package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huynhhuuloc129/todo/jwt"
	"github.com/huynhhuuloc129/todo/models"
)
type responseToken struct{ //response token
	Message string
	Token string
}

func Register(w http.ResponseWriter, r *http.Request) { // Handle register with method post
	var user, user1 models.NewUser
	var passwordCrypt string
	_ = json.NewDecoder(r.Body).Decode(&user)
	_, ok := models.CheckUsername(user.Username);

	if  ok { // Check username exist or not
		http.Error(w, "this username already exist", http.StatusNotAcceptable)
		return
	}

	passwordCrypt, _ = models.Hash(user.Password) // hash password
	user1 = models.NewUser{
		Username: user.Username,
		Password: passwordCrypt,
	}

	err := models.InsertUser(user1) // insert new user to database
	if err != nil {
		http.Error(w, "insert user failed.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user) // response message and token back to view
	if err != nil {
		http.Error(w, "encode failed.", http.StatusCreated)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) { // handle login with method post
	var user models.NewUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	user1, ok := models.CheckUsername(user.Username)
	if !ok { // check username exist or not
		http.Error(w, "account doesn't exist", http.StatusNotFound)
		return
	}
	if ok := models.CheckUser(user); !ok { // check if user input valid or not
		http.Error(w, "account input invalid", http.StatusNotFound)
		return
	}
	fmt.Println(user1.Password, user.Password)
	err := models.CheckPasswordHash(user1.Password ,user.Password) // check if password correct or not
	if err != nil {
		http.Error(w, "password incorrect", http.StatusAccepted)
		return
	}
	
	token, err := jwt.Create(w, user.Username) // Create token
	if err != nil {
		http.Error(w,  "internal server error", 500)
		return
	}
	jwt.CheckToken(token)

	w.Header().Set("Content-Type", "application/json")
	resToken := responseToken{ 
		Message: "login success",
		Token: token,
	}
	err = json.NewEncoder(w).Encode(resToken)// response token back to client
	if err != nil {
		http.Error(w, "encode failed", http.StatusFailedDependency)
	}
}
