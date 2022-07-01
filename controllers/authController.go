package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/huynhhuuloc129/todo/jwt"
	"github.com/huynhhuuloc129/todo/models"
)

type BaseHandler struct {
	BaseCtrl *models.DbConn
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(BC *models.DbConn) *BaseHandler {
	return &BaseHandler{
		BaseCtrl: BC,
	}
}

//response token
type ResponseToken struct {
	Message string
	Token   string
}

// Handle register with method post
func (h *BaseHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user, user1 models.NewUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	user1 = models.NewUser{
		Username: user.Username,
		Password: user.Password,
	}
	ok := models.CheckUserInput(user1)
	if !ok {
		http.Error(w, "registered failed", http.StatusBadRequest)
		return
	}

	if strings.ToLower(user1.Username) != "admin" {
		user1.LimitTask = 10
	} else {
		user1.LimitTask = 0
	}
	if err := h.BaseCtrl.InsertUser(user1); err != nil { // insert new user to database
		http.Error(w, "insert user failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user1); err != nil { // response message and token back to view
		http.Error(w, "encode failed, err: "+err.Error(), http.StatusCreated)
		return
	}
}

// handle login with method post
func (h *BaseHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.NewUser

	_ = json.NewDecoder(r.Body).Decode(&user)
	user1, ok := h.BaseCtrl.CheckUserNameExist(user.Username)
	if !ok { // check username exist or not
		http.Error(w, "account doesn't exist", http.StatusNotFound)
		return
	}

	if ok := models.CheckUserInput(user); !ok { // check if user input valid or not
		http.Error(w, "account input invalid", http.StatusNotFound)
		return
	}
	if err := models.CheckPasswordHash(user1.Password, user.Password); err != nil { // check password correct or not
		http.Error(w, "password incorrect, err: "+err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := jwt.Create(user.Username, int(user1.Id)) // Create token
	if err != nil {
		http.Error(w, "internal server error, err: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resToken := ResponseToken{
		Message: "login success",
		Token:   token,
	}

	if err = json.NewEncoder(w).Encode(resToken); err != nil { // response token back to client
		http.Error(w, "encode failed, err: " + err.Error(), http.StatusFailedDependency)
	}
}
