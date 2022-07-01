package intergrationTest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	AdminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY0MzQ0NzgsImlkIjoxLCJ1c2VybmFtZSI6ImFkbWluIn0.fprKS6TBv8L95_ZqD_jwbGRblm9hnWKi5vQVdGQEtqM"
	AdminID    = 1
)

func TestGetAllUsersHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	var users []models.NewUser
	users = append(users, models.RandomNewUser(), models.RandomNewUser())

	for _, user := range users {
		bh.BaseCtrl.InsertUser(user)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified)
	r.HandleFunc("/users", bh.ResponseAllUser).Methods("GET")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	returnedUsers := []models.User{}
	if err := json.Unmarshal(data, &returnedUsers); err != nil {
		t.Errorf("Returned user list is invalid JSON. Got: %s", data)
	}
	// delete after insert to check
	bh.BaseCtrl.DeleteUser(int(returnedUsers[len(returnedUsers)-1].Id))
	bh.BaseCtrl.DeleteUser(int(returnedUsers[len(returnedUsers)-2].Id))

	if len(returnedUsers) < len(users) {
		t.Errorf("Returned user list is an invalid length. Expected %d. Got %d instead", len(users), len(returnedUsers))
	}
	count := 0
	for _, returnUser := range returnedUsers {
		if returnUser.Username == users[0].Username || returnUser.Username == users[1].Username {
			count++
		}
	}
	if count != len(users) {
		t.Fatal("Returned user list is different")
	}
}

func TestGetOneUserHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "id", 1)

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified, middlewares.MiddlewareID)
	r.HandleFunc("/users/{id}", bh.ResponseOneUser).Methods("GET")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	var returnedUsers models.NewUser
	if err := json.NewDecoder(w.Body).Decode(&returnedUsers); err != nil {
		t.Errorf("Returned user is invalid JSON. Got: %s", returnedUsers.Username)
	}

	if returnedUsers.Username != "admin" {
		t.Fatal("Get return wrong user expected: admin, got: " + returnedUsers.Username)
	}
}

func TestDeleteUserHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	user := models.RandomNewUser()
	bh.BaseCtrl.InsertUser(user)
	users, err := bh.BaseCtrl.GetAllUser()
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/users/%v", users[len(users)-1].Id), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "id", users[len(users)-1].Id)

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified, middlewares.MiddlewareID)
	r.HandleFunc("/users/{id}", bh.DeleteFromUser).Methods("DELETE")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Compare(string(data), "message: delete success") != 0 {
		t.Error("Expected message: delete success, got " + string(data))
	}
}

func TestUpdateUserHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	user := models.RandomNewUser()
	users, err := bh.BaseCtrl.GetAllUser()
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/users/%v", users[len(users)-1].Id), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "id", users[len(users)-1].Id)
	newRequestBody, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified, middlewares.MiddlewareID)
	r.HandleFunc("/users/{id}", bh.UpdateToUser).Methods("PUT")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}
	users, _ = bh.BaseCtrl.GetAllUser()
	// Read the response body.
	var returnedUsers models.NewUser
	if err := json.NewDecoder(w.Body).Decode(&returnedUsers); err != nil {
		t.Errorf("Returned user is invalid JSON. Got: %s", returnedUsers.Username)
	}

	if strings.Compare(returnedUsers.Username, users[len(users)-1].Username) != 0 {
		t.Fatal("Get return wrong user expected: " + users[len(users)-1].Username + ", got: " + returnedUsers.Username)
	}
}

func TestCreateUserHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)

	user := models.RandomNewUser()
	newRequestBody, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified)
	r.HandleFunc("/users", bh.CreateUser).Methods("POST")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}
	users, _ := bh.BaseCtrl.GetAllUser()
	// Read the response body.
	var returnedUsers models.NewUser
	if err := json.NewDecoder(w.Body).Decode(&returnedUsers); err != nil {
		t.Errorf("Returned user is invalid JSON. Got: %s", returnedUsers.Username)
	}

	bh.BaseCtrl.DeleteUser(int(users[len(users)-1].Id))

	if strings.Compare(returnedUsers.Username, users[len(users)-1].Username) != 0 {
		t.Fatal("Get return wrong user expected: " + users[len(users)-1].Username + ", got: " + returnedUsers.Username)
	}
}
