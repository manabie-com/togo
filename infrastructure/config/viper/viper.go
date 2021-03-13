package viper

import (
	"github.com/spf13/viper"

	"github.com/valonekowd/togo/infrastructure/config"
)

type ViperConfiger struct {
	v           *viper.Viper
	files       []*config.File
	readFromEnv bool
}

type Option func(*ViperConfiger)

func ConfigerFiles(files ...*config.File) Option {
	return func(vc *ViperConfiger) {
		vc.files = files
	}
}

func ConfigerReadFromEnv(want bool) Option {
	return func(vc *ViperConfiger) {
		vc.readFromEnv = want
	}
}

func NewConfiger(opts ...Option) config.Configer {
	vc := &ViperConfiger{
		v:           viper.New(),
		readFromEnv: false,
	}

	for _, o := range opts {
		o(vc)
	}

	return vc
}

func (vc *ViperConfiger) loadFromFiles() error {
	for i, f := range vc.files {
		vc.v.SetConfigName(f.Name)
		vc.v.SetConfigType(f.Type)
		for _, p := range f.Paths {
			vc.v.AddConfigPath(p)
		}

		fn := vc.v.ReadInConfig
		if i > 0 {
			fn = vc.v.MergeInConfig
		}

		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (vc *ViperConfiger) Load(val interface{}) error {
	err := vc.loadFromFiles()
	if err != nil {
		return err
	}

	if vc.readFromEnv {
		vc.v.AutomaticEnv()
	}

	return vc.v.Unmarshal(&val)
}

func (vc *ViperConfiger) SetDefault(key string, value interface{}) {
	vc.v.SetDefault(key, value)
}

func (vc *ViperConfiger) Set(key string, value interface{}) {
	vc.v.Set(key, value)
}
