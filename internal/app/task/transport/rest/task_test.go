package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/app/task/transport/rest/mock"
	"github.com/manabie-com/togo/internal/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDelivery_RetrieveTasks(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		mockReq         func(endpoint string) (*http.Request, error)
		mockTaskService func() TaskService
		wantStatus      int
		wantData        string
	}{
		{
			name: "created_date is empty",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("created_date", "")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockTaskService: func() TaskService {
				return nil
			},
			wantStatus: http.StatusBadRequest,
			wantData:   "{\"error\":\"created_date is missing\"}\n",
		},
		{
			name: "created_date is invalid",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("created_date", "123")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockTaskService: func() TaskService {
				return nil
			},
			wantStatus: http.StatusBadRequest,
			wantData:   "{\"error\":\"created_date is invalid\"}\n",
		},
		{
			name: "unable to retrieve tasks",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("created_date", "2020-06-29")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockTaskService: func() TaskService {
				ts := mock.NewMockTaskService(gomock.NewController(t))
				ts.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error")).AnyTimes()
				return ts
			},
			wantStatus: http.StatusInternalServerError,
			wantData:   "{\"error\":\"mock error\"}\n",
		},
		{
			name: "success",
			mockReq: func(endpoint string) (*http.Request, error) {
				r, err := http.NewRequest(http.MethodGet, endpoint, nil)
				if err != nil {
					return nil, err
				}
				q := r.URL.Query()
				q.Add("created_date", "2020-06-29")
				r.URL.RawQuery = q.Encode()
				return r, nil
			},
			mockTaskService: func() TaskService {
				ts := mock.NewMockTaskService(gomock.NewController(t))
				ts.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any()).Return([]model.Task{
					{
						ID:          "ID",
						Content:     "Content",
						UserID:      "UserID",
						CreatedDate: "2020-06-29",
					},
				}, nil).AnyTimes()
				return ts
			},
			wantStatus: http.StatusOK,
			wantData:   "{\"data\":[{\"id\":\"ID\",\"content\":\"Content\",\"user_id\":\"UserID\",\"created_date\":\"2020-06-29\"}]}\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Delivery{
				taskService: tt.mockTaskService(),
				restUtil:    util.NewRestUtil(),
			}
			ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				d.RetrieveTasks(writer, request)
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

func TestDelivery_AddTask(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		mockReq         func(endpoint string) (*http.Request, error)
		mockMiddleware  func(req *http.Request)
		mockTaskService func() TaskService
		wantStatus      int
	}{
		{
			name: "invalid JSON data",
			mockReq: func(endpoint string) (*http.Request, error) {
				var b []byte
				body := bytes.NewBuffer(b)
				body.Write([]byte("invalid_json"))
				r, err := http.NewRequest(http.MethodPost, endpoint, body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			mockTaskService: func() TaskService {
				return nil
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "unable to add task",
			mockReq: func(endpoint string) (*http.Request, error) {
				var b []byte
				body := bytes.NewBuffer(b)
				d, err := json.Marshal(map[string]string{
					"content": "something like this",
				})
				if err != nil {
					return nil, err
				}
				body.Write(d)
				r, err := http.NewRequest(http.MethodPost, endpoint, body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			mockMiddleware: func(req *http.Request) {
				*req = *req.WithContext(util.SetUserIDToContext(req.Context(), "123"))
			},
			mockTaskService: func() TaskService {
				ts := mock.NewMockTaskService(gomock.NewController(t))
				ts.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(model.Task{}, errors.New("mock error")).AnyTimes()
				return ts
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			mockReq: func(endpoint string) (*http.Request, error) {
				var b []byte
				body := bytes.NewBuffer(b)
				d, err := json.Marshal(map[string]string{
					"content": "something like this",
				})
				if err != nil {
					return nil, err
				}
				body.Write(d)
				r, err := http.NewRequest(http.MethodPost, endpoint, body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			mockMiddleware: func(req *http.Request) {
				*req = *req.WithContext(util.SetUserIDToContext(req.Context(), "123"))
			},
			mockTaskService: func() TaskService {
				ts := mock.NewMockTaskService(gomock.NewController(t))
				ts.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(model.Task{}, nil).AnyTimes()
				return ts
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Delivery{
				taskService: tt.mockTaskService(),
				restUtil:    util.NewRestUtil(),
			}
			ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if tt.mockMiddleware != nil {
					tt.mockMiddleware(request)
				}
				d.AddTask(writer, request)
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
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status got %d, want %d", resp.StatusCode, tt.wantStatus)
				return
			}
		})
	}
}
