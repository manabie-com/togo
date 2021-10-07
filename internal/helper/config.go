package helper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func AutoBindConfig(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("%v is not a file", filePath)
	}
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")

	return viper.MergeInConfig()
}
