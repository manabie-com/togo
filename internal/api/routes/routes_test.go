package routes

import (
	"testing"

	"example.com/m/v2/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestSetupRoute(t *testing.T) {
	type args struct {
		router  *gin.Engine
		service handlers.MainUseCase
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupRoute(tt.args.router, tt.args.service)
		})
	}
}
