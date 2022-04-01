package db

type SqlDbConfig struct {
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int    `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
}
