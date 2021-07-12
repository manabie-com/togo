package config

import (
	"time"

	"github.com/spf13/viper"
)

func init() {
	load()
}

var (
	JWTKey       string
	JWTExpiresIn time.Duration
	HTTPPort     int
)

func load() {
	v := viper.New()
	v.SetConfigType("json")
	v.AutomaticEnv()

	v.SetDefault("HTTP_PORT", 5050)
	v.SetDefault("JWT_KEY", "wqGyEBBfPK9w3Lxw")
	v.SetDefault("JWT_EXPIRES_IN", 15*time.Minute) // 15 minutes

	HTTPPort = v.GetInt("HTTP_PORT")
	JWTKey = v.GetString("JWT_KEY")
	JWTExpiresIn = v.GetDuration("JWT_EXPIRES_IN")

}
