package services

import (
	"net/http"

	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/storages/postgres"
	utils "github.com/manabie-com/togo/internal/utils"
)

type loginRequest struct {
	UserId   string `json:"user_id" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,alphanum"`
}

type loginResponse struct {
	Data string `json:"data"`
}

func (sc *ServiceController) loginHandler(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")

	// valid := sc.validateUser(req.Context(), id, value(req, "password"))
	valid := sc.Store.ValidateUser(req.Context(), postgres.ValidateUserParams{
		ID:       id,
		Password: value(req, "password"),
	})

	if !valid {
		// means that not valid userid/password so response invalid
		common.CommonResponse(resp, http.StatusUnauthorized, map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}

	token, err := utils.CreateToken(id.String, sc.Config.JWTKey)
	if err != nil {
		common.CommonResponse(resp, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	common.CommonResponse(resp, http.StatusOK, map[string]string{
		"data": token,
	})
}

// func (sc *ServiceController) validateUser(ctx context.Context, userId, pwd sql.NullString) bool {
// 	return sc.Store.ValidateUser(ctx, postgres.ValidateUserParams{
// 		ID:       userId,
// 		Password: pwd,
// 	})
// }
