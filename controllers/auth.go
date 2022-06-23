package controllers

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"

	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/users/signup", "/api/users/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                    //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header
		if tokenHeader == "" {                       //Token is missing, returns with error code 403 Unauthorized
			u.Respond(w, http.StatusForbidden, "Failure", "Missing auth token", nil)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			u.Respond(w, http.StatusForbidden, "Failure", "Invalid/Malformed auth token", nil)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_TOKEN")), nil
		})
		if err != nil { //Malformed token, returns with http code 403 as usual
			u.Respond(w, http.StatusForbidden, "Failure", "Malformed authentication token", err.Error())
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			u.Respond(w, http.StatusForbidden, "Failure", "Token is not valid.", nil)
			return
		}
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		//Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx)) //proceed in the middleware chain!
	})
}
