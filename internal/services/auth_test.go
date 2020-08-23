package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phuwn/togo/internal/storages/mocks"
	"github.com/phuwn/togo/util"
	"github.com/stretchr/testify/mock"
)

var (
	validPasswordForExampleUserID = "password"
	jwtKey                        = "wqGyEBBfPK9w3Lxw"

	// token generated at 2020-08-20 20:34:58
	validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTc5NTY1OTgsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.jfKWtsLnc-nq0E4sC7WKBSGwI5MpgdaeuxF7zw_vbZQ"
)

func TestToDoService_getAuthToken(t *testing.T) {
	util.MockRuntimeFunc()

	store := &mocks.Store{}
	store.On("ValidateUser",
		mock.Anything,
		sql.NullString{exampleUserID, true},
		validPasswordForExampleUserID,
	).Return(true)
	store.On("RetrieveTasks", mock.Anything, mock.Anything, mock.Anything).Return(false)

	s := &ToDoService{
		JWTKey: jwtKey,
		Store:  store,
	}

	type args struct {
		userID   string
		password string
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			"happy case",
			args{exampleUserID, validPasswordForExampleUserID},
			http.StatusOK,
			fmt.Sprintf(`{"data":"%v"}`+"\n", validToken),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/login", nil)
			if err != nil {
				t.Errorf("failed to create request, err: %s\n", err.Error())
				return
			}
			q := req.URL.Query()
			q.Add("user_id", tt.args.userID)
			q.Add("password", tt.args.password)
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()
			s.getAuthToken(rr, req)
			if rr.Code != tt.code {
				t.Errorf("unexpected response: status code want %v, got %v, error message: %s", tt.code, rr.Code, rr.Body.String())
				return
			}

			if rr.Body.String() != tt.want {
				t.Errorf("unexpected output, got %s, want %s", rr.Body.String(), tt.want)
				return
			}
		})
	}
}
