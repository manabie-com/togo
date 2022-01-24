package mysql

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"togo/common"
)

type mySQL struct {
	prefix string
	uri    string
	db     *gorm.DB
}

func NewMySQL(prefix string) *mySQL {
	return &mySQL{prefix: prefix}
}

func (s *mySQL) Get() interface{} {
	return s.db
}

func (s *mySQL) Run() error {
	s.initFlags()

	if err := s.configure(); err != nil {
		return err
	}

	db, err := gorm.Open(mysql.Open(s.uri))
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *mySQL) initFlags() {
	uri := os.Getenv("MYSQL_URI")
	s.uri = uri
}

func (s *mySQL) configure() error {
	if s.prefix == "" {
		return errors.New(common.MySQLUriNullErr)
	}

	if s.uri == "" {
		return errors.New(common.MySQLUriNullErr)
	}

	return nil
}

func (s *mySQL) GetPrefix() string{
	return s.prefix
}

func (s *mySQL) Stop() <-chan bool{
	stop := make(chan bool)
	go func() {
		stop <- true
	}()
	fmt.Printf("%v is stopped\n", s.GetPrefix())
	return stop
}