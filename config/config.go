package config

type MySQLConf struct {
	Host                   string `mapstructure:"host"`
	Port                   int    `mapstructure:"port"`
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	AuthenticationDatabase string `mapstructure:"authetication_database"`
}

type Config struct {
	MySQLConf MySQLConf `mapstructure:"mysql"`
}
