package mux

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

// Init new router with default log
func Init() *chi.Mux {
	return chi.NewRouter()
}

//InitWithLogger init with zap log param
func InitWithLogger(logger *zap.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(ChiLogger(logger))
	return router
}

// ChiLogger is a middleware logger for go-chi
func ChiLogger(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)
			defer func() {
				l.Info("served",
					zap.String("proto", r.Proto),
					zap.String("path", r.URL.Path),
					zap.Duration("latency_nanosecond", time.Since(t1)), // the unit is ns
					zap.Int("status", ww.Status()),
					zap.Int("size_byte", ww.BytesWritten()),                // unit bytes
					zap.String("req_id", middleware.GetReqID(r.Context()))) // for the future, maybe when using api gateway, it'll add reqID
			}()
		}
		return http.HandlerFunc(fn)
	}
}
