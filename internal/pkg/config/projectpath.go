package config

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	root = filepath.Join(filepath.Dir(b), "../../..")
)

func RootPath() string {
	envRootPath := os.Getenv("ROOT_PATH")
	if envRootPath == "" {
		return root
	}

	return envRootPath
}
