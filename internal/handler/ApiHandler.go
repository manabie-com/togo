package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const errorHeader = "X-Error-Message"

type ApiHandler interface {
	CreateTask() http.HandlerFunc
	GetTaskByDate() http.HandlerFunc
	Login() http.HandlerFunc
	AuthenticationMiddleware(next http.Handler) http.Handler
}

type apiHandlerImpl struct {
	JWTKey string
}

func NewApiHandler(jwtKey string) ApiHandler {
	return &apiHandlerImpl{JWTKey: jwtKey}
}

func (h *apiHandlerImpl) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *apiHandlerImpl) GetTaskByDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *apiHandlerImpl) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *apiHandlerImpl) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]

		claims := make(jwt.MapClaims)
		t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(h.JWTKey), nil
		})

		if err == nil && t.Valid {
			userId, ok := claims["user_id"].(string)
			if ok {
				ctx := context.WithValue(r.Context(), "userId", userId)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		log.WithFields(log.Fields{
			"error": err,
			"token": token,
		}).Info("Invalid token")

		w.Header().Set(errorHeader, "Invalid token")
		w.WriteHeader(401)
	})
}

func (h *apiHandlerImpl) writeJsonRes(w http.ResponseWriter, code int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(buf.Bytes())
}
