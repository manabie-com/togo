package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthController struct {
	AuthLoginAction  AuthLoginAction
	AuthLogoutAction AuthLogoutAction
}

// name...
func (ctrl AuthController) name() string {
	return "auth.AuthController"
}

// Login ...
func (ctrl AuthController) Login(w http.ResponseWriter, r *http.Request) {
	payload := &AuthLoginPayload{}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid payload")
		return
	}

	if validErrs := payload.Validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validErrs)
		return
	}

	accessToken, err := ctrl.AuthLoginAction.Execute(payload.Username, payload.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Logout ...
func (ctrl AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	err := ctrl.AuthLogoutAction.Execute()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(fmt.Sprintf("Unauthorized : %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
