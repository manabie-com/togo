package rest

import (
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/utils"
	"github.com/manabie-com/togo/internal/utils/constants"
	"go.uber.org/zap"
	"net/http"
)

const (

)

func (s *Serializer) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	userID := utils.Value(req, constants.UserIDKey)
	password := utils.Value(req, constants.PasswordKey)

	token, err := s.TodoService.GetAuthToken(userID, password)
	if err != nil {
		zap.L().Error("get token Error.", zap.Error(err))
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

func (s *Serializer) ValidToken(resp http.ResponseWriter, req *http.Request) (*http.Request, bool) {
	token := req.Header.Get(constants.AuthorizationHeader)
	userID, valid := s.TodoService.ValidToken(token)
	if !valid {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(
			utils.BuildErrorResponseRequest(&utils.Meta{
				Code:    http.StatusUnauthorized,
				Message: utils.UnauthorizedRequestMessage,
			}))
		return nil, false
	}

	req = req.WithContext(context.WithValue(req.Context(), utils.UserAuthKey(0), userID))
	return req, true
}
