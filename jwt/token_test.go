package jwt

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

const (
	Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYxMDUzNDEsImlkIjozLCJ1c2VybmFtZSI6Imh1dWxvYyJ9.MqHypeN42fopG5jHWEjk6bu9m7wSENqLewBGq9VC3sA"
)

func TestTokenParse(t *testing.T) {
	info := ParseToken(Token)
	username := fmt.Sprintf("%v", info["username"])
	id := fmt.Sprintf("%v", info["id"])

	if username != "huuloc" || id != "3" {
		log.Fatal("Token parse failed")
	}
}

func TestCreateToken(t *testing.T) {
	var w http.ResponseWriter
	newToken, err := Create(w, "huuloc", 3)
	if err != nil {
		log.Fatal("Create token failed")
	}
	if len(newToken) < 32 {
		log.Fatal("Invalid token created")
	}
}
