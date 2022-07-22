package tasklimit

import "testing"

func TestGetUserLimit(t *testing.T) {
	userLimit := GetUserLimiSvc()

	var userID uint64 = 1

	userLimitFirstTime := userLimit.GetUserLimit(userID)
	userLimitSecondTime := userLimit.GetUserLimit(userID)

	if userLimitFirstTime != userLimitSecondTime {
		t.Errorf("Expected user limit first time %d to be equal second time %d", userLimitFirstTime, userLimitSecondTime)
	}
}
