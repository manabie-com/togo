package middleware

import (
	"errors"
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
			utils.ERROR(w, http.StatusBadRequest, errors.New("invalid authorization header"), "")
		} else {
			jwtToken := authHeader[1]

			token, err := utils.DecodeToken(jwtToken)
			if err != nil {
				utils.ERROR(w, http.StatusInternalServerError, err, "failed to decode token!")
				return
			}
			if token["username"] == nil {
				utils.ERROR(w, http.StatusBadRequest, errors.New("invalid token"), "")
				return
			}
			username := token["username"].(string)

			checkUserExist, err := model.CheckUserExist(username)
			if err != nil {
				utils.ERROR(w, http.StatusInternalServerError, err, "failed to check user exist!")
				return
			}

			if checkUserExist {
				context.Set(r, "username", username)

				next.ServeHTTP(w, r)
			} else {
				utils.ERROR(w, http.StatusBadRequest, errors.New("authorize failed"), "")
				return
			}
		}

	})
}
