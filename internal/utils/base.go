package utils

import (
	"database/sql"
	"net/http"
)

func Value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}
