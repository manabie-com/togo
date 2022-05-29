package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dinhquockhanh/togo/internal/app/auth"
	"github.com/dinhquockhanh/togo/internal/app/user"
	"github.com/dinhquockhanh/togo/internal/pkg/token"
	"github.com/dinhquockhanh/togo/internal/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetByUserName(t *testing.T) {
	usrRandom := createMockUser()
	mockPayload := token.Payload{
		Username: usrRandom.Username,
		TierID:   int(usrRandom.TierID),
	}
	mockToken := "Bearer eyJhbGciOi"
	testCases := []struct {
		caseName      string
		username      string
		token         string
		buildStubs    func(us *user.MockService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			caseName: "OK",
			username: usrRandom.Username,
			token:    mockToken,
			buildStubs: func(us *user.MockService) {
				us.EXPECT().
					GetByUserName(gomock.Any(), gomock.Eq(usrRandom.Username)).
					Times(1).
					Return(usrRandom, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				bytes, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				var gotUser user.UserSafe
				err = json.Unmarshal(bytes, &gotUser)
				require.NoError(t, err)
				require.Equal(t, usrRandom.Safe(), &gotUser)
			},
		},
		//{
		//	caseName: "NotFound",
		//	username: usrRandom.Username,
		//	buildStubs: func(us *user.MockService) {
		//		us.EXPECT().
		//			GetByUserName(gomock.Any(), gomock.Eq(usrRandom.Username)).
		//			Times(1).
		//			Return(nil, sql.ErrNoRows)
		//	},
		//	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusNotFound, recorder.Code)
		//	},
		//},
		//{
		//	caseName: "Internal Error",
		//	username: usrRandom.Username,
		//	buildStubs: func(us *user.MockService) {
		//		us.EXPECT().
		//			GetByUserName(gomock.Any(), gomock.Eq(usrRandom.Username)).
		//			Times(1).
		//			Return(nil, errors.New("internal error"))
		//	},
		//	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusInternalServerError, recorder.Code)
		//	},
		//},
		//{
		//	caseName:   "Bad request",
		//	username:   "",
		//	buildStubs: func(us *user.MockService) {},
		//	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusBadRequest, recorder.Code)
		//	},
		//},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserSvc := user.NewMockService(ctrl)
	//mockTaskSvc := task.NewMockService(ctrl)
	//mockLimitSvc := limit.NewMockService(ctrl)
	mockUserSvc.EXPECT().
		GetByUserName(gomock.Any(), &user.GetUserByUserNameReq{UserName: usrRandom.Username}).
		Times(1).
		Return(usrRandom, nil)
	mockAuthSvc := auth.NewMockService(ctrl)
	mockTokenizer := token.NewMockTokenizer(ctrl)
	mockTokenizer.EXPECT().Verify(gomock.Any()).Return(&mockPayload, nil)
	h := Handler{
		//task: task.NewHandler(nil, mockUserSvc, mockLimitSvc),
		user: user.NewHandler(mockUserSvc),
		auth: auth.NewHandler(mockTokenizer, mockAuthSvc),
	}

	router, err := NewRouter(&h)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			url := fmt.Sprintf("/api/v1/users/%s", tc.username)
			// build stubs
			//tc.buildStubs(mockUserSvc)
			// start recorder
			recorder := httptest.NewRecorder()
			//send request
			request, err := http.NewRequest(http.MethodGet, url, nil)
			request.Header.Set("Authorization", tc.token)
			require.NoError(t, err)
			//start router
			router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})

	}

}

func createMockUser() *user.User {
	return &user.User{
		Username: util.RandomOwner(),
		Email:    util.RandomEmail(),
		TierID:   1,
	}
}
