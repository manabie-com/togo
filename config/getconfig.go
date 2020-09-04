package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/manabie-com/togo/internal/storages"
)

const (
	//POSTGRESQL to get env variable postgre
	POSTGRESQL = "postgres"
	//SQLITE to get env variable sqlite
	SQLITE = "sqlite"
)

// Postgres ..
type Postgres struct {
	// host = "localhost" when run go run main.go on local machine
	// host = "database" // when run in container
	// PostgreSQL container is running on local machine, so we need to connect with localhost. If it is running on a specific server, use your server IP.
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// Connect connecter interface
func (p *Postgres) Connect() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Password, p.Name)
	log.Println(psqlInfo)
	return sql.Open("postgres", psqlInfo)
}

// Sqlite ..
type Sqlite struct {
	DataPath string
}

// Connect connecter interface
func (l *Sqlite) Connect() (*sql.DB, error) {
	return sql.Open("sqlite3", l.DataPath)
}

// GetDBConnecter ...
func GetDBConnecter() (storages.DBConnecter, error) {
	dbSys := os.Getenv("DB_SYS")
	// log.Printf("DB Type: %v\n", dbSys)
	switch dbSys {
	case POSTGRESQL:
		log.Printf("DB Type: %v\n", dbSys)
		dbHost := os.Getenv("DB_HOST")

		dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		return &Postgres{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
		}, nil

	case SQLITE:
		log.Printf("DB Type: %v\n", dbSys)
		dbDataPath := os.Getenv("DB_DATA_PATH")

		return &Sqlite{
			DataPath: dbDataPath,
		}, nil
	default:
		log.Printf("DB Type: %v\n", dbSys)
		return nil, errors.New("Can not read config database")
	}
}
