package test

import (
	"math/rand"
	"net/http/httptest"
	"testing"

	"api_service/connection"
	"api_service/proto"

	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var username = randSeq(10)
var password = randSeq(10)

func TestCreateAccount(t *testing.T) {

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	var createRequest = proto.CreateRequest{
		Username: username,
		Password: password,
		Name:     "TestCreateAccount",
		Email:    "test-create-account@yopmail.com",
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Create(ctx, &createRequest)
	if err != nil {
		t.Error("CreateAccount error", err)
	}
	if response.IsCreated != true {
		t.Errorf("Unexpected response: %v, wanted %v", response.IsCreated, "true")
	}
}

// func TestLoginAccount(t *testing.T) {

// }

// func TestLogout(t *testing.T) {

// }

// func TestCreateTodo(t *testing.T) {

// }

// func TestGetTodo(t *testing.T) {

// }
