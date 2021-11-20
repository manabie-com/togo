package server

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

const (
	TSHARK_PREFER_SETTING = "gui.column.format:\"No.\",\"%m\",\"Time\",\"%t\",\"Source\",\"%s\",\"Destination\"," +
		"\"%d\", \"Protocol\",\"%p\",\"Length\",\"%L\",\"Info\",\"%i\""
	LocalTimeFormat = "2006-01-02 15:04:05"
)

func ConvertTimestamp2LocalTime(t *timestamp.Timestamp) string {
	localTime, err := ptypes.Timestamp(t)
	if err != nil {
		panic(err)
	}
	localTime = localTime.In(time.Local)
	return localTime.Format(LocalTimeFormat)
}

func ConvertTimestamp2RFC3339(t *timestamp.Timestamp) string {
	t_, err := ptypes.Timestamp(t)
	if err != nil {
		panic(err)
	}
	return t_.Format(time.RFC3339)
}

// CurrentTimeBefore ()
func CurrentTimeBefore(m uint64) time.Time {
	currentTime := time.Now()
	result := currentTime.Add(time.Duration(-m) * time.Second)
	return result
}
