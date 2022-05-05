package api

import (
	"fmt"
	"github.com/jmsemira/togo/internal/auth"
	"github.com/jmsemira/togo/internal/database"
	"github.com/jmsemira/togo/internal/models"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type LoginTestSuite struct {
	suite.Suite
}

func (suite *LoginTestSuite) TestLoginHandler() {
	var tests = []struct {
		mockRequest *http.Request
		expected    string
	}{
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/login?username=user1&password=user1", nil)
				return req
			}(),
			expected: `
			{
				"status":"OK",
				"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1c2VyMSIsIlJhdGVMaW1pdCI6MSwiZXhwIjoxNjUxNzgyMDIwfQ.LVL1iDyzClpRyCJk616AJvvpyWYenBQNkgHitYhnn_k"
			}
			`,
		},
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/login?username=user1&password=user2", nil)
				return req
			}(),
			expected: `
			{
				"status":"error",
				"err_msg": "Invalid username and password"
			}
			`,
		},
	}

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		LoginHandler(recorder, test.mockRequest)

		// then
		suite.JSONEq(test.expected, recorder.Body.String())
	}
}

func initializeDatabase() {
	dbSettings := database.DBSettings{
		Type: "sqlite",
		Name: "test.db",
	}

	// initialize and migrate schema
	database.InitializeDB(dbSettings, &models.User{}, &models.Todo{})

	// check if the system already had a user
	users := []models.User{}

	db := database.GetDB()
	db.Find(&users)

	if len(users) == 0 {
		// initialize system users
		for _, i := range []int{1, 2, 3} {
			user := models.User{}
			user.Username = fmt.Sprintf("user%v", i)
			user.Password = auth.HashPass(user.Username)

			user.RateLimitPerDay = i
			db.Create(&user)
		}
	}
}

func TestLoginSuite(t *testing.T) {
	auth.GenerateExpirationTime = func() time.Time {
		timeLayout := "2006-01-02 15:04:05"
		strTime := "2022-05-05 20:20:20"

		timeT, _ := time.Parse(timeLayout, strTime)

		return timeT
	}
	initializeDatabase()
	suite.Run(t, new(LoginTestSuite))
}
