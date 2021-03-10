package sqllite

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"log"
)

// LiteDB for working with sqllite
var db *sql.DB


func InitSqlLiteRepository(dataSource string) {
	sqlDB, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	db = sqlDB
	userRepository := &userRepository{}
	taskRepository := &taskRepository{}
	storages.SetUserRepository(userRepository)
	storages.SetTaskRepository(taskRepository)

}
