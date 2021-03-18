package services

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

func (s *ToDoService) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	parId := req.FormValue("user_id")
	parPassword := req.FormValue("password")
	if !s.Store.ValidateUser(Value(parId), Value(parPassword)) {
		writeHeader(&resp, http.StatusUnauthorized)
		json := make(map[string]string)
		json["error"] = "incorrect user_id/pwd"
		responseJson(&resp, json)

		return
	}

	resp.Header().Set("Content-Type", "application/json")
	token, err := s.createToken(parId)
	if err != nil {
		writeHeader(&resp, http.StatusInternalServerError)
		json := make(map[string]string)
		json["error"] = err.Error()
		responseJson(&resp, json)

		return
	}

	json := make(map[string]string)
	json["data"] = token
	responseJson(&resp, json)
}

func (s *ToDoService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

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

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}
