package controllers_test

import (
	"TOGO/configs"
	"TOGO/controllers"
	"TOGO/middleware"
	"TOGO/models"

	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
}

var NewId primitive.ObjectID

var a App
var NewToken string

// token from
var tokenAdmin string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA5NTM1MDAsImlkIjoiNjJiZDY0NDRlNTIyYjdhYmQwODY1Mzg3Iiwicm9sZSI6ImFkbWluIn0.iHhmOtWBXa7kYLe4z-3MbIOwAqPxvWjirZwmZ7OSdkA"
var tokenMain string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA5NTM0MDcsImlkIjoiNjJiZDY4MjYyOWFmNTIwMzU2ZjhmYzBhIiwicm9sZSI6InVzZXIifQ.lne_0e56hJGui22EWmZN6RcSS4iB0eu4oSyMQTTcQoo"
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func TestMain(m *testing.M) {
	a.Router = mux.NewRouter()
	//run database
	configs.ConnectDB()

	UserRoute((a.Router))
	//routes
	code := m.Run()
	defer os.Exit(code)
}

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/{Id}", controllers.GetUser()).Methods("GET")
	router.HandleFunc("/user/{Id}", middleware.AuthMiddleware(controllers.DeleteUser())).Methods("DELETE")
	router.Handle("/task/{id}", middleware.AuthMiddleware(controllers.GetOneTask())).Methods("GET")
	router.Handle("/task/{id}", middleware.AuthMiddleware(controllers.UpdateTask())).Methods("PUT")
	router.Handle("/task/status/{id}", middleware.AuthMiddleware(controllers.UpdateTaskStatus())).Methods("PUT")

}

func ExcuteRoute(r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.Handler(a.Router)
	handler.ServeHTTP(rr, r)
	return rr
}

func CreateTestUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		Id := primitive.NewObjectID()
		username := "testuser"
		password := "123456"
		name := "test"

		hashPwd, _ := models.HashPassword(password)
		newUser := models.User{
			Id:       Id,
			Username: username,
			Password: hashPwd,
			Name:     name,
			Limit:    10,
			Role:     "user",
		}
		NewId = Id
		// add obj
		_, _ = userCollection.InsertOne(ctx, newUser)
		Token, _ := middleware.CreateToken(newUser.Id, newUser.Role)
		NewToken = "Bearer " + Token
	}
}
