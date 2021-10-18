package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/jericogantuangco/togo/internal/storages/mock"
	"github.com/jericogantuangco/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
)

func TestLoginAPIOK(t *testing.T) {
	user := postgres.User{
		Username: "testUser",
		Password: "password",
		MaxTodo:  5,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		RetrieveUser(gomock.Any(), gomock.Eq(user.Username)).
		Times(1).
		Return(user, nil)

	server, err := NewServer(store)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	url := "/login"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	q := req.URL.Query()

	q.Add("user_id", user.Username)
	q.Add("password", user.Password)
	req.URL.RawQuery = q.Encode()

	addAuth(t, req, server.TokenMaker, authorizationTypeBearer, "testUser", 10*time.Minute)
	server.Router.ServeHTTP(recorder, req)
	checkResponse(t, recorder)

}
