package storages

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Config contains storages db configuration
type Config struct {
	Postgres PostgresConfig `json:"postgres"`
}

// PostgresConfig contains postgres db connStr parameters:
// https://pkg.go.dev/github.com/lib/pq
type PostgresConfig struct {
	// DbName - The name of the database to connect to
	DbName string `connStr:"dbname"`
	// User - The user to sign in as
	User string `connStr:"user"`
	// Password - the user's password
	Password string `connStr:"password"`
	// Host - the host to connect to. Values that start with / are for unix
	Host string `connStr:"host"`
	// Port - the port to bind to. (default is 5432)
	Port string `connStr:"port"`
	// ConnectionTimeout - Maximum wait for connection, in seconds. Zero or
	// not specified means wait indefinitely.
	ConnectionTimeout uint `connStr:"connect_timeout"`
	// SSLMode - Whether or not to use SSL (default is require, this is not
	// the default for libpq)
	SSLMode string `connStr:"sslmode"`
	// SSLCert - Cert file location. The file must contain PEM encoded data.
	SSLCert string `connStr:"sslcert"`
	// SSLKey - Key file location. The file must contain PEM encoded data.
	SSLKey string `connStr:"sslkey"`
	// SSLRootCert - The location of the root certificate file. The file
	// must contain PEM encoded data.
	SSLRootCert string `connStr:"sslrootcert"`

	// MaxIdleConns is to tune max idle conns using
	// (*sql.DB).SetMaxIdleConns(n int) (default is 2)
	MaxIdleConns int `connStr:"max_idle_conns"`
}

// Validate checks if required postgres configs are set
func (c PostgresConfig) Validate() error {
	var errstrings []string

	if len(c.DbName) == 0 {
		errstrings = append(errstrings, "DbName is not set")
	}

	if len(c.Host) == 0 {
		errstrings = append(errstrings, "Host is not set")
	}

	if len(c.User) == 0 {
		errstrings = append(errstrings, "User is not set")
	}

	if len(c.Password) == 0 {
		errstrings = append(errstrings, "Password is not set")
	}

	if len(c.SSLMode) == 0 {
		errstrings = append(errstrings, "SSLMode is not set")
	}

	if len(errstrings) > 0 {
		return errors.New(strings.Join(errstrings, " | "))
	}
	return nil
}

// Build return postgres config connStr
func (c PostgresConfig) Build() (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c.buildReflect(), nil
}

func (c PostgresConfig) buildReflect() string {
	var props []string
	val := reflect.ValueOf(c)

	for i := 0; i < val.NumField(); i++ {
		vField := val.Field(i).Interface()
		if !isZeroOfUnderlyingType(vField) {
			tField := val.Type().Field(i)
			tag := tField.Tag.Get("connStr")
			props = append(props, fmt.Sprintf("%v=%v", tag, vField))
		}
	}

	return strings.Join(props, " ")
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
