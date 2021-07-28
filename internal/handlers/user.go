package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/services/users"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
)

func getAuthToken(service users.ToDoService) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")

		u := &storages.User{}
		err := json.NewDecoder(req.Body).Decode(u)
		defer req.Body.Close()
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		username := utils.ConvertStringToSqlNullString(u.Username)
		password := utils.ConvertStringToSqlNullString(u.Password)
		token, err := service.GetAuthToken(username, password)
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

func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service users.ToDoService) {
	r.Handle("/login", n.With(
		negroni.Wrap(getAuthToken(service)),
	)).Methods("POST").Name("login")
}
