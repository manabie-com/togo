package intergrationTest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/joho/godotenv"
)

func TestRegisterHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}

	user := models.RandomNewUser()
	newRequestBody, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))

	r := mux.NewRouter()
	r.HandleFunc("/register", middlewares.ValidUsernameAndHashPassword(bh, bh.Register)).Methods("POST")
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

func TestLoginHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}

	user := models.NewUser{
		Username: "admin",
		Password: "admin",
	}
	newRequestBody, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))

	r := mux.NewRouter()
	r.HandleFunc("/register", bh.Login).Methods("POST")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}
	// Read the response body.
	var data controllers.ResponseToken
	
	if err = json.NewDecoder(w.Body).Decode(&data); err != nil {
		t.Fatal("read body failed")
	}
	users, _ := bh.BaseCtrl.GetAllUser()
	bh.BaseCtrl.DeleteUser(int(users[len(users)-1].Id))

	if strings.Compare(data.Message, "login success") !=0 {
		t.Fatal("Login test failed, expected success got:"+data.Message)
	}
}