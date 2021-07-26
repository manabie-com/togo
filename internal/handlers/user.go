package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/services/user"
	"github.com/manabie-com/togo/internal/utils"
)

func getAuthToken(service user.ToDoService) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")

		username := utils.Value(req, "user_name")
		token, err := service.GetAuthToken(username, utils.Value(req, "password"))
		if err != nil {
			resp.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": "incorrect user_id/pwd",
			})
			return
		}

		json.NewEncoder(resp).Encode(map[string]string{
			"data": token,
		})
	})
}

func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.ToDoService) {
	r.Handle("/login", n.With(
		negroni.Wrap(getAuthToken(service)),
	)).Methods("GET", "OPTIONS").Name("login")
}
