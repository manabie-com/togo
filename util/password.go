package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"manabie-com/togo/global"
)

func computeEmailAndPassword(id, password string) []byte {
	type Combine struct {
		ID       string
		Password string
	}
	var c = Combine{ID: id, Password: password}
	jsonData, _ := json.Marshal(c)

	h := hmac.New(sha256.New, []byte(global.Config.HMACCombinePasswordKey))
	h.Write(jsonData)
	return h.Sum(nil)
}

func HashPassword(id, password string) string {
	var preparation = computeEmailAndPassword(id, password)
	bytes, _ := bcrypt.GenerateFromPassword(preparation, bcrypt.DefaultCost)
	return string(bytes)
}

func CompareHashPasswordAndPassword(hashedPassword, email, password string) bool {
	var preparation = computeEmailAndPassword(email, password)
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), preparation) != nil {
		return true
	}
	return false
}
