package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var SignUp = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		IsPayment:     false,
		LimitDayTasks: 10,
	}
	// validate
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input format: "+err.Error(), nil)
		return
	}

	validate := validator.New()

	if err = validate.Struct(user); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input field: "+err.Error(), nil)
		return
	}
	// insert database
	err = db.QueryRow(`INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id, name, email`, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		u.Respond(w, http.StatusServiceUnavailable, "Failure", "OOPS, Something went wrong, please try again later", nil)
		return
	}

	// send token jwt here
	tk := &models.Token{UserId: user.ID, LimitDayTasks: user.LimitDayTasks}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	// response and send token to client
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
	var (
		email    string
		password string
	)
	err = db.QueryRow(`SELECT id, name, email, password, limit_day_tasks FROM users WHERE email = $1`, user.Email).Scan(&user.ID, &user.Name, &email, &password, &user.LimitDayTasks)
	if err != nil {
		u.Respond(w, http.StatusNotFound, "Failure", "Your email invalid", nil)
		return
	}
	// if email exist and password incorrect
	if email != user.Email || password != user.Password {
		u.Respond(w, http.StatusUnauthorized, "Failure", "Password incorrect", nil)
		return
	}
	// if email and password OK
	//Create JWT token
	tk := &models.Token{
		UserId:        user.ID,
		LimitDayTasks: user.LimitDayTasks,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	// response and send token to client
	u.Respond(w, http.StatusOK, "Success", "Login Success", map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"token": tokenString,
	})
}

var GetMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get decoded token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	user := &models.User{}
	// query
	err := db.QueryRow(`SELECT name, email, is_payment, limit_day_tasks FROM users WHERE id = $1`, decoded.UserId).Scan(&user.Name, &user.Email, &user.IsPayment, &user.LimitDayTasks)

	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when collect your account. Please try again", nil)
		return
	}
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
	user := &models.User{}
	// convert json -> user object
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Invalid input format: "+err.Error(), nil)
		return
	}
	// get user info
	err = db.QueryRow(`SELECT name, email FROM users WHERE user_id = $1`, decoded.UserId).Scan(&user.Name, &user.Email)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when collect your account. Please try again", nil)
		return
	}
	// update me
	err = db.QueryRow(`UPDATE users, email SET name = $1, email = $2 WHERE id = $3 RETURNING name, email`, user.Name, user.Email, decoded.UserId).Scan(&user.Name, &user.Email)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong when update your account.. Please try again", nil)
		return
	}

	u.Respond(w, http.StatusOK, "Success", "Success update your account!", map[string]interface{}{
		"name":            user.Name,
		"email":           user.Email,
		"is_payment":      user.IsPayment,
		"limit_day_tasks": user.LimitDayTasks,
	})
}
