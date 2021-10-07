package postgres

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
	Timezone string
}

func NewDBConnection() *DBConnection {
	return &DBConnection{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetInt("postgres.port"),
		Username: viper.GetString("postgres.username"),
		Password: viper.GetString("postgres.password"),
		Database: viper.GetString("postgres.database"),
		SSLMode:  viper.GetString("postgres.ssl_mode"),
		Timezone: viper.GetString("postgres.time_zone"),
	}
}

func (c DBConnection) ToConnectionString() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		c.Host,
		c.Username,
		c.Password,
		c.Database,
		c.Port,
		c.SSLMode,
		c.Timezone)
}

func NewGormDB() (*gorm.DB, error) {
	c := NewDBConnection().ToConnectionString()
	return gorm.Open(postgres.Open(c))
}
