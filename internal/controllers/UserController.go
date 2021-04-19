package controllers

import (
	"encoding/json"
	"database/sql"
	"log"
	"net/http"
	"github.com/manabie-com/togo/internal/services"
)	

type UserController struct {
	services.IUserService
}

func (controller *UserController) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)

	id, password := req.FormValue("user_id"), req.FormValue("password")

	// log.Println(id, password)
	if !controller.ValidateUser(id, password ) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := controller.CreateToken(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})	

}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

