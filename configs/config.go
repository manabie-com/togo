package configs

type Config struct {
	Server     *ServerConfig     `mapstructure:"SERVER"`
	PostgreSQL *PostgreSQLConfig `mapstructure:"POSTGRESSQL"`
}

type ServerConfig struct {
	Port     string `mapstructure:"PORT"`
	Timezone string `mapstructure:"TZ"`
}

type PostgreSQLConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
	SslMode  string `mapstructure:"SSL_MODE"`
}
