package endpoint

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/usecase/auth"
	"github.com/HoangVyDuong/togo/internal/usecase/task"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

type ToDoHandler struct {
	JWTKey string
	AuthService auth.Service
	TaskService task.Service
}

func (s *ToDoHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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
		AuthAccount(s.AuthService, s.JWTKey)(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = validToken(req, s.JWTKey)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			ListTasks(s.TaskService)(resp, req)
		case http.MethodPost:
			AddTask(s.TaskService)(resp, req)
		}

		return
	}
}

type userAuthKey int8

func validToken(req *http.Request, jwtKey string) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
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

func createToken(id string, jwtKey string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}