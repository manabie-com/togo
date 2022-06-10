package common

import "fmt"

const (
	ServiceNullErr = "SERVICE IS NULL"
	PortNullErr    = "PORT IS NULL"

	MySQLUriNullErr = "MySQL's URI IS NULL"
)

var DataIsNullErr = func(obj string) string {
	return fmt.Sprintf("%v CanNotBeNull", obj)
}
