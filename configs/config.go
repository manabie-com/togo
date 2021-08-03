package configs

type Config struct {
	DBAddress    string
	RedisAddress string
	JwtKey       string

	Port int
}
