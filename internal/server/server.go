package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/banhquocdanh/togo/internal/services"
	"log"
	"net/http"
)

type ToDoHttpServer struct {
	srv    *services.ToDoService
	JWTKey string
}

func NewToDoHttpServer(jwtKey string, srv *services.ToDoService) *ToDoHttpServer {
	return &ToDoHttpServer{
		srv:    srv,
		JWTKey: jwtKey,
	}

}

func (s *ToDoHttpServer) Listen(port int) error {
	log.Printf("Listen service on port :%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s)
}

func (s *ToDoHttpServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.login(resp, req)
		return
	case "/tasks":
		token := req.Header.Get("Authorization")
		userID, ok := s.srv.ValidToken(token, s.JWTKey)
		if !ok {
			HttpUnauthorized(resp, "invalid token")
			return
		}
		req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), userID))

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	}
}

func WriteHttpResponse(resp http.ResponseWriter, code int, data Response) {
	if resp.Header().Get("Content-Type") == "" {
		resp.Header().Set("Content-Type", "application/json")
	}
	resp.WriteHeader(code)
	_ = json.NewEncoder(resp).Encode(data)
}

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func HttpResponseSuccess(resp http.ResponseWriter, data interface{}) {
	WriteHttpResponse(resp, http.StatusOK, Response{Data: data})
}

func HttpResponseBadRequest(resp http.ResponseWriter, msg string) {
	WriteHttpResponse(resp, http.StatusBadRequest, Response{Error: msg})
}

func HttpResponseInternalError(resp http.ResponseWriter, err error) {
	WriteHttpResponse(resp, http.StatusInternalServerError, Response{Error: err.Error()})
}

func HttpUnauthorized(resp http.ResponseWriter, msg string) {
	WriteHttpResponse(resp, http.StatusUnauthorized, Response{Error: msg})
}

func (s *ToDoHttpServer) listTasks(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	createDate := req.FormValue("created_date")
	if createDate == "" {
		HttpResponseBadRequest(resp, "Invalid create_date")
		return
	}
	tasks, err := s.srv.ListTasks(req.Context(), userID, createDate)

	if err != nil {
		HttpResponseInternalError(resp, err)
		return
	}

	HttpResponseSuccess(resp, tasks)
}

type AddTaskRequest struct {
	Content string `json:"content"`
}

func (s *ToDoHttpServer) addTask(resp http.ResponseWriter, req *http.Request) {
	addTaskReq := &AddTaskRequest{}
	err := json.NewDecoder(req.Body).Decode(addTaskReq)
	defer req.Body.Close()
	if err != nil {
		HttpResponseInternalError(resp, err)
		return
	}

	if addTaskReq.Content == "" {
		HttpResponseBadRequest(resp, "invalid task's content")
		return
	}
	userID, _ := userIDFromCtx(req.Context())

	task, err := s.srv.AddTask(req.Context(), userID, addTaskReq.Content)
	if err != nil {
		HttpResponseInternalError(resp, err)
		return
	}

	HttpResponseSuccess(resp, task)
}

func (s *ToDoHttpServer) login(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("user_id")
	pw := req.FormValue("password")

	token, err := s.srv.Login(req.Context(), id, pw, s.JWTKey)
	if err != nil {
		HttpUnauthorized(resp, "incorrect user_id/pwd")
		return
	}

	HttpResponseSuccess(resp, token)
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
