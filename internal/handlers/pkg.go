package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/handlers/auth"
	"github.com/manabie-com/togo/internal/handlers/tasks"
	requestUtils "github.com/manabie-com/togo/internal/utils/request"
	"github.com/manabie-com/togo/internal/utils/response"
	"github.com/manabie-com/togo/internal/utils/token"
)

type HandleFunc func(http.ResponseWriter, *http.Request) error

type RESTHandler interface {
	Post(http.ResponseWriter, *http.Request) error
	Get(http.ResponseWriter, *http.Request) error
}

type Handlers struct {
	JWTSecret          string
	DB                 *sql.DB
	RESTHandlerMap     map[string]RESTHandler
	CustomerHandlerMap map[string]HandleFunc
	AuthHandler        auth.AuthHandler
}

func (h *Handlers) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	req = requestUtils.SetJWTSecret(req, h.JWTSecret)
	handleFunc := h.getRestHandleFunc(req)
	if handleFunc == nil {
		response.NotFoundPath(resp)
		return
	}
	err := h.authMiddleware(resp, req, handleFunc)
	if err != nil {
		h.handlerErrors(resp, err)
	}
}

// middleware for validating token
func (h *Handlers) authMiddleware(response http.ResponseWriter, request *http.Request, handler HandleFunc) error {
	if request.URL.Path == LoginPath {
		return handler(response, request)
	}
	// validate user here
	jwtToken := token.GetToken(request)
	isValid, claim := token.Validate(jwtToken, h.JWTSecret)
	if !isValid {
		return consts.ErrUnauthorized
	}
	request = requestUtils.SetUserID(request, claim.UserID)
	return handler(response, request)
}

// handlers
func (h *Handlers) LoadHandlers() {
	h.RESTHandlerMap = map[string]RESTHandler{
		TaskPath: tasks.NewHandler(h.DB),
	}
	h.AuthHandler = auth.NewAuthHandler(h.DB)
}

// Get associate handle func base on path and method
func (h *Handlers) getRestHandleFunc(req *http.Request) HandleFunc {
	if req.URL.Path == LoginPath && req.Method == http.MethodGet {
		return h.AuthHandler.Login
	}
	restHandler := h.RESTHandlerMap[req.URL.Path]
	if restHandler == nil {
		return nil
	}
	switch req.Method {
	case http.MethodGet:
		return restHandler.Get
	case http.MethodPost:
		return restHandler.Post
	}
	return nil
}

// Handle return errors depend on the error type
func (h *Handlers) handlerErrors(resp http.ResponseWriter, err error) {
	errObject, ok := err.(consts.Error)
	// Unknown error
	if !ok {
		resp.WriteHeader(http.StatusInternalServerError)
		_ = response.Error(resp, err)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	switch errObject {
	case consts.ErrInvalidAuth, consts.ErrUnauthorized:
		resp.WriteHeader(http.StatusBadRequest)
		_ = response.Error(resp, errObject)
		return
	}
}
