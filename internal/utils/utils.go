package utils

import "database/sql"

func SqlString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}
