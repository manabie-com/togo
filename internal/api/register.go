package api

import (
	"github.com/jmsemira/togo/internal/auth"
	"github.com/jmsemira/togo/internal/helper"
	"github.com/jmsemira/togo/internal/models"
	"net/http"
	"strconv"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}

	user := models.User{}

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	limit, _ := strconv.ParseInt(r.FormValue("rate_limit"), 10, 64)
	user.RateLimitPerDay = int(limit)
	response["status"] = "error"
	if user.Username == "" {
		response["err_msg"] = "Username is a required field"
		helper.ReturnJSON(w, response)
		return
	}

	if user.Password == "" {
		response["err_msg"] = "Password is a required field"
		helper.ReturnJSON(w, response)
		return
	}

	err := auth.Register(&user)
	if err != nil {
		response["status"] = "error"
		response["err_msg"] = err.Error()
		helper.ReturnJSON(w, response)
		return
	}

	response["status"] = "ok"

	helper.ReturnJSON(w, response)
}
