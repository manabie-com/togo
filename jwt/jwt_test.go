package jwt

import (
	"fmt"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYxMDUzNDEsImlkIjozLCJ1c2VybmFtZSI6Imh1dWxvYyJ9.MqHypeN42fopG5jHWEjk6bu9m7wSENqLewBGq9VC3sA"
	TokenUsername = "huuloc"
	TokenId = "3"
)

// test function token parse
func TestTokenParse(t *testing.T) {
	info := ParseToken(Token)
	username := fmt.Sprintf("%v", info["username"])
	id := fmt.Sprintf("%v", info["id"])

	if username != TokenUsername || id != TokenId {
		log.Fatal("Token parse failed")
	}
}

// test function create token
func TestCreateToken(t *testing.T) {
	newToken, err := Create(TokenUsername, 3)
	if err != nil {
		log.Fatal("Create token failed")
	}
	if len(newToken) < 32 {
		log.Fatal("Invalid token created")
	}
}


func TestCheckToken(t *testing.T){
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/tasks", nil)
	req.Header.Set("token", Token)
	username, id, check := CheckToken(w, req)
	if username != TokenUsername || strings.Compare(fmt.Sprintf("%v",id), TokenId) != 0 || check==false{
		t.Fatal("Token check failed")
	}
}