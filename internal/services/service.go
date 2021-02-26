package services

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	authSubKey string = "sub"
	authExpKey        = "exp"
)

var (
	authTokenIsNotValid = errors.New("auth token is not valid")
)

// ToDoService implement HTTP server
type ToDoService struct {
	jwtKey string
	pg     postgres.Database

	server    *http.Server
	serverErr chan error
}

func NewToDoService(jwtKey string, addr string, pg postgres.Database) *ToDoService {
	s := &ToDoService{
		jwtKey: jwtKey,
		pg:     pg,
		server: &http.Server{
			Addr: addr,
		},
		serverErr: make(chan error, 1),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", s.setHeaders(s.createTokenHandler))
	mux.HandleFunc("/tasks", s.setHeaders(s.authHandler(s.tasksHandler())))
	s.server.Handler = mux

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.serverErr <- err
		}
	}()

	return s
}

func (s *ToDoService) HttpServerErr() <-chan error {
	return s.serverErr
}

func (s *ToDoService) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *ToDoService) createToken(id interface{}) (string, error) {
	claims := jwt.MapClaims{
		authSubKey: id,
		authExpKey: time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtKey))
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, error) {
	authToken := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	parsedToken, err := jwt.ParseWithClaims(authToken, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
	})
	if err != nil {
		return req, err
	}

	if !parsedToken.Valid {
		return req, authTokenIsNotValid
	}

	id, ok := claims[authSubKey].(float64)
	if !ok {
		return req, authTokenIsNotValid
	}

	req = req.WithContext(context.WithValue(req.Context(), authSubKey, int(id)))
	return req, nil
}

func userIDFromCtx(ctx context.Context) (int, bool) {
	v := ctx.Value(authSubKey)
	id, ok := v.(int)
	return id, ok
}
