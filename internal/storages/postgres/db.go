package postgres

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"log"
)

// PostgreDB for working with postgre
var db *sql.DB

func InitPostgreDBRepository(dataSource string) {
	sqlDB, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	db = sqlDB
	userRepository := &userRepository{}
	taskRepository := &taskRepository{}
	storages.SetUserRepository(userRepository)
	storages.SetTaskRepository(taskRepository)

}
