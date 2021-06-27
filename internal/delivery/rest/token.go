package rest

import (
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/utils"
	"net/http"
)

const (
	userIDKey   = "user_id"
	passwordKey = "password"
)

func (s *Serializer) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	userID := utils.Value(req, userIDKey)
	password := utils.Value(req, passwordKey)

	token, err := s.TodoService.GetAuthToken(userID, password)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(
			utils.BuildErrorResponseRequest(&utils.Meta{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		return
	}

	responseData := map[string]string{
		"token": token,
	}

	json.NewEncoder(resp).Encode(
		utils.BuildSuccessResponseRequest(&utils.Meta{
			Code:    http.StatusCreated,
			Message: utils.SuccessRequestMessage,
		}, responseData))
}

func (s *Serializer) ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	userID, valid := s.TodoService.ValidToken(token)
	if !valid {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), utils.UserAuthKey(0), userID))
	return req, true
}
