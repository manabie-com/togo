package config

import (
	"flag"
	"fmt"
	"github.com/manabie-com/togo/pkg/common/sql"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

var (
	flConfigFile = ""
	flExample    = false
)

func InitFlags() {
	flag.StringVar(&flConfigFile, "config-file", "", "Path to config file")
	flag.BoolVar(&flExample, "example", false, "Print example config then exit")
}

func ParseFlags() {
	flag.Parse()
}

func LoadWithDefault(v, def interface{}) (err error) {
	defer func() {
		if flExample {
			if err != nil {
				log.Fatal("Error while loading config", err.Error())
			}
			PrintExample(v)
			os.Exit(2)
		}
	}()

	if flConfigFile != "" {
		err = LoadFromFile(flConfigFile, v)
		if err != nil {
			log.Fatalf("can not load config from file: %v (%v)", flConfigFile, err)
		}
		return err
	}
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(def))
	return nil
}

// LoadFromFile loads config from file
func LoadFromFile(configPath string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return LoadFromYaml(data, v)
}

func LoadFromYaml(input []byte, v interface{}) (err error) {
	return yaml.Unmarshal(input, v)
}

// PrintExample prints example config
func PrintExample(cfg interface{}) {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(data))
}

type Postgres = sql.ConfigPostgres

func DefaultPostgres() Postgres {
	return Postgres{
		Protocol: "postgres",
		Host:     "postgres",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "test",
		SSLMode:  "",
	}
}