package middleware

import (
	"errors"
	"fmt"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			utils.ERROR(w, http.StatusBadRequest, errors.New("malformed Token"))
		} else {
			jwtToken := authHeader[1]

			token, err := utils.DecodeToken(jwtToken)
			if err != nil {
				utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error(), "when decode token"))
			}

			username := token["username"].(string)
			if model.CheckUserExist(username) {
				context.Set(r, "username", username)

				next.ServeHTTP(w, r)
			} else {
				utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("authorize failed"))
			}
		}

	})
}
