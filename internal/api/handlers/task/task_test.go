package task

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/api/handlers"
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
