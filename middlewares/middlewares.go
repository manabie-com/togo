package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/jwt"
	"github.com/huynhhuuloc129/todo/models"
)

// check if logging as admin or not
func AdminVerified(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, userid, ok := jwt.CheckToken(w, r)
		if !ok || strings.ToLower(username) != "admin" {
			http.Error(w, "you need to login as ADMIN first to perform this action", http.StatusUnauthorized)
			return
		}
		context.Set(r, "userid", userid)
		next.ServeHTTP(w, r)
	})
}

// check if logging or not
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, userid, ok := jwt.CheckToken(w, r)
		if !ok {
			http.Error(w, "you need to login first to perform this action", http.StatusUnauthorized)
			return
		}
		context.Set(r, "userid", userid)
		context.Set(r, "id", userid)

		next.ServeHTTP(w, r)
	})
}

// check ID is a number or not
func MiddlewareID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "id url need to be a number", http.StatusBadRequest)
			return
		}
		context.Set(r, "id", id)
		next.ServeHTTP(w, r)
	})
}

// check username duplicate/valid or not and hash password incoming
func ValidUsernameAndHashPassword(bh *controllers.BaseHandler, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "read body request failed, err: "+err.Error(), http.StatusBadRequest)
			return
		}

		var bodyJSON models.NewUser
		err = json.Unmarshal(requestBody, &bodyJSON)
		if err != nil {
			http.Error(w, "unmarshal body failed, err: "+err.Error(), http.StatusBadRequest)
			return
		}
		bodyJSON.Password, err = models.Hash(bodyJSON.Password)
		if err != nil {
			http.Error(w, "hash password failed, err: "+err.Error(), http.StatusBadRequest)
			return
		}

		newRequestBody, err := json.Marshal(bodyJSON)
		if err != nil {
			http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, ok := bh.BaseCtrl.CheckUserNameExist(bodyJSON.Username); ok { // Check username exist or not
			http.Error(w, "this username already exist", http.StatusNotAcceptable)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))
		next.ServeHTTP(w, r)
	})
}

func CheckLimitTaskUserMiddleware(bh *controllers.BaseHandler, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userid, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "userid"))) // get userid from login
		if ok, err := bh.BaseCtrl.CheckLimitTaskUser(userid); !ok {
			if err != nil {
				http.Error(w, err.Error(), http.StatusFailedDependency)
				return
			}
			http.Error(w, "The limit of today is full or error occur during check limit", http.StatusFailedDependency)
			return
		}
		next.ServeHTTP(w, r)
	})
}
