package configs

import (
	"bytes"
	_ "embed"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/vchitai/l"
)

var ll = l.New()

// DefaultConfig Provide a default config
//go:embed config.yaml
var DefaultConfig []byte

// Config holds all settings
type Config struct {
	Base                     `mapstructure:",squash"`
	MySQL                    *MySQL `yaml:"mysql" mapstructure:"mysql"`
	Redis                    *Redis `yaml:"redis" mapstructure:"redis"`
	ToDoListAddLimitedPerDay int64  `yaml:"default_todo_list_add_limited_per_day" mapstructure:"default_todo_list_add_limited_per_day"`
}

func (cfg *Config) Validate() error {
	return nil
}

// MySQL is settings of a MySQL server. It contains almost same fields as mysql.Config,
// but with some different field names and tags.
type MySQL struct {
	Username  string            `yaml:"username" mapstructure:"username"`
	Password  string            `yaml:"password" mapstructure:"password"`
	Protocol  string            `yaml:"protocol" mapstructure:"protocol"`
	Address   string            `yaml:"address" mapstructure:"address"`
	Database  string            `yaml:"database" mapstructure:"database"`
	Params    map[string]string `yaml:"params" mapstructure:"params"`
	Collation string            `yaml:"collation" mapstructure:"collation"`
	Loc       *time.Location    `yaml:"location" mapstructure:"loc"`
	TLSConfig string            `yaml:"tls_config" mapstructure:"tls_config"`

	Timeout      time.Duration `yaml:"timeout" mapstructure:"timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`

	AllowAllFiles           bool   `yaml:"allow_all_files" mapstructure:"allow_all_files"`
	AllowCleartextPasswords bool   `yaml:"allow_cleartext_passwords" mapstructure:"allow_cleartext_passwords"`
	AllowOldPasswords       bool   `yaml:"allow_old_passwords" mapstructure:"allow_old_passwords"`
	ClientFoundRows         bool   `yaml:"client_found_rows" mapstructure:"client_found_rows"`
	ColumnsWithAlias        bool   `yaml:"columns_with_alias" mapstructure:"columns_with_alias"`
	InterpolateParams       bool   `yaml:"interpolate_params" mapstructure:"interpolate_params"`
	MultiStatements         bool   `yaml:"multi_statements" mapstructure:"multi_statements"`
	ParseTime               bool   `yaml:"parse_time" mapstructure:"parse_time"`
	GoogleAuthFile          string `yaml:"google_auth_file" mapstructure:"google_auth_file"`
}

// FormatDSN returns MySQL DSN from settings.
func (m *MySQL) FormatDSN() string {
	um := &mysql.Config{
		User:                    m.Username,
		Passwd:                  m.Password,
		Net:                     m.Protocol,
		Addr:                    m.Address,
		DBName:                  m.Database,
		Params:                  m.Params,
		Collation:               m.Collation,
		Loc:                     m.Loc,
		TLSConfig:               m.TLSConfig,
		Timeout:                 m.Timeout,
		ReadTimeout:             m.ReadTimeout,
		WriteTimeout:            m.WriteTimeout,
		AllowAllFiles:           m.AllowAllFiles,
		AllowCleartextPasswords: m.AllowCleartextPasswords,
		AllowOldPasswords:       m.AllowOldPasswords,
		ClientFoundRows:         m.ClientFoundRows,
		ColumnsWithAlias:        m.ColumnsWithAlias,
		InterpolateParams:       m.InterpolateParams,
		MultiStatements:         m.MultiStatements,
		ParseTime:               m.ParseTime,
		AllowNativePasswords:    true,
	}
	return um.FormatDSN()
}

// Redis ...
type Redis struct {
	Addr     string `yaml:"addr" mapstructure:"addr"`
	Password string `yaml:"password" mapstructure:"password"`
	DB       int    `yaml:"db" mapstructure:"db"`
}

type Base struct {
	HTTPAddress int `yaml:"http_address" mapstructure:"http_address"`
	GRPCAddress int `mapstructure:"grpc_address"`
	Environment string
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(DefaultConfig))
	if err != nil {
		ll.Fatal("Failed to read viper config", l.Error(err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		ll.Fatal("Failed to unmarshal config", l.Error(err))
	}
	if err := cfg.Validate(); err != nil {
		ll.Fatal("Failed to validate config", l.Error(err))
	}

	ll.Info("Config loaded", l.Object("config", cfg))
	return cfg
}

// MySQLDSN returns the MySQL DSN from config.
func (cfg *Config) MySQLDSN() string {
	return cfg.MySQL.FormatDSN()
}
