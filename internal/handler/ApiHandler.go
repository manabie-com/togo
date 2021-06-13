package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/ent"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const ErrorHeader = "X-Error-Message"

type ApiHandler interface {
	CreateTask() http.HandlerFunc
	GetTaskByDate() http.HandlerFunc
	Login() http.HandlerFunc
	AuthenticationMiddleware(next http.Handler) http.Handler
}

type apiHandlerImpl struct {
	jwtKey      string
	authService services.AuthService
	todoService services.ToDoService
}

func NewApiHandler(jwtKey string, client *ent.Client) ApiHandler {
	return &apiHandlerImpl{
		jwtKey:      jwtKey,
		authService: services.NewAuthService(jwtKey, client),
		todoService: services.NewToDoService(client),
	}
}

func (h *apiHandlerImpl) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request model.TaskCreationRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			h.writeJsonRes(w, 400, errors.New("invalid request"))
			return
		}

		task, err := h.todoService.CreateTask(r.Context(), request)
		if err == nil {
			h.writeJsonRes(w, 201, task)
		} else {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("create task failed")

			h.writeJsonRes(w, 500, errors.New("internal server error"))
		}

	}
}

func (h *apiHandlerImpl) GetTaskByDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.Query().Get("created_date")
		createdDate, err := time.Parse("2006-01-02", str)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("invalid created_date")
			h.writeJsonRes(w, 400, errors.New("invalid created_date"))
			return
		}

		tasks, err := h.todoService.GetTaskByDate(r.Context(), createdDate)
		if err == nil {
			h.writeJsonRes(w, 200, tasks)
		} else {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("retrieve task failed")

			h.writeJsonRes(w, 500, errors.New("internal server error"))
		}
	}
}

func (h *apiHandlerImpl) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credential model.LoginCredential
		err := json.NewDecoder(r.Body).Decode(&credential)
		if err != nil {
			h.writeJsonRes(w, 400, errors.New("invalid request"))
			return
		}

		token, err := h.authService.Login(r.Context(), credential)
		if err != nil {
			h.writeJsonRes(w, 401, errors.New("invalid username or password"))
			return
		} else {
			h.writeJsonRes(w, 200, token)
		}
	}
}

func (h *apiHandlerImpl) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		var err error
		if len(splitToken) == 2 {
			token = splitToken[1]

			claims := make(jwt.MapClaims)
			var t *jwt.Token
			t, err = jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
				return []byte(h.jwtKey), nil
			})

			if err == nil && t.Valid {
				userId, ok := claims["user_id"].(string)
				if ok {
					ctx := context.WithValue(r.Context(), "userId", userId)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
		} else {
			err = errors.New("invalid bearer")
		}

		log.WithFields(log.Fields{
			"error": err,
			"token": token,
		}).Info("Invalid token")

		h.writeJsonRes(w, 401, errors.New("invalid token"))
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

	if code <= 200 && code < 300 {
		w.WriteHeader(code)
		_, _ = w.Write(buf.Bytes())
	} else {
		msg := fmt.Sprintf("%v", v)
		w.Header().Set(ErrorHeader, msg)
		w.WriteHeader(code)
		_, _ = w.Write(buf.Bytes())
	}
}
