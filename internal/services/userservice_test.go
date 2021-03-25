package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/configurations"
	mockdb "github.com/manabie-com/togo/internal/storages/mock"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
)

type loginHandlerTCsResult struct {
	Code  int
	Data  string
	Valid bool
}

var loginHandlerTCs = []struct {
	LoginRequest *http.Request
	LoginInfo    postgres.ValidateUserParams
	LoginResult  loginHandlerTCsResult
}{
	{
		LoginRequest: &http.Request{
			Method: http.MethodGet,
			Form: map[string][]string{
				"user_id":  {"firstUser"},
				"password": {"example"},
			},
		},
		LoginInfo: postgres.ValidateUserParams{
			ID:       sql.NullString{String: "firstUser", Valid: true},
			Password: sql.NullString{String: "example", Valid: true},
		},
		LoginResult: loginHandlerTCsResult{
			Code:  200,
			Valid: true,
		},
	},
	{
		LoginRequest: &http.Request{
			Method: http.MethodGet,
			Form: map[string][]string{
				"user_id":  {"secondUser"},
				"password": {"example"},
			},
		},
		LoginInfo: postgres.ValidateUserParams{
			ID:       sql.NullString{String: "secondUser", Valid: true},
			Password: sql.NullString{String: "example", Valid: true},
		},
		LoginResult: loginHandlerTCsResult{
			Code:  401,
			Data:  `{"error": "incorrect user_id/pwd"}`,
			Valid: false,
		},
	},
}

func TestLoginHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockStore := mockdb.NewMockStore(mockCtrl)

	config, err := configurations.LoadConfig("../../resources")
	require.Nil(t, err)

	for i, tc := range loginHandlerTCs {
		fmt.Print("Running TestLoginHandler TC: ", i, ": ")
		mockStore.EXPECT().ValidateUser(gomock.Any(), tc.LoginInfo).Return(tc.LoginResult.Valid).Times(1)
		sc := mockServiceController(t, config, mockStore)
		resp := httptest.NewRecorder()

		require.NotNil(t, sc)

		sc.loginHandler(resp, tc.LoginRequest)
		require.Equal(t, tc.LoginResult.Code, resp.Code)
	}
}
