package router

import (
	"encoding/json"
	"net/http"
	"todo/be/db"
	"todo/be/env"
	"todo/be/utils"

	"github.com/gorilla/mux"
)

const (
	http_POST       = "POST"
	key_TokenAes    = "kdfjka475ushs7483y9e8f8f84ihg834"
	key_TokenHeader = "Token"
)

func InitRouter(router *mux.Router) {
	initRouterApi(router)
	// init another router
}

func initUserToken() db.User {
	return db.User{
		UserId: utils.RandomString(env.LENGTH_USER_ID),
		Limit:  utils.RandomNumber(env.MIN_TASK, env.MAX_TASK),
	}
}

func getUserData(userData string) (db.User, bool) {
	if len(userData) == 0 {
		return initUserToken(), true
	}
	userJson, ok := utils.DecryptAES(key_TokenAes, userData)
	if !ok {
		return initUserToken(), true
	}
	var user db.User
	err := json.Unmarshal(userJson, &user)
	if utils.IsError(err) {
		return initUserToken(), true
	}
	return user, false
}

func getUserToken(response http.ResponseWriter, request *http.Request) db.User {
	userData := request.Header.Get(key_TokenHeader)
	user, isNew := getUserData(userData)
	if isNew {
		bytes, _ := json.Marshal(user)
		token, _ := utils.EncryptAES(key_TokenAes, bytes)
		response.Header().Set(key_TokenHeader, token)
	}
	return user
}
