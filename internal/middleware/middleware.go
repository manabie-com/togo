package middleware

import (
	"errors"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/internal/usecase"
	"lntvan166/togo/pkg"
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			pkg.ERROR(w, http.StatusBadRequest, errors.New("invalid authorization header"), "")
		} else {
			jwtToken := authHeader[1]

			token, err := pkg.DecodeToken(jwtToken)
			if err != nil {
				pkg.ERROR(w, http.StatusInternalServerError, err, "failed to decode token!")
				return
			}
			if token["username"] == nil {
				pkg.ERROR(w, http.StatusBadRequest, errors.New("invalid token"), "")
				return
			}
			username := token["username"].(string)

			checkUserExist, err := usecase.CheckUserExist(username)
			if err != nil {
				pkg.ERROR(w, http.StatusInternalServerError, err, "failed to check user exist!")
				return
			}

			if checkUserExist {
				context.Set(r, "username", username)

				next.ServeHTTP(w, r)
			} else {
				pkg.ERROR(w, http.StatusBadRequest, errors.New("authorize failed"), "")
				return
			}
		}

	})
}

func AdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			pkg.ERROR(w, http.StatusBadRequest, errors.New("invalid authorization header"), "")
		} else {
			jwtToken := authHeader[1]

			token, err := pkg.DecodeToken(jwtToken)
			if err != nil {
				pkg.ERROR(w, http.StatusInternalServerError, err, "failed to decode token!")
				return
			}
			if token["username"] == nil {
				pkg.ERROR(w, http.StatusBadRequest, errors.New("invalid token"), "")
				return
			}
			username := token["username"].(string)

			checkUserExist, err := usecase.CheckUserExist(username)
			if err != nil {
				pkg.ERROR(w, http.StatusInternalServerError, err, "failed to check user exist!")
				return
			}

			if checkUserExist {

				if username == config.ADMIN {
					context.Set(r, "username", username)

					next.ServeHTTP(w, r)
				} else {
					pkg.ERROR(w, http.StatusBadRequest, errors.New("authorize failed"), "")
					return
				}
			} else {
				pkg.ERROR(w, http.StatusBadRequest, errors.New("authorize failed"), "")
				return
			}
		}

	})
}
