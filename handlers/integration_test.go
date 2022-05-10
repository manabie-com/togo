package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/models"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	ginRouter     *gin.Engine
	testUserModel *models.UserModel
	testTaskModel *models.TaskModel

	/* Testing User variables */
	testUserVar []models.User
)

func TestMain(m *testing.M) {
	/* Testing models */
	var db *gorm.DB
	/* First load config from .env.test */
	/* Change workspace dir to read .env */
	_, filename, _, _ := runtime.Caller(0)
	os.Chdir(path.Join(path.Dir(filename), ".."))

	wd, _ := os.Getwd()
	err := godotenv.Load(filepath.Join(wd, ".env.test"))
	if err != nil {
		log.Fatal("Cannot load .env.test file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	if db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("[MODEL TESTING] Cannot connect to database")
	}

	/* Database migration */
	// Drop all tables before testing
	db.Migrator().DropTable(&models.Task{}, &models.User{})

	// Migrating database
	if err = db.AutoMigrate(
		&models.User{},
		&models.Task{},
	); err != nil {
		log.Fatal("[MODEL TESTING] Failed when migrate database")
	}

	/* Init models */
	testUserModel = models.NewUserModel(db)
	testTaskModel = models.NewTaskModel(db)

	/* Init handlers */
	testHandlers := &Handlers{
		Task: NewTaskHandler(&models.Models{
			User: testUserModel,
			Task: testTaskModel,
		}),
	}
	/* Init testing variables */
	testUserVar = append(testUserVar,
		models.User{UserID: "User_ID#10", DailyTasksLimit: 8, MaxDailyTasks: 8},   // Test normal without reaching limit
		models.User{UserID: "User_ID#11", DailyTasksLimit: 10, MaxDailyTasks: 10}, // Test normal limit in a day
		models.User{UserID: "User_ID#12", DailyTasksLimit: 16, MaxDailyTasks: 16}, // Test daily limit with day time increase
	)
	gin.SetMode(gin.TestMode)
	ginRouter = gin.Default()
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(gin.Logger())
	apiGroup := ginRouter.Group("/api")
	{
		apiGroup.PUT("/tasks", testHandlers.Task.UpdateUserTask)
	}
	go ginRouter.Run(":8080")
	time.Sleep(time.Second * 2)
	os.Exit(m.Run())
}

func SendRequestWithBody(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	return w
}
func TestUpdateUserTask(t *testing.T) {
	var i int
	var reqBody []byte
	/* Setup UserTask testing variables */
	var testUserTask []models.UserTask
	for _, v := range testUserVar {
		testUserTask = append(testUserTask, models.UserTask{
			UserID:        v.UserID,
			MaxDailyLimit: v.MaxDailyTasks,
			TodoTask:      fmt.Sprintf("Tesing Todo task content for UserID %s", v.UserID),
		})
	}
	var err error
	// Case 1: -Normal case- Send ToDo task for a UserID in a day without reaching Daily tasks limit value
	for i = 0; i < int(testUserTask[0].MaxDailyLimit)-rand.Intn(int(testUserTask[0].MaxDailyLimit)-1); i++ {
		reqBody, err = json.Marshal(testUserTask[0])
		require.NoError(t, err, "Cannot encoding JSON")
		w := SendRequestWithBody(ginRouter, "PUT", "/api/tasks", string(reqBody))
		require.Equal(t, http.StatusOK, w.Code)
	}

	// Case 2: -Normal case with Limit- Send ToDo task for UserID in a day with out of Daily tasks limit value
	for i = 0; i < int(testUserTask[1].MaxDailyLimit)+rand.Intn(10); i++ {
		reqBody, err = json.Marshal(testUserTask[1])
		require.NoError(t, err, "Cannot encoding JSON")
		w := SendRequestWithBody(ginRouter, "PUT", "/api/tasks", string(reqBody))
		if i < int(testUserTask[1].MaxDailyLimit) {
			require.Equal(t, http.StatusOK, w.Code, "Successfully send ToDo task")
		} else {
			require.Equal(t, http.StatusNotImplemented, w.Code, "User out of Daily tasks limit value")
		}
	}

	// Case 3: -Day increase case- Send ToDo task with out of Daily tasks limit, then increase day to test refull limit value
	// var tasks int = 0
	// for i = 0; i < int(testUserTask[1].MaxDailyLimit)+rand.Intn(100); i++ {
	// 	reqBody, err = json.Marshal(testUserTask[1])
	// 	require.NoError(t, err, "Cannot encoding JSON")
	// 	w := SendRequestWithBody(ginRouter, "PUT", "/api/tasks", string(reqBody))
	// 	if tasks < int(testUserTask[1].MaxDailyLimit) {
	// 		require.Equal(t, http.StatusOK, w.Code, "Successfully send ToDo task")
	// 		tasks++
	// 	} else {
	// 		require.Equal(t, http.StatusNotImplemented, w.Code, "User out of Daily tasks limit value")
	// 		// Reset i and increase a day
	// 		increaseSystemDay(rand.Intn(30))
	// 		tasks = 0
	// 	}
	// }
}

func increaseSystemDay(nextDay int) error {
	var err error
	if _, err = exec.LookPath("date"); err != nil {
		return err
	}
	now := time.Now()
	nextDate := time.Date(now.Year(), now.Month(), now.Day()+nextDay, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	nextDateFmt := nextDate.Format("12 Apr 2006 12:12:12")
	args := []string{"--set", nextDateFmt}
	if err = exec.Command("date", args...).Run(); err != nil {
		return err
	}
	return nil
}
