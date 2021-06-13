package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/ent"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const XErrorMessage = "X-Error-Message"

type ApiHandler interface {
	CreateTask() http.HandlerFunc
	GetTaskByDate() http.HandlerFunc
	Login() http.HandlerFunc
}

type apiHandlerImpl struct {
	jwtKey      string
	authService services.AuthService
	todoService services.TaskService
}

func NewApiHandler(jwtKey string, client *ent.Client) ApiHandler {
	return &apiHandlerImpl{
		jwtKey:      jwtKey,
		authService: services.NewAuthService(jwtKey, client),
		todoService: services.NewTaskService(client),
	}
}

func (h *apiHandlerImpl) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request model.TaskCreationRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeJsonRes(w, 400, errors.New("invalid request"))
			return
		}

		task, err := h.todoService.CreateTask(r.Context(), request)
		if err == nil {
			writeJsonRes(w, 201, task)
		} else {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("create task failed")

			writeJsonRes(w, 500, errors.New("internal server error"))
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
			writeJsonRes(w, 400, errors.New("invalid created_date"))
			return
		}

		tasks, err := h.todoService.GetTaskByDate(r.Context(), createdDate)
		if err == nil {
			writeJsonRes(w, 200, tasks)
		} else {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("retrieve task failed")

			writeJsonRes(w, 500, errors.New("internal server error"))
		}
	}
}

func (h *apiHandlerImpl) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credential model.LoginCredential
		err := json.NewDecoder(r.Body).Decode(&credential)
		if err != nil {
			writeJsonRes(w, 400, errors.New("invalid request"))
			return
		}

		token, err := h.authService.Login(r.Context(), credential)
		if err != nil {
			writeJsonRes(w, 401, errors.New("invalid username or password"))
			return
		} else {
			writeJsonRes(w, 200, token)
		}
	}
}

func writeJsonRes(w http.ResponseWriter, code int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if 200 <= code && code < 300 {
		w.WriteHeader(code)
		_, _ = w.Write(buf.Bytes())
	} else {
		msg := fmt.Sprintf("%v", v)
		w.Header().Set(XErrorMessage, msg)
		w.WriteHeader(code)
		_, _ = w.Write(buf.Bytes())
	}
}
