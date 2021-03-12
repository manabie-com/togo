package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// HandlerFunc as proxy to embed middleware
type HandlerFunc func(http.ResponseWriter, *http.Request)

// ContextKey ...
type ContextKey string

// ContextUserKey ...
const ContextUserKey ContextKey = "user_id"

// normalizeAuthorizationHeader normalize Authorization header
func normalizeAuthorizationHeader(authHeader string) (string, error) {
	authHeader = strings.Trim(authHeader, " ")
	if len(authHeader) == 0 {
		return "", errors.New("Unauthorized")
	}

	// Support both 'Authorization: Bearer token' and 'Authorization: token'
	authHeaderParts := strings.Split(authHeader, " ")
	if authHeaderParts[0] == "Bearer" {
		if len(authHeaderParts) > 1 {
			authHeader = authHeaderParts[1]
		} else {
			return "", errors.New("Invalid header format")
		}
	}
	return authHeader, nil
}

// ErrorJSONResponse represent a reponse contain an error message
type ErrorJSONResponse struct {
	Error string `json:"error"`
}

// SuccessJSONResponse represent a reponse contain a success message
type SuccessJSONResponse struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data"`
}

func returnErrorJSONResponse(w http.ResponseWriter, msg string, status int) {
	returnJSON(w, ErrorJSONResponse{Error: msg}, status)
}

func returnSuccessJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	returnJSON(w, SuccessJSONResponse{Data: data}, status)
}

func returnJSON(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(content)
}
