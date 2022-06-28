package controllers

import (
	"database/sql"
	"encoding/json"
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
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input format: "+err.Error())
		return
	}
	// validate user object
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input field: "+err.Error())
		return
	}
	// insert database
	err = user.InsertUser(db)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Your email is duplicated, Please try again")
		return
	}
	// send token jwt here
	tk := &models.Token{UserId: user.ID, LimitDayTasks: user.LimitDayTasks}
	tokenString := tk.CreateToken()
	// everything Ok
	u.SuccessRespond(w, http.StatusCreated, "Created Account", map[string]interface{}{
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
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input format: "+err.Error())
		return
	}

	var (
		email    string = user.Email
		password string = user.Password
	)
	// get user by email
	err = user.GetUserByEmail(db)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Not found account with your email, Please provide a valid email address!")
		return
	}
	// if email exist and password incorrect
	if email != user.Email || password != user.Password {
		u.FailureRespond(w, http.StatusUnauthorized, "Password incorrect")
		return
	}
	// create message case 1: Active user, case 2:  unActive user
	var message = "Login Success"
	if !user.IsActive {
		message = "Welcome back"
		err = user.ActiveUser(db)
		if err != nil {
			u.FailureRespond(w, http.StatusInternalServerError, "Something went wrong"+err.Error())
			return
		}
	}
	// if email and password OK
	//Create JWT token
	tk := &models.Token{
		UserId:        user.ID,
		LimitDayTasks: user.LimitDayTasks,
	}
	tokenString := tk.CreateToken()
	// everything Ok
	u.SuccessRespond(w, http.StatusOK, message, map[string]interface{}{
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
	err := user.GetUserById(db)
	if err != nil {
		u.FailureRespond(w, http.StatusInternalServerError, "Something went wrong when collect your account. Please try again"+err.Error())
		return
	}
	// everything Ok
	u.SuccessRespond(w, http.StatusOK, "Success", map[string]interface{}{
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
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input format: "+err.Error())
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
	if err := validate.Var(email, "email,min=10,max=30"); email != "" && err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input email: "+err.Error())
		return
	}

	if err := validate.Var(name, "min=5,max=20"); name != "" && err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input name: "+err.Error())
		return
	}
	err = user.GetUserById(db)
	if err != nil {
		u.FailureRespond(w, http.StatusInternalServerError, "Somethings went wrong. Please try again"+err.Error())
		return
	}
	// confirm password
	if password != user.Password {
		u.FailureRespond(w, http.StatusUnauthorized, "Password incorrect. Please try again")
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
	err = user.UpdateUser(db)
	if err != nil {
		u.FailureRespond(w, http.StatusInternalServerError, "Something went wrong when update your account. Please try again"+err.Error())
		return
	}
	// everything Ok
	u.SuccessRespond(w, http.StatusOK, "Success update your account!", map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
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
		u.FailureRespond(w, http.StatusBadRequest, "Invalid input format: "+err.Error())
		return
	}
	// get user info
	inputPassword := user.Password
	err = user.GetUserById(db)
	if err != nil {
		u.FailureRespond(w, http.StatusInternalServerError, "Something went wrong when collect your account. Please try again"+err.Error())
		return
	}
	// if input password not equal to database password
	if inputPassword != user.Password {
		u.FailureRespond(w, http.StatusUnauthorized, "Password incorrect. Please try again")
		return
	}
	// update field is_active to false
	err = user.DeleteUser(db)
	if err != nil {
		u.FailureRespond(w, http.StatusInternalServerError, "Something went wrong when delete your account. Please try again"+err.Error())
		return
	}
	// everything Ok
	u.SuccessRespond(w, http.StatusNoContent, "Success delete your account!", nil)
}
