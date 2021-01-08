package usecase_test

import (
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/spf13/viper"
)

type TestHelper struct {
	Usecase *usecase.Usecase
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

	a, err := usecase.New(conf)
	if err != nil {
		return nil, err
	}

	th := &TestHelper{
		Usecase: a,
		//Store:   a.Store,
	}

	return th, nil
}

func (th *TestHelper) Teardown() {
	th.Usecase.Store.DropAllRecords()
}
