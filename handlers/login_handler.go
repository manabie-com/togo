package handlers

import (
	"encoding/json"
	"errors"
	"github.com/manabie-com/togo/pkg/core/servehttp"
	"github.com/manabie-com/togo/pkg/utils"
	"github.com/manabie-com/togo/usecases"
	"log"
	"net/http"
)

type LoginHandler struct {
	Uc usecases.LoginUseCase
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	input, err := getLoginRequest(r)
	if err != nil {
		log.Printf("Validation error, detail: %v", err.Error())
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	userToken, err := h.Uc.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Verify login error, detail: %v", err.Error())
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	servehttp.ResponseSuccessJSON(w, map[string]interface{}{
		"token": userToken,
	})
	return
}

func getLoginRequest(r *http.Request) (*usecases.LoginInput, error) {
	defer r.Body.Close()
	var input *usecases.LoginInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, errors.New("cannot parse login request")
	}

	validator, _ := utils.NewGoPlayground()

	err := validator.Validate(input)
	if err != nil {
		return nil, err
	}

	if len(validator.Messages()) > 0 {
		return nil, errors.New(validator.Messages()[0])
	}

	return input, nil
}
