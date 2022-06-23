package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}

	validate := validator.New()

	if err = validate.Struct(user); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	// insert database
	err = db.QueryRow(`INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id, name, email`, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	fmt.Println(user)
	// send token jwt here
	tk := &models.Token{UserId: user.ID}
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
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	var (
		email    string
		password string
	)
	err = db.QueryRow(`SELECT id, name, email, password FROM users WHERE email = $1`, user.Email).Scan(&user.ID, &user.Name, &email, &password)
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
	fmt.Println(user)
	tk := &models.Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	fmt.Println(os.Getenv("SECRET_TOKEN"))
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	// response and send token to client
	u.Respond(w, http.StatusOK, "Success", "Login Success", map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"token": tokenString,
	})
}

var GetMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var userID = r.Context().Value("user").(uint32)

	var user = &models.User{
		ID: userID,
	}

	err := db.QueryRow(`SELECT name, email, is_payment, limit_day_tasks FROM users WHERE id = $1`, user.ID).Scan(&user.Name, &user.Email, &user.IsPayment, &user.LimitDayTasks)

	if err != nil {
		u.Respond(w, http.StatusNotFound, "Failure", err.Error(), nil)
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
	user := models.User{}
	// success
	u.Respond(w, http.StatusOK, "Success", "Success", map[string]interface{}{
		"name":            user.Name,
		"email":           user.Email,
		"is_payment":      user.IsPayment,
		"limit_day_tasks": user.LimitDayTasks,
	})
}
