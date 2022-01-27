package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Asset(name string) ([]byte, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	base := filepath.Join(rootPath, "core/config")
	if strings.Contains(name, "..") {
		panic(fmt.Sprintf("invalid name (%v)", name))
	}
	return ioutil.ReadFile(filepath.Join(base, name))
}

type DBConfig struct {
	PostgresDB PostgresConfig `json:"postgres_db"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}

type Config struct {
	Databases DBConfig `json:"databases"`
}
