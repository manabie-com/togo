package common

import (
	"github.com/lib/pq"
)

/// normalize common sql errors
/// if error is not supported, the same error is returned back
func NormalizeSqlError(iPqError error) error {
	ret := iPqError
	if pqErr, ok := iPqError.(*pq.Error); ok {
		if pqErr.Code == "40001" {
			ret = SqlSerializableTransactionError
		}
	} 
	return ret
}