package middleware

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"github.com/spf13/viper"
	"reflect"
	"testing"
)

func Test_authenticateToken(t *testing.T) {
	viper.Set("jwt.key", "abcdef")
	TrueToken, _ := utils.CreateToken("testID", viper.GetString("abcdef"))
	FalseToken := "falsetoken"
	tests := []struct {
		name         string
		token         string
		wantResponse interface{}
		wantErr      error
	}{
		{
			name: "Success Case",
			token: TrueToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := authenticateToken(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("authenticateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("authenticateToken() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	type args struct {
		endpoint kit.Endpoint
	}
	tests := []struct {
		name string
		args args
		want kit.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Authenticate(tt.args.endpoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}