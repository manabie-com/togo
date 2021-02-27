package postgres

import "fmt"

type Config struct {
	Host string
	Port string
	Usr  string
	Pwd  string
	Db   string
}

func (c *Config) toConnStr() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.Usr, c.Pwd, c.Host, c.Port, c.Db)
}
