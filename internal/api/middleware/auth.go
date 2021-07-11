package middleware

import (
	"context"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"github.com/manabie-com/togo/internal/pkg/token"
	"net/http"
)

var (
	loginPath = "/login"
)

const (
	AuthorizationHeader = "Authorization"
)

func ValidToken(next http.Handler, cfg *config.Config, generator token.Generator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case loginPath:
			mbLogger.Infoln("login request, skip validate")
		default:
			tokenStr := r.Header.Get(AuthorizationHeader)

			id, err := generator.ValidateToken(tokenStr)
			if err != nil {
				mbLogger.Errorf("validate token err: %v", err)
				http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
				return
			}

			mbLogger = logger.WithFields(map[string]interface{}{
				"user_id": id,
				"path":    r.URL.Path,
				"method":  r.Method,
			})

			r = r.WithContext(context.WithValue(r.Context(), logger.MBLoggerConText, mbLogger))
			r = r.WithContext(context.WithValue(r.Context(), token.UserIDField, id))
		}

		next.ServeHTTP(w, r)
	})
}
