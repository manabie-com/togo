package rest

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
)

// AuthHandler ...
type AuthHandler struct {
	authService services.AuthService
}

// NewAuthHandler ...
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register ...
func (ah *AuthHandler) Register(mux *http.ServeMux) {
	mux.Handle("/login", MethodMiddleware(LoggingMiddleWare(http.HandlerFunc(ah.Login)), http.MethodPost))
}

// Login handles /login route
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody loginRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		returnErrorJSONResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(reqBody.UserID) == 0 || len(reqBody.Password) == 0 {
		returnErrorJSONResponse(w, "Missing user_id or password", http.StatusBadRequest)
		return
	}

	token, err := ah.authService.AuthenticateUser(r.Context(), reqBody.UserID, reqBody.Password)
	if len(token) == 0 {
		if err != nil {
			returnErrorJSONResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		returnErrorJSONResponse(w, "Wrong username or password", http.StatusUnauthorized)
		return
	}
	returnSuccessJSONResponse(w, loginRequestResponse{Token: token}, http.StatusOK)
	return
}

type loginRequestBody struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type loginRequestResponse struct {
	Token string `json:"token"`
}
