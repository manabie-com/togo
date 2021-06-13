package handler

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const XRateLimit = "X-Rate-Limit"

func NewRateLimiter(pool *redis.Pool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userId").(string)
			maxTodo := r.Context().Value("maxTodo").(int)

			conn := pool.Get()
			defer func(conn redis.Conn) {
				_ = conn.Close()
			}(conn)

			val, err := redis.Int(conn.Do("GET", userId))
			if err != nil {
				_, _ = conn.Do("SET", userId, 1)
				_, _ = conn.Do("EXPIRE", userId, 24*60*60) //1 days
				w.Header().Set(XRateLimit, "1")
			} else {
				limit := val + 1
				if limit > maxTodo {
					log.WithFields(log.Fields{
						"userId": userId,
					}).Info("max rate limiting Reached,")

					writeJsonRes(w, 429, errors.New("max rate limiting reached, please try after some time"))
					return
				}
				_, _ = conn.Do("SET", userId, limit)
				w.Header().Set(XRateLimit, strconv.Itoa(limit))

			}
			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
