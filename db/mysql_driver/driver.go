package mysql_driver

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"mini_project/db/model"

	"log"
	"os"
	"sync"
	"time"

	gorm_driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var lock = &sync.Mutex{}

type database struct {
	db *gorm.DB
	m  *sync.RWMutex
}

var singleInstanceDb model.DatabaseAPI

func GetConnection(dbUrl map[string]string) model.DatabaseAPI {
	if singleInstanceDb == nil {
		lock.Lock()
		defer lock.Unlock()
		//double check
		if singleInstanceDb == nil {
			mysqlDB, err := open(dbUrl)

			if err != nil {
				log.Fatalln("failed to connect database")
			}
			singleInstanceDb = &database{db: mysqlDB, m: &sync.RWMutex{}}
		}
	}
	return singleInstanceDb
}

func GetMigrator(dbUrl map[string]string) gorm.Migrator {
	mysqlDB, err := open(dbUrl)

	if err != nil {
		log.Fatalln("failed to connect database")
	}
	return mysqlDB.Migrator()
}

func open(dburl map[string]string) (*gorm.DB, error) {
	// url would look like user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", dburl["User"], dburl["Password"], dburl["Host"], dburl["Name"])
	sqlDB, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             1 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,     // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // enable color
		},
	)
	return gorm.Open(gorm_driver.New(gorm_driver.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		PrepareStmt:            true,
		Logger:                 newLogger,
		FullSaveAssociations:   true,
		SkipDefaultTransaction: true,
	})
}

func New(dbUrl map[string]string) {
	mysqlDB, err := open(dbUrl)

	if err != nil {
		panic("failed to connect database")
	}
	// create or update table
	mysqlDB.Set("gorm:table_options", "ENGINE=InnoDB COLLATE=utf8mb4_unicode_ci").AutoMigrate(
		&model.User{},
		&model.Task{},
	)

	db, err := mysqlDB.DB()
	if err != nil {
		panic("failed to connect database")
	}
	// Should config by user
	db.SetMaxOpenConns(500)
	//set max idle connection refers to https://github.com/jinzhu/gorm/issues/246
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Hour)

	singleInstanceDb = &database{
		db: mysqlDB,
		m:  &sync.RWMutex{},
	}
}

func (s *database) HasTable(table interface{}) bool {
	return s.db.Migrator().HasTable(table)
}

func (s *database) DropTable(table interface{}) {
	s.db.Migrator().DropTable(table)
}

// Close handles any necessary cleanup
func (s *database) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func ComputeFlowHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
