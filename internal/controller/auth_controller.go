package controller

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/response"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/services/impl"
	"net/http"
)

type AuthController struct {
	userService services.UsersService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{impl.NewUserServiceImpl(db) }
}

func(a AuthController) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	password := value(req, "password")
	resp.Header().Set("Content-Type", "application/json")
	data, err, code := a.userService.Login(id, password)
	if (err != nil){
		response.RespondWithError(resp, code, err.Error())
	} else {
		response.RespondWithJSON(resp, code, json.NewEncoder(resp).Encode(map[string]string{
			"data": data,
		}))
	}
}

func value(req *http.Request, p string) string {
	return req.FormValue(p)
}
