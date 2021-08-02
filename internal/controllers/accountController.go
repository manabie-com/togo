package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/daos"
	"github.com/manabie-com/togo/internal/middleware"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/utils"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var (
		message = "OK"
	)

	accountDAO := daos.GetAccountDAO()
	//validations
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	var account = models.Account{}
	if err := json.Unmarshal(requestBody, &account); err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	isUsernameValid, err := utils.IsUsernameValid(account.Username)
	if err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	if !isUsernameValid {
		config.ResponseWithError(w, "Invalid username", err)
		return
	}
	isPasswordValid, err := utils.IsPasswordValid(account.Password)
	if err != nil {
		config.ResponseWithError(w, "Malformed data", err)
		return
	}
	if !isPasswordValid {
		config.ResponseWithError(w, "Invalid username", err)
		return
	}
	//create account on db
	result, err := accountDAO.CreateAccount(models.Account{
		Username: account.Username,
		Password: account.Password,
	})
	if err != nil {
		config.ResponseWithError(w, "Create account failed", err)
		return
	}
	token, _ := middleware.CreateAccountJWT(*result)
	resp := map[string]interface{}{
		"username": &result.Username,
		"token":    token,
	}
	config.ResponseWithSuccess(w, message, resp)
}

func ViewProfileHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var (
		// params  = mux.Vars(r)
		message = "OK"
	)
	var accountID uuid.UUID
	var err error

	//If the user is looking for another profile
	if len(r.URL.Query()["account_id"]) > 0 {
		accountID, err = uuid.Parse(r.URL.Query()["account_id"][0])
		if err != nil {
			config.ResponseWithError(w, "malformed uuid", err)
		}
	} else {
		ctx := r.Context()
		accountID, _ = uuid.Parse(fmt.Sprint(ctx.Value("account_id")))
	}

	accountDAO := daos.GetAccountDAO()
	userDetails, err := accountDAO.FindAccountByID(accountID)
	if err != nil {
		config.ResponseWithError(w, "view account failed", err)
		return
	}
	config.ResponseWithSuccess(w, message, userDetails)

}
