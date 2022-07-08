package router

import (
	"testing"
	"todo/be/env"
)

func TestConstValues(t *testing.T) {
	if http_POST != "POST" {
		t.Errorf("Output expect 'POST' instead of %v", http_POST)
	}
	if len(key_TokenAes) != 32 {
		t.Errorf("Output expect %d instead of %d", 32, len(key_TokenAes))
	}
	if key_TokenHeader != "Token" {
		t.Errorf("Output expect 'Token' instead of %v", key_TokenHeader)
	}
}

func TestInitUserToken(t *testing.T) {
	result := initUserToken()
	if result.Limit < env.MIN_TASK || env.MAX_TASK < result.Limit {
		t.Errorf("User.Limit expect in rage [%d, %d] instead of %d", env.MIN_TASK, env.MAX_TASK, result.Limit)
	}
	if len(result.UserId) != env.LENGTH_USER_ID {
		t.Errorf("User.UserId expect length is %d instead of %d", env.LENGTH_USER_ID, len(result.UserId))
	}
}

func TestGetUserData(t *testing.T) {
	result, isNew := getUserData("K6eVE14VqusQPr-rS6HQUFJLEeqPBEuUcp5Y_5NnoQQK41sjhXUK4ARR62U-etIigjfYHXG6VA3ozf4=")
	if isNew {
		t.Errorf("Output expect false instead of true")
	}
	if result.Limit != 8 {
		t.Errorf("Output expect 8 instead of %d", result.Limit)
	}
	if result.UserId != "GerUBjdKysGBRdpPMjF4" {
		t.Errorf("Output expect 'GerUBjdKysGBRdpPMjF4' instead of %v", result.UserId)
	}

	_, isNew = getUserData("")
	if !isNew {
		t.Errorf("Output expect true instead of false")
	}

	_, isNew = getUserData("sample text")
	if !isNew {
		t.Errorf("Output expect true instead of false")
	}
}
