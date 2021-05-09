package storages

import (
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type DatabaseTestSuite struct {
	suite.Suite
	cfg *config.Config
	db IDatabase
	gormDB *gorm.DB
}

func (s *DatabaseTestSuite) SetupSuite() {
	s.cfg = &config.Config{
		MaxTodo: 5,
		Environment: "T",
		SQLite:      "data.test.db",
	}
	os.Remove(s.cfg.SQLite)
	var err error
	s.gormDB, err = gorm.Open(sqlite.Open(s.cfg.SQLite),  &gorm.Config{})
	s.NoError(err)
	s.db, err = NewDatabase(s.cfg)
	s.NoError(err)
}

func (s *DatabaseTestSuite) TearDownSuite() {
	// remove file test db
	os.Remove(s.cfg.SQLite)
}

func (s *DatabaseTestSuite) Test_1_AddUser() {
	userId := "firstUser"
	password := "12345"
	maxTodo := int32(5)
	err := s.db.AddUser(userId, password, maxTodo)
	s.NoError(err)

	user := &models.User{}
	s.NoError(s.gormDB.Model(&models.User{}).Where("id = ?", userId).First(user).Error)

	s.Equal(userId, user.ID)
	s.Equal(password, user.Password)
	s.Equal(maxTodo, user.MaxTodo)
}

func (s *DatabaseTestSuite) Test_2_AddUser_UserExisted() {
	userId := "firstUser"
	password := "12345"
	maxTodo := s.cfg.MaxTodo
	err := s.db.AddUser(userId, password, maxTodo)
	s.Error(err)
}

func (s *DatabaseTestSuite) Test_3_AddTask() {
	id := uuid.New().String()
	userId := "firstUser"
	content := "task1"
	createdDate := "01-01-2021"
	err := s.db.AddTask(&models.Task{ID: id, UserID: userId, Content: content, CreatedDate: createdDate}, func(userId, createdDate string) error {
		return nil
	})
	s.NoError(err)

	task := &models.Task{}
	s.NoError(s.gormDB.Model(task).Where("id = ?", id).First(task).Error)
	s.Equal(id, task.ID)
	s.Equal(userId, task.UserID)
	s.Equal(content, task.Content)
	s.Equal(createdDate, task.CreatedDate)
}

func (s *DatabaseTestSuite) Test_4_CountTasks() {
	userId := "firstUser"
	createdDate := "01-01-2021"
	tasks, err := s.db.CountTasks(userId, createdDate)
	s.NoError(err)
	s.Equal(int32(1), tasks)
}

func (s *DatabaseTestSuite) Test_5_GetMaxTodo() {
	userId := "firstUser"
	maxTodo, err := s.db.GetMaxTodo(userId)
	s.NoError(err)
	s.Equal(s.cfg.MaxTodo, maxTodo)
}

func (s *DatabaseTestSuite) Test_6_RetrieveTasks() {
	userId := "firstUser"
	content := "task1"
	createdDate := "01-01-2021"
	tasks, err := s.db.RetrieveTasks(userId, createdDate)
	s.NoError(err)
	s.Len(tasks, 1)
	s.Equal(userId, tasks[0].UserID)
	s.Equal(content, tasks[0].Content)
	s.Equal(createdDate, tasks[0].CreatedDate)
}

func TestDatabase(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
