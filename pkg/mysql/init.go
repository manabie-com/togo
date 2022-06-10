package mysql

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db, err := gorm.Open(mysql.Open(s.uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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
		return errors.New(common.DataIsNullErr(s.prefix))
	}

	if s.uri == "" {
		return errors.New(common.DataIsNullErr("Mysql's URI"))
	}

	return nil
}

func (s *mySQL) GetPrefix() string {
	return s.prefix
}

func (s *mySQL) Stop() <-chan bool {
	stop := make(chan bool)
	go func() {
		stop <- true
	}()
	return stop
}
