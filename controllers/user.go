package controllers

import (
	"net/http"

	u "github.com/manabie-com/togo/utils"
)

var SignUp = func(w http.ResponseWriter, r *http.Request) {

	// send token jwt
	// ...
	u.Respond(w, http.StatusCreated, map[string]interface{}{})

}

var Login = func(w http.ResponseWriter, r *http.Request) {
	// get email password
	// if email and password not exist
	u.Respond(w, http.StatusBadRequest, map[string]interface{}{})
	// if email exist and password incorrect
	u.Respond(w, http.StatusUnauthorized, map[string]interface{}{})
	// if email and password OK
	// send token to client
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}

var GetMe = func(w http.ResponseWriter, r *http.Request) {
	// if jwt  invalid
	u.Respond(w, http.StatusBadRequest, map[string]interface{}{})
	// decode jwt -> get userID -> get user
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}

var UpdateMe = func(w http.ResponseWriter, r *http.Request) {
	// success
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}
