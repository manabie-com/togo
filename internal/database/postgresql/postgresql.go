package postgresql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/manabie-com/togo/internal/env"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", getConnectionString())
	if err != nil {
		panic(err)
	}

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(env.MaxIdleConns())
	db.DB().SetMaxOpenConns(env.MaxOpenConns())
	db.SingularTable(false)
	db.LogMode(true)

	return db
}

func getConnectionString() string {
	host := env.GetDBHost()
	port := env.GetDBPort()
	userName := env.GetDBUserName()
	password := env.GetDBPassword()
	databaseName := env.GetDBName()
	sslMode := env.GetDBSSLMode()

	connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
	return fmt.Sprintf(connectionStringTemplate,
		host, port, sslMode,
		userName, password, databaseName)

}
