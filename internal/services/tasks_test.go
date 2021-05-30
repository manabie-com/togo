package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestToDoService_ServeHTTP_login(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name        string
		args        args
		prepareMock func(mock sqlmock.Sqlmock)
		assert      func(t *testing.T, resp *http.Response)
	}{
		{
			name: "when provide correct username and password - return token",
			args: args{
				targetUrl: "/login?user_id=someUser&password=somePassword",
			},
			prepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE id = ? AND password = ?`)).
					WithArgs("someUser", "somePassword").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("someUser"))

			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
				var body map[string]string
				assert.Nil(t, json.NewDecoder(resp.Body).Decode(&body))
				assert.Contains(t, body, "data")
			},
		},
		{
			name: "when provide wrong username or password - return 400",
			args: args{
				targetUrl: "/login?user_id=someUser&password=somePassword",
			},
			prepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE id = ? AND password = ?`)).
					WithArgs("someUser", "somePassword").
					WillReturnError(sql.ErrNoRows)
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
				var body map[string]string
				assert.Nil(t, json.NewDecoder(resp.Body).Decode(&body))
				assert.Contains(t, body, "error")
				assert.Equal(t, "incorrect user_id/pwd", body["error"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			s := &ToDoService{
				JWTKey: "someToken",
				Store:  &sqllite.LiteDB{DB: db},
			}

			tt.prepareMock(mock)

			req := httptest.NewRequest("GET", tt.args.targetUrl, nil)
			w := httptest.NewRecorder()
			s.ServeHTTP(w, req)
			resp := w.Result()
			tt.assert(t, resp)
		})
	}
}
func TestToDoService_ServeHTTP_listTasks(t *testing.T) {
	type args struct {
		targetUrl string
		userId    string
	}
	tests := []struct {
		name      string
		args      args
		mockToken func(generatedToken string) string

		prepareMock func(mock sqlmock.Sqlmock)
		assert      func(t *testing.T, resp *http.Response)
	}{
		{
			name: "when provide token and created_date - return tasks for that date",
			args: args{
				targetUrl: "/tasks?created_date=2021-05-05",
				userId:    "someUser",
			},
			mockToken: func(generatedToken string) string {
				return generatedToken
			},
			prepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`)).
					WithArgs("someUser", "2021-05-05").
					WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).AddRow(1, "content", "someUser", "2021-05-05"))
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.JSONEq(t, `{"data":[{"id":"1","content":"content","user_id":"someUser","created_date":"2021-05-05"}]}`,
					string(body))
				println(string(body))
			},
		},
		{
			name: "when db have error - return status code 500",
			args: args{
				targetUrl: "/tasks?created_date=2021-05-05",
				userId:    "someUser",
			},
			mockToken: func(generatedToken string) string {
				return generatedToken
			},
			prepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`)).
					WithArgs("someUser", "2021-05-05").
					WillReturnError(sql.ErrNoRows)
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.Contains(t, string(body), sql.ErrNoRows.Error())
			},
		},
		{
			name: "when provide wrong token - return 401",
			args: args{
				targetUrl: "/tasks?created_date=2021-05-05",
				userId:    "someUser",
			},
			mockToken: func(generatedToken string) string {
				return ""
			},
			prepareMock: func(mock sqlmock.Sqlmock) {
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.Empty(t, body)
				println(string(body))

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.prepareMock(mock)
			s := &ToDoService{
				JWTKey: "someToken",
				Store:  &sqllite.LiteDB{DB: db},
			}
			token, err := s.createToken(tt.args.userId)
			req := httptest.NewRequest("GET", tt.args.targetUrl, nil)
			req.Header.Add("Authorization", tt.mockToken(token))
			w := httptest.NewRecorder()
			s.ServeHTTP(w, req)
			tt.assert(t, w.Result())

		})
	}
}

func TestToDoService_listTasks(t *testing.T) {
	type fields struct {
		JWTKey string
		Store  storages.DB
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		prepareMock func(mockDB *mocks.DB)
		assert      func(t *testing.T, resp *http.Response)
	}{
		{
			name:   "when s.Store.RetrieveTasks() return errors - response code 500",
			fields: fields{},
			args:   args{},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveTasks", mock.Anything, mock.Anything, mock.Anything).Return([]*storages.Task{}, errors.New("some error"))
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
		{
			name:   "when s.Store.RetrieveTasks() return tasks - response code 200 with tasks lists",
			fields: fields{},
			args:   args{},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveTasks", mock.Anything, mock.Anything, mock.Anything).Return([]*storages.Task{
					{
						ID:          "121-123-123",
						Content:     "someContent",
						UserID:      "someUserID",
						CreatedDate: "someCreatedDate",
					},
				}, nil)
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.JSONEq(t, `{"data":[{"id":"121-123-123","content":"someContent","user_id":"someUserID","created_date":"someCreatedDate"}]}`, string(body))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &mocks.DB{}
			tt.prepareMock(mockDB)
			s := &ToDoService{
				JWTKey: "someToken",
				Store:  mockDB,
			}
			req := httptest.NewRequest("GET", "/tasks", nil)
			w := httptest.NewRecorder()
			s.listTasks(w, req)
			tt.assert(t, w.Result())
		})
	}
}

