package services

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"
	"time"
)

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	auth := storages.GetLogin()
	user := storages.GetUser()
	err := json.NewDecoder(req.Body).Decode(&auth)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.Store.Where("id = ? and password = ?",  auth.UserID, auth.Password).Find(&user).Error
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(user.ID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

func (s *ToDoService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
