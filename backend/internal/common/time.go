package common

import (
	"time"
	"database/sql/driver"
	"fmt"
)

type Time struct {
	time.Time
}

func MakeTime(
	iTime time.Time,
) Time {
	return Time {
		Time: iTime,
	}
}


func (t Time) Value() (driver.Value, error) {
	return t.Format(time.RFC3339), nil
}

func (t *Time) Scan(value interface{}) error {
	/// will be in local time zone
    if base, ok := value.(time.Time); !ok {
		return fmt.Errorf("is not a time")
    } else {
		base = base.In(time.Local)
		*t = MakeTime(base)
		return nil
	}
}


func (t Time) IsEqual(iTime Time) bool {
	return t.Format(time.RFC3339) == iTime.Format(time.RFC3339)
}