package controller

import (
	"encoding/json"
	"net/http"
	"time"
	"togo/common/response"
	"togo/models"
	"togo/service"

	"github.com/rs/zerolog"
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
	logger      zerolog.Logger
}

// Define a Constructor
// Dependency Injection for `User` Controller
func NewUserController(service service.UserService, logger zerolog.Logger) UserController {
	return &usercontroller{
		userservice: service,
		logger:      logger,
	}
}

// Register a new `User`
// Route: POST /register
// Access: public
func (c *usercontroller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get POST body
	var user models.User
	c.logger.Info().Msg("parsing post request")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.logger.Info().Msgf("parsing unsuccessful: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Validate User
	c.logger.Info().Msg("validating registration")
	err = c.userservice.ValidateRegistration(&user)
	if err != nil {
		c.logger.Info().Msgf("validating registration failed: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Register User
	c.logger.Info().Msg("registering new user")
	_, err = c.userservice.Register(&user)
	if err != nil {
		c.logger.Info().Msgf("registration failed: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.logger.Info().Msg("registration successful")
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
	c.logger.Info().Msg("parsing post request")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.logger.Info().Msgf("parsing unsuccessful: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Validate User Login
	c.logger.Info().Msg("validating login unsuccessful")
	err = c.userservice.ValidateLogin(&user)
	if err != nil {
		c.logger.Info().Msgf("validating login failed: %v", err.Error())
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
	c.logger.Info().Msg("generating jwt token")
	token, err := c.userservice.GenerateJWT(&user, expiration)
	if err != nil {
		c.logger.Info().Msgf("generating jwt token failed: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Set cookie
	c.logger.Info().Msg("setting http-only cookie")
	http.SetCookie(w,
		&http.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  expiration,
			HttpOnly: true,
		})

	// Login
	c.logger.Info().Msg("login successful")
	c.userservice.Login(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Status: "Success",
		Code:   http.StatusOK,
	})
}
