package initializer

import (
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/require"
)

var globalTestConfig GlobalParams = GlobalParams{
	Enviroments: &EnviromentParams{},
	Components:  &ComponentParams{},
}

func TestLoadEnviromemts(t *testing.T) {
	// Test cover reading .env.test file

	/* Change workspace dir to read .env.test */
	_, filename, _, _ := runtime.Caller(0)
	os.Chdir(path.Join(path.Dir(filename), ".."))

	err := globalTestConfig.LoadEnviromemts(".env-test")
	require.NoError(t, err)
	require.NotEmpty(t, globalTestConfig.Enviroments.DbHost)
	require.NotEmpty(t, globalTestConfig.Enviroments.DbName)
	require.NotEmpty(t, globalTestConfig.Enviroments.DbPassword)
	require.NotEmpty(t, globalTestConfig.Enviroments.DbUser)
	require.NotEmpty(t, globalTestConfig.Enviroments.DbPort)
}

func TestLoadComponents(t *testing.T) {
	var wait int
	var err error
	for wait = 0; wait < 50; wait++ {
		err = globalTestConfig.LoadComponents(globalTestConfig.Enviroments)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	require.NoError(t, err)
}
