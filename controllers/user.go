package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var SignUp = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Had field invalid"))
	}
	defer db.Close()
	fmt.Println(user)
	var userId int
	err = db.QueryRow(`INSERT INTO users(name, email, password) VALUES($1,$2,$3) RETURNING id;`, user.Name, user.Email, user.Password).Scan(&userId)
	fmt.Println(userId)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Invalid request"))
	}
	// send token jwt here
	// ...
	u.Respond(w, http.StatusCreated, u.Message(true, "Created Account"))
}

var Login = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	// get email password
	// if email and password not exist
	u.Respond(w, http.StatusBadRequest, map[string]interface{}{})
	// if email exist and password incorrect
	u.Respond(w, http.StatusUnauthorized, map[string]interface{}{})
	// if email and password OK
	// send token to client
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}

var GetMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// if jwt  invalid
	u.Respond(w, http.StatusBadRequest, map[string]interface{}{})
	// decode jwt -> get userID -> get user
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}

var UpdateMe = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// success
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}
