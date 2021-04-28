package driver

import (
	"fmt"
	"log"

	"github.com/manabie-com/togo/define"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type ConnectDB struct {
	Conn *gorm.DB
}

var Driver = &ConnectDB{}

func CreateConnect() (*ConnectDB, error) {
	var (
		dbConn string
		dbType = define.ConfigDatabase["type"]
		dbHost = define.ConfigDatabase["host"]
		dbPort = define.ConfigDatabase["port"]
		dbname = define.ConfigDatabase["name"]
		dbUser = define.ConfigDatabase["user"]
		dbPass = define.ConfigDatabase["pass"]
	)

	// create connection string with postgredb
	dbConn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbname)

	db, err := gorm.Open(dbType, dbConn)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	Driver.Conn = db
	return Driver, nil
}

// GetConnection func
func GetConnection() (*ConnectDB, error) {
	db, err := CreateConnect()
	if err != nil {
		return db, err
	}
	err = db.Conn.DB().Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

// CloseConnection func close connection db
func CloseConnection(db *ConnectDB) {
	err := db.Conn.Close()
	if err != nil {
		log.Fatal("error opening db", err)
	}
}
