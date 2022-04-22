package dbcontext

import (
	"ManabieProject/env"
	"ManabieProject/src/dbcontrol/mongodb"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net"
)

const (
	MONGODB   = "mongodb"
	ORACLE    = "oracle"
	MYSQL     = "mysql"
	SQLSERVER = "sqlserver"
)

// DataContext interface Data Context
type DataContext interface {
	Connect(string) error
	Close()
	Ping() error
	InsertOne(string, string, interface{}) (interface{}, error)
	ReplaceOne(string, string, interface{}, interface{}) (interface{}, error)
	UpdateOne(string, string, interface{}, interface{}) (interface{}, error)
	InsertMany(string, string, []interface{}) ([]interface{}, error)
	DeleteOne(string, string, interface{}) (int64, error)
	DeleteMany(string, string, interface{}) (int64, error)
	Query(string, string, interface{}, interface{}) ([]bson.M, error)
}

var Context DataContext

func Init() error {
	host := env.Value("DB_HOST")
	if net.ParseIP(host) == nil {
		fmt.Print(host + " is not in the correct IP format !")
		return errors.New(host + " is not in the correct IP format !")
	}
	name := env.Value("DB_NAME")
	switch name {
	case MONGODB:
		Context = new(mongodb.MongoDB)
		// mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]
		var applyURI string
		if env.Value("DB_USERNAME") == "" {
			applyURI = name + ":" + "//" + env.Value("DB_HOST") + ":" + env.Value("DB_PORT")
		} else {
			applyURI = name + ":" + "//" + env.Value("DB_USERNAME") + ":" + env.Value("DB_PASSWORD") + env.Value("DB_HOST") + ":" + env.Value("DB_PORT")
		}
		err := Context.Connect(applyURI)
		return err
	case MYSQL:
		// todo
	case SQLSERVER:
		// todo
	case ORACLE:
		// todo
	default:
		fmt.Print(name + " is not supported !")
	}
	return nil
}

func Ping() error {
	err := Context.Ping()
	if err != nil {
		fmt.Println("error !")
	} else {
		fmt.Println("ok !")
	}
	return err
}
