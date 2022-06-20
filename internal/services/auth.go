package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/manabie-com/togo/internal/storages/entities"
)

type userAuthKey int8

var (
	userClaimKey = "user_id"
)

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	var user entities.User
	if err := json.Unmarshal(body, &user); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if !s.Store.ValidateUser(req.Context(), user.ID, user.Password) {
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

func (s *ToDoService) createToken(userID string) (string, error) {
	atClaims := jwt.MapClaims{
		userClaimKey: userID,
		"exp":        time.Now().Add(time.Hour * 180).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")
	// we should remove bearer from token
	token = strings.Replace(token, "bearer ", "", 1)
	token = strings.Replace(token, "Bearer ", "", 1)

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims[userClaimKey].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}
