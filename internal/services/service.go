package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	authUserIdKey string = "user_id"
	authExpKey           = "exp"
)

var (
	authTokenIsNotValid  = errors.New("auth token is not valid")
	authUserIdIsRequired = errors.New(fmt.Sprintf(`'%v' is required`, authUserIdKey))
)

// ToDoService implement HTTP server
type ToDoService struct {
	server    *http.Server
	serverErr chan error

	JWTKey string
	Store  *sqllite.LiteDB
}

func NewToDoService(db *sql.DB) *ToDoService {
	s := &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
		server: &http.Server{
			Addr: ":5050",
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

func (s *ToDoService) createToken(id string) (string, error) {
	claims := jwt.MapClaims{
		authUserIdKey: id,
		authExpKey:     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTKey))
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, error) {
	authToken := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	parsedToken, err := jwt.ParseWithClaims(authToken, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		return req, err
	}

	if !parsedToken.Valid {
		return req, authTokenIsNotValid
	}

	id, ok := claims[authUserIdKey].(string)
	if !ok {
		return req, authUserIdIsRequired
	}

	req = req.WithContext(context.WithValue(req.Context(), authUserIdKey, id))
	return req, nil
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(authUserIdKey)
	id, ok := v.(string)
	return id, ok
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}