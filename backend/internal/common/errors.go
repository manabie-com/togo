package common

import "errors"

var NotFound = errors.New("NotFound")
var SqlSerializableTransactionError = errors.New("SqlSerializableTransactionError")