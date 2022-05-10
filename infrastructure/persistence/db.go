package persistence

import (
	"fmt"

	"github.com/jfzam/togo/domain/entity"
	"github.com/jfzam/togo/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	Task repository.TaskRepository
	db   *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		Task: NewTaskRepository(db),
		db:   db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	db, err := s.db.DB()
	db.Close()
	return err
}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.Task{})
}
