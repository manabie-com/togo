package common

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"

	// "grpc_server/audit_middleware"
	"mini_project/config"
	"mini_project/db/model"

	// "grpc_server/logutil"
	// "grpc_server/rbac"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gophercloud/gophercloud"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ParseToken(token string, publicKey *ecdsa.PublicKey) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			fmt.Println("Unexpected signing method: ", zap.Reflect("alg", t.Header["alg"]))
			return nil, fmt.Errorf("invalid token")
		}
		return publicKey, nil
	})
	if err == nil && jwtToken.Valid {
		return jwtToken, nil
	}
	return jwtToken, err
}

func ApplyAuth(ctx context.Context) (context.Context, error) {
	return nil, status.Errorf(codes.Unimplemented, "Please implement auth func for this service")
}

func VerifyTokenBearerOrBasic(ctx context.Context, fullMethodName string,
	db model.DatabaseAPI, authClient *gophercloud.ServiceClient) (context.Context, error) {
	newCtx, err, caller := verifyTokenBasicOrBearer(ctx, fullMethodName, db, authClient)
	fmt.Println("VerifyTokenBearerOrBasic >> caller >>", caller)
	return newCtx, err
}

func verifyTokenBasicOrBearer(ctx context.Context, fullMethodName string,
	db model.DatabaseAPI, authClient *gophercloud.ServiceClient) (context.Context, error, string) {
	var caller = "unknown user"
	val := metautils.ExtractIncoming(ctx).Get("authorization")
	if val == "" {
		return nil, status.Errorf(codes.Unauthenticated, "Request unauthenticated with bear or basic"), caller
	}
	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return nil, status.Errorf(codes.Unauthenticated, "Bad authorization string"), caller
	}

	if strings.EqualFold(splits[0], "basic") {
		payload, _ := base64.StdEncoding.DecodeString(splits[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) < 2 {
			return nil, status.Errorf(codes.Unauthenticated, "Bad authorization string"), caller
		}
		caller = pair[0]
		// if pair[0] != "admin" {
		// user, err := db.GetUserByName(pair[0])
		// if err != nil {
		// 	return nil, status.Errorf(codes.Unauthenticated, "try to login with unknown user or deleted user"), caller
		// }
		// caller := user.Name
		// err = checkUserPermission(ctx, fullMethodName, user, ac)
		// if err != nil {
		// 	return nil, err, caller
		// }
		// }
		newCtx := context.WithValue(ctx, "caller", caller)
		return context.WithValue(newCtx, "token", pair), nil, caller
	} else if strings.EqualFold(splits[0], "bearer") {
		newCtx, err, caller := verifyTokenBearer(ctx, fullMethodName, db, authClient)
		fmt.Println("verifyTokenBasicOrBearer bearer >>>", caller)
		return newCtx, err, caller
	}
	return nil, status.Errorf(codes.Unauthenticated, "Request unauthenticated with bear or basic"), caller
}

func VerifyTokenBearer(ctx context.Context, fullMethodName string,
	db model.DatabaseAPI) (context.Context, error) {
	newCtx, err, caller := verifyTokenBearer(ctx, fullMethodName, db, nil)
	// if err != nil {
	// 	// export security metrics
	// 	// audit_middleware.ExportSecurityMetrics(ctx, fullMethodName, err, caller)
	// }
	fmt.Println("VerifyTokenBearer caller", caller)
	return newCtx, err
}

func verifyTokenBearer(ctx context.Context, fullMethodName string,
	db model.DatabaseAPI, authClient *gophercloud.ServiceClient) (context.Context, error, string) {
	caller := "unknown user"
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error()), caller
	}
	jwtToken, err := ParseToken(token, config.JwtPublicKey)
	if err != nil {
		if jwtToken != nil {
			caller = getCallerFromTokenExpired(jwtToken, db)
		}
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err), caller
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err), caller
	}
	userID := claims["sub"]
	user, err := db.GetUser(fmt.Sprintf("%v", userID))
	if err != nil {
		caller = "deleted user or unknown user"
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err), caller
	}
	caller = user.Name
	audit := extractAudit(claims)
	// if user.AuditIDs != strings.Join(audit, ";") {
	// 	return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", "Token was revoke"), caller
	// }

	// if user.Name != "admin" {
	// 	err = checkUserPermission(ctx, fullMethodName, user, ac)
	// 	if err != nil {
	// 		return nil, err, caller
	// 	}
	// }

	if authClient != nil {
		authClient.SetToken(token)
	}
	newCtx := context.WithValue(ctx, "caller", caller)
	newCtx = context.WithValue(newCtx, "audit_ids", strings.Join(audit, ";"))

	return context.WithValue(newCtx, "user_id", userID), nil, caller
}

func getCallerFromTokenExpired(token *jwt.Token, db model.DatabaseAPI) string {
	caller := "unknow user"
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return caller
	}
	userId := claims["sub"]

	user, err := db.GetUser(fmt.Sprintf("%v", userId))
	if err != nil {
		caller = "deleted user or unknown user"
		return caller
	}
	caller = user.Name
	return caller
}

func extractAudit(claims jwt.MapClaims) []string {
	audit := claims["openstack_audit_ids"].([]interface{})
	data := make([]string, len(audit))
	for i, v := range audit {
		data[i] = fmt.Sprint(v)
	}
	return data
}
