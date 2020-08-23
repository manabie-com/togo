package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	userRepository "github.com/manabie-com/togo/internal/repository"
)

// AuthHandler handles authentication
type AuthHandler struct {
	JWTKey string
	Repo   *userRepository.UserRepository
}

//Login func handles user login
func (handler *AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	var res map[string]interface{}

	json.NewDecoder(req.Body).Decode(&res)

	username := res["username"]

	password := res["password"]

	user, err := handler.Repo.ValidateUser(username.(string), password.(string))

	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}

	token, err := handler.createToken(user.ID)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]map[string]string{
		"data": {
			"token":    token,
			"id":       user.ID,
			"userName": user.Username,
		},
	})
}

func (handler *AuthHandler) createToken(ID string) (string, error) {
	atClaims := jwt.MapClaims{}

	atClaims["userId"] = ID

	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(handler.JWTKey))

	if err != nil {
		return "", err
	}
	return token, nil
}
