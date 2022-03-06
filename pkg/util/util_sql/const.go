package util_sql

const (
	SimpleSelect = "select %s from %s where %s order by %s;"
	InnerJoin    = " inner join %[1]s on %[1]s.%[3]s = %[2]s.%[4]s "
)
