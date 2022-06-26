package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var SignUp = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		IsPayment:     false,
		LimitDayTasks: 10,
	}
	// decode json body to user
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input format: "+err.Error(), nil)
		return
	}
	// validate user object
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input field: "+err.Error(), nil)
		return
	}
	// insert database
	if err := user.InsertOne(db); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Your email is duplicated, Please try again", nil)
		return
	}
	// send token jwt here
	tk := &models.Token{UserId: user.ID, LimitDayTasks: user.LimitDayTasks}
	tokenString := tk.CreateToken()
	// everything Ok
	u.Respond(w, http.StatusCreated, "Success", "Created Account", map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"token": tokenString,
	})
}

var Login = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	// validate
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input format: "+err.Error(), nil)
		return
	}
	fmt.Println(user)
	var (
		email    string = user.Email
		password string = user.Password
	)
	// get user by email
	if err := user.GetOneByEmail(db); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Not found account with your email, Please provide a valid email address!", nil)
		return
	}
	// if email exist and password incorrect
	if email != user.Email || password != user.Password {
		u.Respond(w, http.StatusUnauthorized, "Failure", "Password incorrect", nil)
		return
	}
	// create message case 1: Active user, case 2:  unActive user
	var message = "Login Success"
	if !user.IsActive {
		message = "Welcome back"
		_, _ = db.Exec(`UPDATE users SET is_active = $1 WHERE id = $2`, true, user.ID)
	}
	// if email and password OK
	//Create JWT token
	tk := &models.Token{
		UserId:        user.ID,
		LimitDayTasks: user.LimitDayTasks,
	}
	tokenString := tk.CreateToken()
	// everything Ok
	u.Respond(w, http.StatusOK, "Success", message, map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"token": tokenString,
	})
}

var GetMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get decoded token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	user := &models.User{
		ID: decoded.UserId,
	}
	// query
	if err := user.GetOneById(db); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when collect your account. Please try again", nil)
		return
	}
	// everything Ok
	u.Respond(w, http.StatusOK, "Success", "Success", map[string]interface{}{
		"name":            user.Name,
		"email":           user.Email,
		"is_payment":      user.IsPayment,
		"limit_day_tasks": user.LimitDayTasks,
	})
}

var UpdateMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get decoded token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	user := &models.User{
		ID: decoded.UserId,
	}
	// convert json -> user object
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input format: "+err.Error(), nil)
		return
	}
	// get user info
	var (
		name     string = user.Name
		email    string = user.Email
		password string = user.Password
	)
	// validate input
	validate := validator.New()
	if err := validate.Var(email, "email"); email != "" && err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input email: "+err.Error(), nil)
		return
	}

	if err := validate.Var(name, "min=5,max=20"); name != "" && err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input name: "+err.Error(), nil)
		return
	}

	if err := user.GetOneById(db); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when collect your account. Please try again", nil)
		return
	}
	// confirm password
	if password != user.Password {
		u.Respond(w, http.StatusUnauthorized, "Failure", "Password incorrect. Please try again", nil)
		return
	}
	// if valid value => overwrite new value
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	// update me
	_, err = db.Exec(`UPDATE users SET name = $1, email = $2 WHERE id = $3`, user.Name, user.Email, user.ID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when update your account. Please try again", nil)
		return
	}
	// everything Ok
	u.Respond(w, http.StatusOK, "Success", "Success update your account!", map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"is_payment": user.IsPayment,
	})
}

var DeleteMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get decoded token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	user := &models.User{
		ID:       decoded.UserId,
		IsActive: false,
	}
	// convert json -> user object
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input format: "+err.Error(), nil)
		return
	}
	// get user info
	inputPassword := user.Password
	if err := user.GetOneById(db); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when collect your account. Please try again", nil)
		return
	}
	// if input password not equal to database password
	if inputPassword != user.Password {
		u.Respond(w, http.StatusUnauthorized, "Failure", "Password incorrect. Please try again", nil)
		return
	}
	// update field is_active to false
	_, err = db.Exec(`UPDATE users SET is_active = $1 WHERE id = $2`, user.IsActive, user.ID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when delete your account. Please try again", nil)
		return
	}
	// everything Ok
	u.Respond(w, http.StatusOK, "Success", "Success delete your account!", nil)
}
