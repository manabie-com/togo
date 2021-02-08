package services

import (
	"context"
	"testing"
)

func TestUserIdFromContext(t *testing.T) {
	contextData := context.WithValue(context.Background(), userAuthKey(0), "firstUser")
	userId, ok := userIDFromCtx(contextData)
	if !(userId == "firstUser" && ok) {
		t.Errorf("Expect (userId,ok) is (\"firstUser\",true), but the fact is (%s,%t)", userId, ok)
	}
}
