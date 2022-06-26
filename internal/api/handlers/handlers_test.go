package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetValueCookieFromCtx(t *testing.T) {
	type args struct {
		ctx       *gin.Context
		keyCookie string
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetValueCookieFromCtx(tt.args.ctx, tt.args.keyCookie)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValueCookieFromCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetValueCookieFromCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
