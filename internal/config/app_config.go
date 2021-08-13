package config

type Config struct {
	DB struct {
		Host     string
		Name     string
		User     string `default:"postgres"`
		Password string `required:"true" env:"DBPassword"`
		Port     string   `default:"2345"`
	}
}