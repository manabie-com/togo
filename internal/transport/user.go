package transport

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/utils"
	"net/http"
	"time"
)

type UserTransport interface{
	Login( resp http.ResponseWriter, req *http.Request)
}
type userTransport struct {
	userService services.UserService
}

func NewUserTransport(us services.UserService) UserTransport{
	return &userTransport{
		userService: us,
	}
}

func (ut *userTransport) Login( resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userId := req.FormValue("user_id")
	password := req.FormValue("password")
	isValidUser, err := ut.userService.GetAuthToken(ctx, userId, password)

	if !isValidUser {
		utils.HttpResponseUnauthorized(resp, err.Error())
		return
	}

	token, err := createToken(userId)
	if err != nil {
		utils.HttpResponseInternalServerError(resp, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

func createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("wqGyEBBfPK9w3Lxw"))
	if err != nil {
		return "", err
	}
	return token, nil
}

