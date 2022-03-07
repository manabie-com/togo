package utils

import "github.com/sirupsen/logrus"

func ConvertToFloat64(myInterface interface{}) float64 {
	switch myInterface.(type) {
	case int64:
		// v is an int here, so e.g. v + 1 is possible.
		return float64(myInterface.(int64))
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		return myInterface.(float64)
	}
	return 0
}

func ErrorLog(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

func FatalLog(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
