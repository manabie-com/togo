package auth

import (
	"errors"
	"fmt"
	"github.com/jmsemira/togo/internal/database"
	"github.com/jmsemira/togo/internal/models"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
	"time"
)

type AuthTestSuite struct {
	suite.Suite
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
			user.Password = HashPass(user.Username)

			user.RateLimitPerDay = i
			db.Create(&user)
		}
	}
}

func (suite *AuthTestSuite) TestHashPass() {
	var tests = []struct {
		password string
		expected error
	}{
		{
			password: "password",
			expected: nil,
		},
	}

	for _, test := range tests {
		hashedPass := HashPass(test.password)

		err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(test.password+SALT))

		suite.Equal(err, test.expected)
	}
}

func (suite *AuthTestSuite) TestLogin() {
	var tests = []struct {
		username      string
		password      string
		expectedError error
	}{
		{
			username:      "user1",
			password:      "user1",
			expectedError: nil,
		},
		{
			username:      "user1",
			password:      "user2",
			expectedError: errors.New("Invalid username and password"),
		},
	}

	for _, test := range tests {
		user, err := Login(test.username, test.password)

		db := database.GetDB()
		expectedUser := models.User{}
		db.Where("username = ?", test.username).First(&expectedUser)

		suite.Equal(err, test.expectedError)
		if err == nil {
			suite.Equal(*user, expectedUser)
		}
	}
}

func (suite *AuthTestSuite) TestRegister() {
	var tests = []struct {
		input         models.User
		expectedError error
	}{
		{
			input: models.User{
				Username: "user4",
				Password: "user4",
			},
			expectedError: nil,
		},
		{
			input: models.User{
				Username: "user4",
				Password: "user4",
			},
			expectedError: errors.New("Username Exist!"),
		},
	}

	for _, test := range tests {
		err := Register(&test.input)
		suite.Equal(err, test.expectedError)
	}
}

func (suite *AuthTestSuite) TestGenerateJWTToken() {
	var tests = []struct {
		input         models.User
		expected      string
		expectedError error
	}{
		{
			input: models.User{
				ID:              1,
				Username:        "user4",
				Password:        "user4",
				RateLimitPerDay: 1,
			},
			expected:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1c2VyNCIsIlJhdGVMaW1pdCI6MSwiZXhwIjoxNjUxNzgyMDIwfQ.NqZQ_jggYb_A8JMzciE6DepL8sX1djkoXkkD8gH4xBo",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		token, err := GenerateJWTToken(&test.input)
		suite.Equal(err, test.expectedError)
		suite.Equal(test.expected, token)
	}
}

func TestAuthSuite(t *testing.T) {
	GenerateExpirationTime = func() time.Time {
		timeLayout := "2006-01-02 15:04:05"
		strTime := "2022-05-05 20:20:20"

		timeT, _ := time.Parse(timeLayout, strTime)

		return timeT
	}
	initializeDatabase()
	suite.Run(t, new(AuthTestSuite))
	os.Remove("test.db")
}
