package taskstorage_test

import (
	"errors"
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/japananh/togo/modules/task/taskstorage"
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

// TODO: remove loadEnv when integration test finish
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

func TestSqlStore_CreateTask(t *testing.T) {
	db, err := setUpTest(t)
	require.Nil(t, err)

	task := taskmodel.TaskCreate{
		Title:       "task 1",
		Description: "description 1",
		CreatedBy:   1,
	}

	err = taskstorage.NewSQLStore(db).CreateTask(nil, &task)
	require.Nil(t, err, "failed to insert task to gorm db")
}

func TestSqlStore_FindTaskByCondition(t *testing.T) {
	db, err := setUpTest(t)
	require.Nil(t, err)

	task, err := taskstorage.NewSQLStore(db).FindTaskByCondition(nil, map[string]interface{}{"id": 1})
	require.Nil(t, err)
	require.NotNil(t, task)
}
