package util

import "database/sql"

func ConvertSQLNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}
