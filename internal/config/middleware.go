package config

import (
	"github.com/manabie-com/togo/internal/response"
	"github.com/manabie-com/togo/internal/utils"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt := utils.NewJwt()

		var ok bool
		r, ok = jwt.ValidToken(r)
		if !ok {
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		// Access context values in handlers like this
		// props, _ := r.Context().Value("props").(jwt.MapClaims)
		next.ServeHTTP(w, r)
	})
}
