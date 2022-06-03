package controller

import (
	"encoding/json"
	"net/http"
	"time"
	"togo/common/response"
	"togo/models"
	"togo/service"
)

// Define an interface for the `User` controller
// `Register` and `Login` will be used by the endpoints for `User` related actions
type UserController interface {
	// Accept POST body, validate `User`, then register new `User`
	Register(w http.ResponseWriter, r *http.Request)

	// Accept POST body, validate `User`, generate JWT token, set cookie, then Login
	Login(w http.ResponseWriter, r *http.Request)
}

// Define a Controller struct that contains
// the `User` Service (business logic for `User`) attribute
type usercontroller struct {
	userservice service.UserService
}

// Define a Constructor
// Dependency Injection for `User` Controller
func NewUserController(service service.UserService) UserController {
	return &usercontroller{
		userservice: service,
	}
}

// Register a new `User`
// Route: POST /register
// Access: public
func (c *usercontroller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get POST body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Validate User
	err = c.userservice.ValidateRegistration(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Register User
	_, err = c.userservice.Register(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Status: "Success",
		Code:   http.StatusOK,
		Data:   map[string]int{"tasks per day": user.Limit},
	})
}

// Login existing `User`
// Route: POST /login
// Access: public
func (c *usercontroller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get POST body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Validate User Login
	err = c.userservice.ValidateLogin(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Set expiration time of the token
	expiration := time.Now().Add(time.Minute * 60)

	// Generate JWT token
	// Return token string to set Cookie
	token, err := c.userservice.GenerateJWT(&user, expiration)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Set cookie
	http.SetCookie(w,
		&http.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  expiration,
			HttpOnly: true,
		})

	// Login
	c.userservice.Login(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Status: "Success",
		Code:   http.StatusOK,
	})
}
