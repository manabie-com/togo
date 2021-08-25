package util

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

func ConvertSQLNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
