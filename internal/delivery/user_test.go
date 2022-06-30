package delivery

import (
	"bytes"
	"lntvan166/togo/internal/config"
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/pkg/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const ADMIN = "admin"

// var token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY1Mjk2MzcsInVzZXJuYW1lIjoiYWRtaW45In0.eVza3kGW3x6lAkvi4wLXDjY-yRfNQBQ2VrB-aaV7QUs"
var userToLogin = &e.User{
	Username: "admin",
	Password: "admin",
	Plan:     "free",
	MaxTodo:  10,
}

var user = &e.User{
	ID:       1,
	Username: "admin",
	Password: "admin",
	Plan:     "free",
	MaxTodo:  10,
}

var user2 = &e.User{
	ID:       2,
	Username: "admin2",
	Password: "admin2",
	Plan:     "free",
	MaxTodo:  10,
}

var userNotRegistered = &e.User{
	ID:       0,
	Username: "admin",
	Password: "admin",
	Plan:     "free",
	MaxTodo:  10,
}

var users = []*e.User{
	user,
	user2,
}

var userAfterUpgrade = &e.User{
	ID:       1,
	Username: "admin",
	Password: "admin",
	Plan:     "vip",
	MaxTodo:  20,
}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY1Mjk2MzcsInVzZXJuYW1lIjoiYWRtaW45In0.eVza3kGW3x6lAkvi4wLXDjY-yRfNQBQ2VrB-aaV7QUs"

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskDelivery := mock.NewMockTaskUsecase(ctrl)

	userDelivery := NewUserDelivery(userUsecase, taskDelivery)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", LOCALHOST+"/user", nil)

	context.Set(r, "username", ADMIN)
	config.ADMIN = ADMIN

	userUsecase.EXPECT().GetAllUsers().Return(users, nil)

	userDelivery.GetAllUsers(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.Equal(t, `[{"id":1,"username":"admin","password":"admin","plan":"free","max_todo":10},{"id":2,"username":"admin2","password":"admin2","plan":"free","max_todo":10}]`, body)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskDelivery := mock.NewMockTaskUsecase(ctrl)

	userDelivery := NewUserDelivery(userUsecase, taskDelivery)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", LOCALHOST+"/user/1", nil)

	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	context.Set(r, "username", ADMIN)
	config.ADMIN = ADMIN

	userUsecase.EXPECT().GetUserByID(user.ID).Return(user, nil)

	userDelivery.GetUser(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.Equal(t, `{"id":1,"username":"admin","password":"admin","plan":"free","max_todo":10}`, body)
}

func TestDeleteUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskDelivery := mock.NewMockTaskUsecase(ctrl)

	userDelivery := NewUserDelivery(userUsecase, taskDelivery)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", LOCALHOST+"/user/1", nil)

	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	context.Set(r, "username", ADMIN)
	config.ADMIN = ADMIN

	userUsecase.EXPECT().DeleteUserByID(user.ID).Return(nil)

	userDelivery.DeleteUserByID(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}
