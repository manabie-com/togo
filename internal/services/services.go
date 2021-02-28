package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/pkg/common/xerrors"
	"log"
	"net/http"
	"reflect"
)

type AuthType int

const (
	None = AuthType(0)
	User = AuthType(1)
)

type Decl struct {
	Auth        AuthType
	HandlerFunc interface{}
}

type ToDoService struct {
	jwtKey      string
	userService *UserService
	taskService *TaskService
	acl         map[string]map[string]Decl
}

func NewToDoService(
	db *sql.DB, JWTKey string,
	maxTodo int) *ToDoService {
	userService := NewUserService(db, maxTodo, JWTKey)
	taskService := NewTaskService(db)

	return &ToDoService{
		jwtKey:      JWTKey,
		userService: userService,
		taskService: taskService,
		acl: map[string]map[string]Decl{
			"/register": {
				http.MethodPost: Decl{
					HandlerFunc: userService.Register,
					Auth:        None,
				},
			},
			"/login": {
				http.MethodPost: Decl{
					HandlerFunc: userService.Login,
					Auth:        None,
				},
			},
			"/tasks": {
				http.MethodGet: Decl{
					HandlerFunc: taskService.ListTasks,
					Auth:        User,
				},
				http.MethodPost: Decl{
					HandlerFunc: taskService.AddTask,
					Auth:        User,
				},
			},
		},
	}
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	handler, ok := s.acl[req.URL.Path]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	decl, ok := handler[req.Method]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	handlerFunc := decl.HandlerFunc
	auth := decl.Auth

	// authorization
	switch auth {
	case User:
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
	case None:
		// no-op
	}

	// Call functions of services
	// Todo: check function (type of arguments, outputs)
	typ := reflect.TypeOf(handlerFunc)
	funcArgs := reflect.New(typ.In(1).Elem()).Interface()

	// parse body and request values
	switch req.Method {
	case http.MethodPost:
		err := json.NewDecoder(req.Body).Decode(funcArgs)
		defer req.Body.Close()
		if err != nil {
			log.Panicf("error %v", err)
			return
		}

	case http.MethodGet:
		funcArgsTyp := reflect.TypeOf(funcArgs).Elem()
		for i := 0; i < funcArgsTyp.NumField(); i++ {
			jsonTag := funcArgsTyp.Field(i).Tag.Get("json")
			reflect.ValueOf(funcArgs).Elem().Field(i).Set(reflect.ValueOf(req.FormValue(jsonTag)))
		}
	}

	ctx := req.Context()
	outs := reflect.ValueOf(handlerFunc).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(funcArgs)})
	resultFunc := outs[0].Interface()
	errFunc := outs[1].Interface()

	resp.Header().Set("Content-Type", "application/json")
	if errFunc != nil {
		err := errFunc.(xerrors.XError)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Message,
		})
		return
	} else {
		json.NewEncoder(resp).Encode(map[string]interface{}{
			"data": resultFunc,
		})
	}
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
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
