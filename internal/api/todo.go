package api

import (
	"github.com/jmsemira/togo/internal/auth"
	"github.com/jmsemira/togo/internal/helper"
	"github.com/jmsemira/togo/internal/models"
	"net/http"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}
	todo := models.Todo{}

	user, _ := r.Context().Value("user").(*auth.Claims)
	todo.UserID = user.ID

	if r.FormValue("name") == "" {
		response["status"] = "error"
		response["err_msg"] = "Name is required!"
		helper.ReturnJSON(w, response)
		return
	}
	todo.Name = r.FormValue("name")

	err := todo.Save(user.RateLimit)
	if err != nil {
		response["status"] = "error"
		response["err_msg"] = err.Error()
		helper.ReturnJSON(w, response)
		return
	}

	response["status"] = "OK"
	helper.ReturnJSON(w, response)
}
