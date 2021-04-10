package util

import (
	"crypto/md5"

	"github.com/google/uuid"
)

func Code2UUID(code string) uuid.UUID {
	return uuid.NewHash(md5.New(), uuid.Nil, []byte(code), 4)
}
