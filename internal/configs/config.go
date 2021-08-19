package configs

type Config struct {
	JWT struct {
		Key string
	}
	DB struct {
		Host     string
		Name     string
		User     string
		Password string
		Port     uint
	}
}
