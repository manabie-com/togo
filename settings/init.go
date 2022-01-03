package settings

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var logger *logrus.Logger
var DbName string
var DbDialector gorm.Dialector
var DbDebug bool
var RestPort int

func init() {
	initLog()
	getEnv()
	initDb()
}

func initLog() {
	logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
}

func GetLogger() *logrus.Logger {
	return logger
}

func setDefaultEnv() {
	viper.SetDefault("DB_NAME", "togo.sqlite")
	viper.SetDefault("DB_DEBUG", false)

	viper.SetDefault("REST_PORT", 8080)
}

func getEnv() {
	setDefaultEnv()

	DbName = viper.GetString("DB_NAME")
	DbDebug = viper.GetBool("DB_DEBUG")
	RestPort = viper.GetInt("REST_PORT")
}

func initDb() {
	DbDialector = sqlite.Open(DbName)
}
