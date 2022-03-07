package util

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"

	"github.com/khoale193/togo/pkg/e"
)

func VerifyPassword(password string, hash string, cryptType string) bool {
	switch cryptType {
	case e.AuthenticationTypeMD5:
		m := md5.New()
		m.Write([]byte(password))
		return hex.EncodeToString(m.Sum(nil)) == hash
	case e.AuthenticationTypeBcrypt:
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		return err == nil
	}
	return false
}
