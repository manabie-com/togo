package transport

import (
	"encoding/json"
	"net/http"

	middleware "github.com/manabie-com/togo/internal/middlewares"
	errutil "github.com/manabie-com/togo/internal/pkg/error_utils"
	"github.com/manabie-com/togo/internal/storages"
)

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func renderSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, err := json.Marshal(data)
	if err != nil {
		renderBadRequest(w, err.Error())
		return
	}

	w.Write(resp)
}

func renderUnauthorizedError(w http.ResponseWriter, message string) {
	err := errutil.NewUnauthorizedError(message)

	renderError(w, err)
}

func renderUnprocessibleEntityError(w http.ResponseWriter, message string) {
	err := errutil.NewUnprocessibleEntityError(message)

	renderError(w, err)
}

func renderBadRequest(w http.ResponseWriter, message string) {
	err := errutil.NewBadRequestError(message)

	renderError(w, err)
}

func renderError(w http.ResponseWriter, err errutil.MessageErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status())

	resp, _ := json.Marshal(err)

	w.Write(resp)
}

func populateLoginResponse(user storages.User) (loginResponse, error) {
	claim := middleware.JWTClaims{
		UserID: user.ID,
	}

	token, err := middleware.GenerateJWT(claim)
	if err != nil {
		return loginResponse{}, err
	}

	return loginResponse{
		AccessToken:  token["access_token"],
		RefreshToken: token["refresh_token"],
	}, nil
}
