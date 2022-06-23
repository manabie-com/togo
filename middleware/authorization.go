package middleware

import (
	"errors"
	"fmt"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
	"strings"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			utils.ERROR(w, http.StatusBadRequest, errors.New("malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token := utils.DecodeToken(jwtToken)
			username := fmt.Sprint(token["username"])
			if model.CheckUserExist(username) {
				next.ServeHTTP(w, r)
			} else {
				utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("authorize failed"))
			}
		}

	})
}
