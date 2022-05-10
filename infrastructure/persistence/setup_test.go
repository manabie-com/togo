package persistence

import (
	"fmt"
	"log"
	"os"

	"github.com/jfzam/togo/domain/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() (*gorm.DB, error) {
	if _, err := os.Stat("./../../.env"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("./../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return LocalDatabase()
	}
	return CIBuild()
}

//Circle CI DB
func CIBuild() (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "togo_test", "password")
	conn, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return conn, nil
}

//Local DB
func LocalDatabase() (*gorm.DB, error) {
	dbdriver := os.Getenv("TEST_DB_DRIVER")
	host := os.Getenv("TEST_DB_HOST")
	password := os.Getenv("TEST_DB_PASSWORD")
	user := os.Getenv("TEST_DB_USER")
	dbname := os.Getenv("TEST_DB_NAME")
	port := os.Getenv("TEST_DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	conn, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", dbdriver)
	}

	err = conn.Debug().Migrator().DropTable(&entity.User{}, &entity.Task{})
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		entity.User{},
		entity.Task{},
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func seedUser(db *gorm.DB) (*entity.User, error) {
	user := &entity.User{
		ID:              1,
		UserName:        "testuser1",
		Password:        "testpassword1",
		TaskLimitPerDay: 2,
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func seedUsers(db *gorm.DB) ([]entity.User, error) {
	users := []entity.User{
		{
			ID:              1,
			UserName:        "testuser1",
			Password:        "testpassword0001",
			TaskLimitPerDay: 1,
		},
		{
			ID:              2,
			UserName:        "testuser2",
			Password:        "testpassword2",
			TaskLimitPerDay: 2,
		},
	}
	for _, v := range users {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func seedTask(db *gorm.DB) (*entity.Task, error) {
	task := &entity.Task{
		ID:          1,
		Title:       "task title",
		Description: "task desc",
		UserID:      1,
	}
	err := db.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func seedTasks(db *gorm.DB) ([]entity.Task, error) {
	tasks := []entity.Task{
		{
			ID:          1,
			Title:       "1st task",
			Description: "first desc",
			UserID:      1,
		},
		{
			ID:          2,
			Title:       "2nd task",
			Description: "second desc",
			UserID:      1,
		},
	}
	for _, v := range tasks {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}
