package api

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"manabie.com/togo"
	"net/http"
	"strconv"
	"strings"
)

const (
	AUTHORIZATION_HEADER string = "Authorization"
	USER_HEADER                 = "UserHeader"
)

type Cfg struct {
	Port               int
	UserCrudService    togo.UserCrudService
	TaskLimiterService togo.TaskLimiterService
}

var port int
var parser = new(jwt.Parser)

func Init(c Cfg) {
	port = c.Port
	userCrudService = c.UserCrudService
	taskLimiterService = c.TaskLimiterService
}

func httpMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) > 2 && r.URL.Path[0:2] == "/e" {
			r.Header.Del(USER_HEADER)
			jwtAuthor := strings.TrimSpace(r.Header.Get(AUTHORIZATION_HEADER))
			if userId, err := getUserIdFromAuthorizationHeader(jwtAuthor); err == nil {
				r.Header.Set(USER_HEADER, userId)
			} else {
				log.Error(err)
			}
		}
		h.ServeHTTP(w, r)
	})
}

func getUserIdFromAuthorizationHeader(jwtData string) (string, error) {
	token, _, err := parser.ParseUnverified(jwtData, jwt.MapClaims{})
	if err != nil {
		return "", errors.Wrapf(err, "fail to parse jwt %s", jwtData)
	}

	if token == nil {
		return "", errors.New("token must not be null")
	}

	if token.Claims == nil {
		return "", errors.New("token is invalid, claims must not be null")
	}

	mapClaims := token.Claims.(jwt.MapClaims)
	if mapClaims == nil {
		return "", errors.New("token is invalid, claims is invalid, claims must be mapClaims type")
	}

	userIdValue, isExisting := mapClaims["userId"]
	if !isExisting {
		return "", errors.New("token is invalid, userId must be existing")
	}

	userId, ok := userIdValue.(float64)

	if !ok {
		return "", errors.New("token is invalid, sub claims must be string type")
	}

	return strconv.Itoa(int(userId)), nil
}

func Run() error {
	r := mux.NewRouter()
	r.Use(httpMiddleware)
	r.HandleFunc("/", handleHome).Methods("GET")
	back := r.PathPrefix("/b/manabie").Subrouter()
	back.HandleFunc("/create-task", CreateTask).Methods("POST")
	back.HandleFunc("/set-limit", SetTaskLimiter).Methods("POST")

	job := r.PathPrefix("/j/manabie").Subrouter()
	job.HandleFunc("/reset-daily-tasks", ResetDailyTask).Methods("POST")

	external := r.PathPrefix("/e/manabie").Subrouter()
	external.HandleFunc("/do-task/{taskId}", DoTask).Methods("POST")

	publish := r.PathPrefix("/p/manabie").Subrouter()
	publish.HandleFunc("/signup", CreateUser).Methods("POST")
	publish.HandleFunc("/login", Login).Methods("POST")

	http.Handle("/", r)

	log.Infof("manabie service is running at port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
