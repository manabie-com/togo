package middleware

import (
	"github.com/manabie-com/togo/internal/pkg/logger"
	"net/http"
)

var mbLogger logger.Logger

func AddCors(next http.Handler) http.Handler {
	mbLogger = logger.WithPrefix("middleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mbLogger.Infoln(r.Method, r.URL.Path)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
