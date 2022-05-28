package connections

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	PostgresConnection = "host=%v user=%v dbname=%v sslmode=disable password=%v port=%v search_path=public"
)

var host, db, user, pass string
var port int

func init() {
	// Todo: implement get database config from secret manager
	if dbConfig, err := config.NewConfig("ini", "conf/db.conf"); err == nil {
		host = dbConfig.String("task::postgreshost")
		port, _ = dbConfig.Int("task::postgresport")
		db = dbConfig.String("task::postgresdb")
		user = dbConfig.String("task::postgresuser")
		pass = dbConfig.String("task::postgrespass")
	} else {
		panic(err)
	}
}
func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(createConnection()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
func createConnection() string {
	return fmt.Sprintf(PostgresConnection, host, user, db, pass, port)
}
