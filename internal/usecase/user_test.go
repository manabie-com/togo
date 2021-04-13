package usecase

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"testing"
)

func Test_uc_Validate(t *testing.T) {
	type fields struct {
		task storages.Task
	}
	type args struct {
		ctx      context.Context
		user     sql.NullString
		password sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &uc{
				task: tt.fields.task,
			}
			if got := u.Validate(tt.args.ctx, tt.args.user, tt.args.password); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uc_CreateToken(t *testing.T) {
	type fields struct {
		task storages.Task
	}
	type args struct {
		id     string
		jwtKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &uc{
				task: tt.fields.task,
			}
			got, err := u.CreateToken(tt.args.id, tt.args.jwtKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uc_ValidToken(t *testing.T) {
	type fields struct {
		task storages.Task
	}
	type args struct {
		token  string
		JWTKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &uc{
				task: tt.fields.task,
			}
			got, got1 := u.ValidToken(tt.args.token, tt.args.JWTKey)
			if got != tt.want {
				t.Errorf("ValidToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ValidToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
