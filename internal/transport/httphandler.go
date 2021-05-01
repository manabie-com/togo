package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/services"
)

// HttpHandler implements http.Handler
type HttpHandler struct {
	ToDoService services.ToDoService
	jwtKey      string
}

// NewHttpHandler returns a HttpHandler
func NewHttpHandler(jwtKey string, service services.ToDoService) *HttpHandler {
	return &HttpHandler{
		jwtKey:      jwtKey,
		ToDoService: service,
	}
}

// ServeHTTP implements http.Handler.ServeHTTP(ResponseWriter, *Request)
func (h *HttpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	switch req.URL.Path {
	case "/login":
		h.getAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = h.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			h.listTasks(resp, req)
		case http.MethodPost:
			h.addTask(resp, req)
		}
		return
	}
}

func (h *HttpHandler) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !h.ToDoService.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := h.createToken(id)
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

func (h *HttpHandler) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := h.ToDoService.ListTasks(req.Context(), id, value(req, "created_date"))
	if err != nil {
		responseError(err, http.StatusInternalServerError, resp)
	}
	responseOk(resp, tasks)
}

func (h *HttpHandler) addTask(resp http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responseError(err, http.StatusInternalServerError, resp)
	}
	userID, _ := userIDFromCtx(req.Context())
	task, err := h.ToDoService.AddTask(req.Context(), userID, reqBody)
	if err != nil {
		if err == services.ErrUserReachDailyRequestLimit {
			responseError(err, http.StatusForbidden, resp)
			return
		}
		responseError(err, http.StatusInternalServerError, resp)
		return
	}
	responseOk(resp, task)
}

func responseError(err error, statusCode int, resp http.ResponseWriter) {
	resp.WriteHeader(statusCode)
	json.NewEncoder(resp).Encode(map[string]string{
		"error": err.Error(),
	})
}

func responseOk(resp http.ResponseWriter, data interface{}) {
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data,
	})
}

func value(req *http.Request, p string) string {
	return req.FormValue(p)
}

func (h *HttpHandler) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(h.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *HttpHandler) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(h.jwtKey), nil
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

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
