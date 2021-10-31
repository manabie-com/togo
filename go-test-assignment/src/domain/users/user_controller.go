package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UserController ...
type UserController struct {
	UserCreateAction UserCreateAction
}

// name...
func (ctrl UserController) name() string {
	return "User.UserController"
}

// Create ...
func (ctrl UserController) Create(w http.ResponseWriter, r *http.Request) {
	payload := &UserCreatePayload{}
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

	UserDetail, err := ctrl.UserCreateAction.Execute(payload.Username, payload.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("While a creating user error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserDetail)
}
