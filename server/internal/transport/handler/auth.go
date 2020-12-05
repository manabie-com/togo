package handler

import (
	"encoding/json"
	"github.com/HoangVyDuong/togo/internal/usecase/auth"
	"net/http"
)

func AuthAccount(authService auth.Service, jwtKey string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("user_id")
		if authService.Auth(r.Context(), id) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "incorrect user_id/pwd",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")

		token, err := createToken(id, jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"data": token,
		})
	}
}
