package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	cacheKey "github.com/quochungphp/go-test-assignment/src/pkgs/cache_key"
	"github.com/quochungphp/go-test-assignment/src/pkgs/redis"
	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := token.ValidToken(r)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}

		// Reflect user info to var
		tk := token.Token{}
		accessUserInfo, err := tk.ExtractToken(r)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(fmt.Sprintf("Unauthorized :%s", err))
			return
		}

		// Check black list token
		cacheKey := cacheKey.TokenBlackListCacheKey(fmt.Sprintf("%s-%d", accessUserInfo.CorrelationID, accessUserInfo.UserID))
		isBlackListToken, err := redis.GetItem(cacheKey, "")
		if isBlackListToken != nil && err == nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(fmt.Sprintf("Expired session:%s ", accessUserInfo.CorrelationID))
			return
		}

		next.ServeHTTP(w, r)
	})
}
