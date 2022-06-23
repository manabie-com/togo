package jwt

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Create(w http.ResponseWriter, username string, id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token expired after 12 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ParseToken(tokenString string) jwt.MapClaims{
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString, claims, nil)
	return claims
}

func CheckToken(w http.ResponseWriter, r *http.Request)(string, int, bool){
	token := r.Header.Get("token")
	result := ParseToken(token)
	username := fmt.Sprintf("%v",result["username"])
	id, err := strconv.Atoi(fmt.Sprintf("%v", result["id"]))
	if result["username"] == nil || err != nil{
		return username, id,false
	} 
	return username,id ,true
}