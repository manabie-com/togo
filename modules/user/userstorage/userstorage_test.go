package userstorage_test

import (
	"errors"
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/japananh/togo/modules/user/userstorage"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"regexp"
	"testing"
)

func loadEnv() error {
	re := regexp.MustCompile(`^(.*` + "manabie-interview-test" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return godotenv.Load(string(rootPath) + `/.env`)
}

func setUpTest(t *testing.T) (*gorm.DB, error) {
	err := loadEnv()
	require.Nil(t, err)
	db, err := gorm.Open(mysql.Open(os.Getenv("TEST_DB_CONNECTION_STR")), &gorm.Config{})
	require.Nil(t, err, "cannot open database connection")
	if db == nil {
		return nil, errors.New("gorm db is null")
	}
	return db, nil
}

func TestSqlStore_CreateUser(t *testing.T) {
	db, err := setUpTest(t)
	require.Nil(t, err)

	user := usermodel.UserCreate{
		Email:          "user@gmail.com",
		Password:       "b0dd9c5cfd02c3e96171ab3f08e67dac",
		Salt:           "BMOcrdlEltpGCZxZkmVyBqyDwxrDXkxPLZMOFDSXNxGqrwKoxt",
		DailyTaskLimit: 6,
	}

	err = userstorage.NewSQLStore(db).CreateUser(nil, &user)
	require.Nil(t, err, "failed to insert user to gorm db")
}

func TestSqlStore_Login(t *testing.T) {
	db, err := setUpTest(t)
	require.Nil(t, err)

	user, err := userstorage.NewSQLStore(db).FindUser(
		nil,
		map[string]interface{}{"email": "user@gmail.com", "password": "b0dd9c5cfd02c3e96171ab3f08e67dac"},
	)
	require.Nil(t, err)
	require.NotNil(t, user)
}
