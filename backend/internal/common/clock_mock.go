package common

import "time"

type ClockMock struct {
	timestamps_ms []int64
	count int
}

func MakeClockMock() *ClockMock {
	return &ClockMock {

	}
}

func (c *ClockMock) AddTimestamps(iTimestamp_ms int64) {
	c.timestamps_ms = append(c.timestamps_ms, iTimestamp_ms)
}

func (c *ClockMock) Now() Time {
	if len(c.timestamps_ms) == 0 {
		panic("No time created for mock clock")
	}
	index := c.count
	if index >= len(c.timestamps_ms) {
		index = len(c.timestamps_ms) - 1
	} else {
		c.count += 1
	}
	timestamp_ms := c.timestamps_ms[index]
	return MakeTime(time.Unix(timestamp_ms / 1e3, (timestamp_ms % 1e3) * 1e6))
}