func TestToDoService_addTask(t *testing.T) {
	type fields struct {
		JWTKey string
		Store  storages.DB
	}
	type args struct {
		reqBody string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		prepareMock func(mockDB *mocks.DB)
		assert      func(t *testing.T, resp *http.Response)
	}{
		{
			name:   "when s.Store.RetrieveUser() return errors - response code 500",
			fields: fields{},
			args:   args{},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveUser", mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
		{
			name:   "when s.Store.CountTasks() return errors - response code 500",
			fields: fields{},
			args:   args{},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveUser", mock.Anything, mock.Anything, mock.Anything).Return(
					&storages.User{
						ID:      "someUser",
						MaxTodo: 5,
					}, nil)

				mockDB.On("CountTasks", mock.Anything, mock.Anything, mock.Anything).Return(sql.NullInt32{}, errors.New("some error"))
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.Contains(t, string(body), "some error")
			},
		},
		{
			name:   "when tasks count > user.MaxTodo - response code 400",
			fields: fields{},
			args:   args{},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveUser", mock.Anything, mock.Anything, mock.Anything).Return(
					&storages.User{
						ID:      "someUser",
						MaxTodo: 5,
					}, nil)

				mockDB.On("CountTasks", mock.Anything, mock.Anything, mock.Anything).Return(sql.NullInt32{Int32: 5, Valid: true}, nil)
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.Contains(t, string(body), "Limited to 5 tasks per day")
			},
		},
		{
			name:   "when parse request body to json - response code 500",
			fields: fields{},
			args: args{
				reqBody: "{some invalid json}",
			},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveUser", mock.Anything, mock.Anything, mock.Anything).Return(
					&storages.User{
						ID:      "someUser",
						MaxTodo: 5,
					}, nil)

				mockDB.On("CountTasks", mock.Anything, mock.Anything, mock.Anything).Return(sql.NullInt32{Int32: 4, Valid: true}, nil)
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.Contains(t, string(body), `invalid character 's' looking for beginning of object key string`)
			},
		},
		{
			name:   "when Store.AddTask() return error - response code 400",
			fields: fields{},
			args: args{
				reqBody: `{"content":"someContent"}`,
			},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveUser", mock.Anything, mock.Anything, mock.Anything).Return(
					&storages.User{
						ID:      "someUser",
						MaxTodo: 5,
					}, nil)

				mockDB.On("CountTasks", mock.Anything, mock.Anything, mock.Anything).Return(sql.NullInt32{Int32: 4, Valid: true}, nil)
				mockDB.On("AddTask", mock.Anything, mock.Anything).Return(errors.New("some error"))
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
				body, _ := ioutil.ReadAll(resp.Body)
				assert.Contains(t, string(body), "some error")
			},
		},
		{
			name:   "when add task successfully - response status code 200",
			fields: fields{},
			args: args{
				reqBody: `{"content": "someContent"}`,
			},
			prepareMock: func(mockDB *mocks.DB) {
				mockDB.On("RetrieveUser", mock.Anything, mock.Anything, mock.Anything).Return(
					&storages.User{
						ID:      "someUser",
						MaxTodo: 5,
					}, nil)

				mockDB.On("CountTasks", mock.Anything, mock.Anything, mock.Anything).Return(sql.NullInt32{Int32: 4, Valid: true}, nil)
				mockDB.On("AddTask", mock.Anything, mock.Anything).Return(nil)
			},
			assert: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
				var body map[string]storages.Task
				assert.Nil(t, json.NewDecoder(resp.Body).Decode(&body))
				assert.NotEmpty(t, body["data"])
				assert.NotEmpty(t, body["data"].ID)
				assert.NotEmpty(t, body["data"].CreatedDate)
				assert.NotEmpty(t, body["data"].UserID)
				assert.Equal(t, body["data"].Content, "someContent")

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &mocks.DB{}
			tt.prepareMock(mockDB)
			s := &ToDoService{
				JWTKey: "someToken",
				Store:  mockDB,
			}
			req := httptest.NewRequest("POST", "/tasks", strings.NewReader(tt.args.reqBody))
			req = req.Clone(context.WithValue(req.Context(), userAuthKey(0), "someUser"))
			w := httptest.NewRecorder()
			s.addTask(w, req)
			tt.assert(t, w.Result())
		})
	}
}
