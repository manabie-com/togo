package common

import (
	"testing"
	"os"
	"github.com/stretchr/testify/require"
	"time"
)

func TestRepositoryFactorySql(t *testing.T) {
	prevOffset := os.Getenv("SIM_OFFSET_DAY")
	defer func () {
		os.Setenv("SIM_OFFSET_DAY", prevOffset)
	}()
	os.Setenv("SIM_OFFSET_DAY", "1")
	clock := MakeClockSim()
	simTime := clock.Now()
	day := simTime.Day()
	
	current := time.Now()
	require.Equal(t, current.Day() + 1, day)
}