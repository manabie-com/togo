package common

import (
	"os"
    "strconv"
	"time"
)

type ClockSim struct {
	offset_day int
}

func MakeClockSim() ClockSim {
	offsetStr := os.Getenv("SIM_OFFSET_DAY")
	offset := 0
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			panic("Invalid SIM_OFFSET_DAY. Must be int")
		}
	}

	clockSim := ClockSim {
		offset_day: offset,
	}
	return clockSim
}

func (c ClockSim) Now() Time {
	current := time.Now().Local()
	return MakeTime(current.Add(time.Duration(c.offset_day * 24) * time.Hour))
}