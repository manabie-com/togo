package api

import (
	"fmt"
	"github.com/jmsemira/togo/internal/auth"
	"github.com/jmsemira/togo/internal/helper"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := auth.Login(username, password)
	if err != nil {
		fmt.Errorf(err.Error())
		response["status"] = "error"
		response["err_msg"] = err.Error()
		helper.ReturnJSON(w, response)
		return
	}

	token, err := auth.GenerateJWTToken(user)

	if err != nil {
		fmt.Errorf(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		helper.ReturnJSON(w, response)
		return
	}

	response["status"] = "OK"
	response["token"] = token
	helper.ReturnJSON(w, response)
}
