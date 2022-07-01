package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA4NzgxNzMsImlkIjoiNjJiYWJkNTA0OTBjMmE0ODc4MTViYzcxIn0.wcvPD8ly0YMoSiRrUkCQ3upS2xjby4hOU7LLybk7pqQ"

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Token   string                 `json:"token"`
	Data    map[string]interface{} `json:"data"`
}

func TestCreateToken(t *testing.T) {
	Id, _ := primitive.ObjectIDFromHex("62babd50490c2a487815bc71")
	Result, _ := CreateToken(Id, "user")

	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(Result, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	Id_User := claims["id"].(string)

	if Id_User != "62babd50490c2a487815bc71" {
		t.Errorf("Expected value: %v, got: %v", "62babd50490c2a487815bc71", Id_User)
	}

}

func TestAuthMiddleware(t *testing.T) {
	authHandler := func(w http.ResponseWriter, r *http.Request) {}
	req := httptest.NewRequest("GET", "/user", nil)
	req.Header.Set("Authorization", token)
	res := httptest.NewRecorder()
	authHandler(res, req)
	AuthMiddleware(authHandler).ServeHTTP(res, req)
}
