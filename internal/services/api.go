package services

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

type API struct {
	Router *mux.Router
	Port int
	db storages.IDatabase
	cfg *config.Config
	redisPool *redis.Pool
}

func NewAPI(cfg *config.Config) (*API, error) {
	db, err := storages.NewDatabase(cfg)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	api := &API{
		Port: cfg.Port,
		db: db,
		Router: mux.NewRouter(),
		cfg: cfg,
		redisPool: newRedisPool(cfg),
	}
	// middleware
	api.Router.Use(api.ValidToken)

	// add task
	todo := NewTodoService(db, cfg.JWTKey, api.redisPool, cfg.MaxTodo)
	todo.AddHandler(api)
	return api, nil
}

func (api *API) Start() {
	defer api.close()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", api.Port), cors.AllowAll().Handler(api.Router)))
}

func (api *API) close() {
	api.redisPool.Close()
}

func (api *API) ValidToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")
		path := req.URL.Path
		if path != "/login" && path != "/signup" {
			var ok bool
			req, ok = api.validToken(req)
			if !ok {
				resp.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(resp, req)
	})
}

func (api *API) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(api.cfg.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}

func newRedisPool(cfg *config.Config) *redis.Pool {
	return &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Redis.Address)
			if err != nil {
				return nil, err
			}
			if cfg.Redis.Password != "" {
				if _, err := c.Do("AUTH", cfg.Redis.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", cfg.Redis.DatabaseNum); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		MaxIdle:         cfg.Redis.MaxIdle,
		MaxActive:       cfg.Redis.MaxActive,
		IdleTimeout:     time.Duration(cfg.Redis.MaxIdleTimeout) * time.Second,
		Wait:            cfg.Redis.Wait,
		MaxConnLifetime: time.Duration(cfg.Redis.ConnectTimeout) * time.Second,
	}
}