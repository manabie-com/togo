package utils

import "time"

const MySQLTimeFormat = "2006-01-02 15:04:05"

func GetCurrentTime() string {
	return time.Now().Format(MySQLTimeFormat)
}
