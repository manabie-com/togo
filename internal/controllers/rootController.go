package controllers

import (
	"encoding/json"
	"net/http"

	config "github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/daos"
	"github.com/manabie-com/togo/internal/middleware"
	"github.com/manabie-com/togo/internal/models"
)

//LoginHandler will handle the login function
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var (
		// params   = mux.Vars(r)
		message = "OK"
	)
	//Get input from request
	var accountInfo = &models.Account{}
	if err := json.NewDecoder(r.Body).Decode(&accountInfo); err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	var err error
	accountDAO := daos.GetAccountDAO()
	result, err := accountDAO.FindAccountByUsernameAndPassword(*accountInfo)
	if err != nil {
		config.ResponseWithError(w, "Invalid username/password", err)
		return
	}
	token, _ := middleware.CreateAccountJWT(*result)
	resp := map[string]interface{}{
		"username": &result.Username,
		"token":    token,
	}
	config.ResponseWithSuccess(w, message, resp)
}
