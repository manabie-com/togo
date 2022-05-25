package common

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	appPort             int
	atExpiry            int
	rtExpiry            int
	secretKey           string
	appEnv              string
	dbConnectionStr     string
	dbConnectionStrTest string
}

func NewConfig() *config {
	return &config{}
}

// Load load config from `.env` file
func (c *config) Load(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	// By default, `viper` set configName = "config",
	// so it will look for `config.env` instead of `.env`
	// Reference: https://github.com/spf13/viper/blob/master/viper.go#L231
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("reading ---", err, path)
		return err
	}

	c.appPort = viper.GetInt("PORT")
	c.appEnv = viper.GetString("APP_ENV")
	c.secretKey = viper.GetString("SYSTEM_KEY")
	c.atExpiry = viper.GetInt("ACCESS_TOKEN_EXPIRY")
	c.rtExpiry = viper.GetInt("REFRESH_TOKEN_EXPIRY")
	c.dbConnectionStr = viper.GetString("DB_CONNECTION_STR")
	c.dbConnectionStrTest = viper.GetString("DB_CONNECTION_STR_TEST")

	return nil
}

func (c *config) DBConnectionURL() string {
	return c.dbConnectionStr
}

func (c *config) DBConnectionURLTest() string {
	return c.dbConnectionStrTest
}

func (c *config) AppPort() int {
	return c.appPort
}

func (c *config) AppEnv() string {
	return c.appEnv
}

func (c *config) SecretKey() string {
	return c.secretKey
}

func (c *config) RtExpiry() int {
	return c.rtExpiry
}

func (c *config) AtExpiry() int {
	return c.atExpiry
}
