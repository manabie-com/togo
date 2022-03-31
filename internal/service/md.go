package service

import (
	"fmt"

	"google.golang.org/grpc/metadata"
)

func getUserID(md metadata.MD) (string, error) {
	userID := md.Get("user-id")
	if len(userID) == 0 {
		return "", fmt.Errorf("user not found")
	}
	return userID[0], nil
}
