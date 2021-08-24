package services

import (
	"net/http"
	"testing"

	"github.com/manabie-com/togo/internal/usecase"
)

func TestToDoService_ServeHTTP(t *testing.T) {
	type fields struct {
		UserUsecase usecase.UserUsecase
		TaskUsecase usecase.TaskUsecase
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToDoService{
				UserUsecase: tt.fields.UserUsecase,
				TaskUsecase: tt.fields.TaskUsecase,
			}
			s.ServeHTTP(tt.args.resp, tt.args.req)
		})
	}
}
