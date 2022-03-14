package usecase

import (
	"testing"
)

//func TestValidateToken(t *testing.T) {
//	type args struct {
//		tokenString string
//	}
//	tests := []struct {
//		name          string
//		args          args
//		wantErrString string
//	}{
//		// TODO: Add test cases.
//		{name: "check valid token", args: struct{ tokenString string }{tokenString:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzQyNTYyNX0.EBdw6nYJhQcJPG42f7gadA85nXcm_rNsSOH8s_hq2oQ"}, wantErrString:"" },
//		{name: "check invalid token", args: struct{ tokenString string }{tokenString:  "asasaeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzQyNTYyNX0.EBdw6nYJhQcJPG42f7gadA85nXcm_rNsSOH8s_hq2oQ"}, wantErrString:"That's not even a token"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotErrString := ValidateToken(tt.args.tokenString); gotErrString != tt.wantErrString {
//				t.Errorf("ValidateToken() = %v, want %v", gotErrString, tt.wantErrString)
//			}
//		})
//	}
//}

func TestValidateToken1(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name          string
		args          args
		wantUserId    string
		wantErrString string
	}{
		// TODO: Add test cases.
		{name: "check valid token", args: struct{ tokenString string }{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzQyNTYyNX0.EBdw6nYJhQcJPG42f7gadA85nXcm_rNsSOH8s_hq2oQ"}, wantErrString: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserId, gotErrString := ValidateToken(tt.args.tokenString)
			if gotUserId != tt.wantUserId {
				t.Errorf("ValidateToken() gotUserId = %v, want %v", gotUserId, tt.wantUserId)
			}
			if gotErrString != tt.wantErrString {
				t.Errorf("ValidateToken() gotErrString = %v, want %v", gotErrString, tt.wantErrString)
			}
		})
	}
}
