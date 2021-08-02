package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/daos"
	"github.com/manabie-com/togo/internal/models"
)

func UserAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			//Check if token is valid.
			isValidToken, token := CheckIfValidToken(r)
			if !isValidToken {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			claims, _ := token.Claims.(jwt.MapClaims)
			//Check if suspended
			_, err := GetAccountFullInfo(claims["account_id"])
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusForbidden)
				return
			}
			ctx := context.WithValue(r.Context(), "UserAccessToken", token)
			ctx = context.WithValue(ctx, "account_id", claims["account_id"])
			ctx = context.WithValue(ctx, "max_daily_tasks_count", claims["max_daily_tasks_count"])
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}
	}
}

func CheckIfValidToken(r *http.Request) (bool, *jwt.Token) {
	if r.Header.Get("Authorization") != "" {
		authorizationToken := r.Header.Get("Authorization")
		customToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
		token, _ := jwt.Parse(customToken, nil)
		if token == nil {
			return false, nil
		} else {
			return true, token
		}
	} else {
		return false, nil
	}
}

func GetAccountFullInfo(id interface{}) (*models.Account, error) {
	accountDAO := daos.GetAccountDAO()
	accountID, _ := uuid.Parse(fmt.Sprint(id))
	accountInfo, err := accountDAO.FindAccountByID(accountID)
	if err != nil {
		return nil, err
	}
	//If account doesn't exist.
	if accountInfo == nil {
		return nil, errors.New("account doesn't exist")
	}
	return accountInfo, nil
}
