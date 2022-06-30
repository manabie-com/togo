package delivery

import (
	"lntvan166/togo/pkg/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var LOCALHOST = "http://localhost:8080"

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)

	authDelivery := NewAuthDelivery(userUsecase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", LOCALHOST+"/auth/register", strings.NewReader(`
	{
		"username": "admin",
		"password":"admin"
	}`))

	userUsecase.EXPECT().Register(userNotRegistered).Return(nil)

	r.Header.Set("Content-Type", "application/json")

	authDelivery.Register(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)

	authDelivery := NewAuthDelivery(userUsecase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", LOCALHOST+"/auth/login", strings.NewReader(`
	{
		"username": "admin",
		"password":"admin"
	}`))

	userUsecase.EXPECT().Login(userToLogin).Return(token, nil).AnyTimes()

	r.Header.Set("Content-Type", "application/json")

	authDelivery.Login(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}
