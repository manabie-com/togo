package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/stretchr/testify/assert"
)

const (
	AdminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY0MzQ0NzgsImlkIjoxLCJ1c2VybmFtZSI6ImFkbWluIn0.fprKS6TBv8L95_ZqD_jwbGRblm9hnWKi5vQVdGQEtqM"
	Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYxMDUzNDEsImlkIjozLCJ1c2VybmFtZSI6Imh1dWxvYyJ9.MqHypeN42fopG5jHWEjk6bu9m7wSENqLewBGq9VC3sA"
)

type httptestHandler struct {
	w http.ResponseWriter
	r *http.Request
}

// test middleware Logging
func TestLoggingVerified(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8000/users/", nil)
	req.Header.Add("token", Token) // add token to header
	Handler := httptestHandler{
		w: w,
		r: req,
	}

	logging := LoggingVerified(Handler)
	logging.ServeHTTP(Handler.w, Handler.r)

	userid := context.Get(Handler.r, "userid")
	id := context.Get(Handler.r, "id")
	if id != 3 && userid != 3 {
		t.Fatal("logging test failed")
	}
}


// test middleware AdminVerified
func TestAdminVerified(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8000/users/", nil)
	req.Header.Add("token", AdminToken) // add token to header
	Handler := httptestHandler{
		w: w,
		r: req,
	}

	adminverified := AdminVerified(Handler)
	adminverified.ServeHTTP(Handler.w, Handler.r)

	userid := context.Get(Handler.r, "userid")
	if userid != 1 {
		t.Fatal("logging test failed")
	}
}

// test middleware MiddlewareID
func TestMiddlewareID(t *testing.T) {
	w := httptest.NewRecorder()
	id := "random text"
	req := httptest.NewRequest(http.MethodGet, "localhost:8000/users", nil)
	vars := map[string]string{
		"id": id,
	}
	req = mux.SetURLVars(req, vars) // set url variable
	Handler := httptestHandler{
		w: w,
		r: req,
	}

	idHandler := MiddlewareID(Handler)
	idHandler.ServeHTTP(Handler.w, Handler.r)
	// newid := context.Get(Handler.r, "id")
	res := w.Result()
	resbody, _ := ioutil.ReadAll(res.Body)
	if strings.Compare(string(resbody), "id url need to be a number") != 1 {
		t.Fatal("test id failed")
	}
}

// test middleware ValidUsernameAndHashPassword
func TestValidUsernameAndHashPassword(t *testing.T) {
	db, mock := models.NewMock()
	dbConn := models.NewdbConn(db)
	bh := controllers.NewBaseHandler(dbConn)
	
	newUser := models.RandomNewUser() // create a new random user
	mock.ExpectQuery(regexp.QuoteMeta(models.QueryAllUsernameText)).WithArgs(newUser.Username).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "limittask"}))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "localhost:8000/users", nil)
	requestBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal("make body request failed")
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody)) // send newuser to body of the request to the middlewares
	
	Handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	valid := ValidUsernameAndHashPassword(bh, Handler)

	valid.ServeHTTP(w, req)
	var bodyJSON models.NewUser
	if err := json.NewDecoder(req.Body).Decode(&bodyJSON); err != nil { // get info out after middleware process
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}
	assert.Equal(t, bodyJSON.Username, newUser.Username)
	assert.NotEqual(t, bodyJSON.Password, newUser.Password) // after hash password can't be the same
	assert.Equal(t, bodyJSON.LimitTask, newUser.LimitTask)
}

func (l httptestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // function to implement interface
}
