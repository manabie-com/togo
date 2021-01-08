package postgresql_test

import (
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	"github.com/spf13/viper"
)

type TestHelper struct {
	Store *postgresql.Store
}

func Setup() (*TestHelper, error) {
	viper.SetConfigName("test.config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf := model.AppSettings{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	s, err := postgresql.NewStore(conf.DatabaseSettings)
	if err != nil {
		return nil, err
	}

	th := &TestHelper{
		Store: s,
	}

	return th, nil
}

func (th *TestHelper) Teardown() {
	th.Store.DropAllRecords()
}
