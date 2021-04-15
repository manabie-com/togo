package rest

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/app/user/transport/rest/mock"
	"github.com/manabie-com/togo/internal/app/user/usecase"
	"github.com/manabie-com/togo/internal/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDelivery_Login(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		mockReq     func(endpoint string) (*http.Request, error)
		mockAuthSrv func() AuthService
		wantStatus  int
		wantData    string
	}{
		{
			name: "user ID is empty",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("user_id", "")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockAuthSrv: func() AuthService {
				return nil
			},
			wantStatus: http.StatusBadRequest,
			wantData:   "{\"error\":\"user_id is missing\"}\n",
		},
		{
			name: "password is empty",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("user_id", "123")
				q.Add("password", "")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockAuthSrv: func() AuthService {
				return nil
			},
			wantStatus: http.StatusBadRequest,
			wantData:   "{\"error\":\"user_id is missing\"}\n",
		},
		{
			name: "invalid user",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("user_id", "123")
				q.Add("password", "password")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockAuthSrv: func() AuthService {
				as := mock.NewMockAuthService(gomock.NewController(t))
				as.EXPECT().GetAuthToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("", usecase.ErrInvalidUser).AnyTimes()
				return as
			},
			wantStatus: http.StatusUnauthorized,
			wantData:   fmt.Sprintf("{\"error\":\"%s\"}\n", usecase.ErrInvalidUser),
		},
		{
			name: "unable to create token",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("user_id", "123")
				q.Add("password", "password")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockAuthSrv: func() AuthService {
				as := mock.NewMockAuthService(gomock.NewController(t))
				as.EXPECT().GetAuthToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("", usecase.ErrUnableToGenerateToken).AnyTimes()
				return as
			},
			wantStatus: http.StatusInternalServerError,
			wantData:   fmt.Sprintf("{\"error\":\"%s\"}\n", usecase.ErrUnableToGenerateToken),
		},
		{
			name: "success",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("user_id", "123")
				q.Add("password", "password")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockAuthSrv: func() AuthService {
				as := mock.NewMockAuthService(gomock.NewController(t))
				as.EXPECT().GetAuthToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("token", nil).AnyTimes()
				return as
			},
			wantStatus: http.StatusOK,
			wantData:   fmt.Sprintf("{\"data\":\"%s\"}\n", "token"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Delivery{
				authService: tt.mockAuthSrv(),
				restUtil:    util.NewRestUtil(),
			}
			ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				d.Login(writer, request)
			}))
			defer ts.Close()

			req, err := tt.mockReq(ts.URL)
			if err != nil {
				t.Error(err)
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Error(err)
				return
			}
			defer resp.Body.Close()
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
				return
			}
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status got %d, want %d", resp.StatusCode, tt.wantStatus)
				return
			}
			if string(data) != tt.wantData {
				t.Errorf("data got %s, want %s", data, tt.wantData)
				return
			}
		})
	}
}

func TestDelivery_isInIgnoreList(t *testing.T) {
	t.Parallel()
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "login",
			args: args{
				path: "/login",
			},
			want: true,
		},
		{
			name: "tasks",
			args: args{
				path: "/tasks",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Delivery{}
			if got := d.isInIgnoreList(tt.args.path); got != tt.want {
				t.Errorf("isInIgnoreList() = %v, want %v", got, tt.want)
			}
		})
	}
}
