package tests

import (
	"fmt"
	"testing"

	"github.com/manabie-com/togo/internal/data/dbtest"
	"github.com/manabie-com/togo/platform/docker"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}
