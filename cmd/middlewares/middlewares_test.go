package middlewares

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetDefaultMiddleWare(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDefaultMiddleWare(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDefaultMiddleWare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
