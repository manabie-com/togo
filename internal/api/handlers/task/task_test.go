package task

import (
	"reflect"
	"testing"

	"example.com/m/v2/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestAddTask(t *testing.T) {
	type args struct {
		service handlers.MainUseCase
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTask(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
