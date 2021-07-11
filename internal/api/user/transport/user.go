package transport

import (
	"encoding/json"
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/api/user/storages"
	userUseCase "github.com/manabie-com/togo/internal/api/user/usecase"
	"github.com/manabie-com/togo/internal/api/utils"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"github.com/manabie-com/togo/internal/pkg/token"
	"net/http"
)

type User struct {
	UserUC         userUseCase.User
	TokenGenerator token.Generator
}

func (s *User) Login(resp http.ResponseWriter, req *http.Request) {
	user := &storages.User{}

	err := json.NewDecoder(req.Body).Decode(user)
	defer req.Body.Close()

	ctx := req.Context()

	valErr := utils.ValidateRequest(user)
	if valErr != nil {
		utils.WriteJSON(ctx, resp, http.StatusBadRequest, nil, valErr)
		return
	}

	isValid, err := s.UserUC.IsValidate(ctx, user.ID, user.Password)
	if err != nil {
		utils.WriteJSON(ctx, resp, http.StatusUnauthorized, nil, err)
		return
	}

	if !isValid {
		utils.WriteJSON(ctx, resp, http.StatusUnauthorized, nil, errors.New(dictionary.IncorrectLogin))
		return
	}

	tokenStr, err := s.TokenGenerator.CreateToken(user.ID)
	if err != nil {
		logger.MBPrintf(ctx, "failed to create token: %s", err)
		utils.WriteJSON(ctx, resp, http.StatusInternalServerError, nil, errors.New(dictionary.FailedToCreateToken))
		return
	}
	utils.WriteJSON(ctx, resp, http.StatusOK, tokenStr, nil)
}
