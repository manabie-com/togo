package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/api/handlers"
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
