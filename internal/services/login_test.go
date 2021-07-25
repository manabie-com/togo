package services

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/surw/togo/internal/storages/mock"
)

func TestLogin(t *testing.T) {
	randomUserID := gofakeit.Username()
	randomPasswd := gofakeit.PetName()

	testCases := []struct {
		name          string
		userID        string
		passwd        string
		gotToken      bool
		buildStubs    func(store *mockdb.MockILiteDB)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			userID:   randomUserID,
			passwd:   randomPasswd,
			gotToken: true,
			buildStubs: func(store *mockdb.MockILiteDB) {
				store.EXPECT().ValidateUser(gomock.Any(), gomock.Eq(randomUserID), gomock.Eq(randomPasswd)).
					Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:     "DB err",
			userID:   randomUserID,
			passwd:   randomPasswd,
			gotToken: true,
			buildStubs: func(store *mockdb.MockILiteDB) {
				store.EXPECT().ValidateUser(gomock.Any(), gomock.Eq(randomUserID), gomock.Eq(randomPasswd)).
					Times(1).Return(errors.New("some err"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockILiteDB(ctrl)
			testCase.buildStubs(store)

			server := NewToDoService(store)

			recorder := httptest.NewRecorder()
			path := "/login"
			req, err := http.NewRequest(http.MethodGet, path, nil)
			require.NoError(t, err)
			q := url.Values{}
			q.Set("user_id", testCase.userID)
			q.Set("password", testCase.passwd)
			req.URL.RawQuery = q.Encode()

			var router = NewRouter()
			server.Register(router)

			router.router.ServeHTTP(recorder, req)
			testCase.checkResponse(t, recorder)
		})
	}
}
