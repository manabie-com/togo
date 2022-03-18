package handler

import (
	"net/http"

	"github.com/manabie-com/togo/pkg/ctx"

	"github.com/manabie-com/togo/pkg/auth"

	"github.com/manabie-com/togo/pkg/httpx"

	"github.com/manabie-com/togo/internal/user/service"
)

type UserMiddleware struct {
	userService service.UserService
}

// NewCustomerMiddleware ...
func NewUserMiddleware(u service.UserService) *UserMiddleware {
	return &UserMiddleware{userService: u}
}

func (m *UserMiddleware) UserOnly(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r)
		if err != nil {
			httpx.WriteError(w, err)
			return
		}
		claims, err := auth.GetCustomClaimsFromRequest(r)
		if err != nil {
			httpx.WriteError(w, err)
			return
		}
		user, err := m.userService.Authenticate(r.Context(), claims.UserID)
		if err != nil {
			httpx.WriteError(w, err)
			return
		}
		next.ServeHTTP(w, ctx.Set(r, ctx.UserKey, user))
	}
	return http.HandlerFunc(fn)
}
