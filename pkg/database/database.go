package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

const (
	// DBMS postgres type
	DBMS = "postgres"
)

// Database struct
type Database struct {
	*gorm.DB
	Connection *sql.DB
}

type DBConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

func NewDatabase(dbConfig DBConfig) (Database, error) {
	config := "host=" + dbConfig.Host + " port=" + dbConfig.Port + " user=" + dbConfig.User + " dbname=" + dbConfig.Name + " sslmode=disable password=" + dbConfig.Pass
	conn, err := sql.Open(DBMS, config)
	if err != nil {
		return Database{}, err
	}
	dialector := postgres.New(postgres.Config{Conn: conn})
	if err != nil {
		return Database{}, err
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})

	return Database{db, conn}, err
}

func Close(db Database) error {
	return db.Connection.Close()
}
