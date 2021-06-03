package model

import (
	"reflect"
	"testing"
)

func TestUser_GetUserById(t *testing.T) {
	type fields struct {
		ID       string
		Password string
		MaxTodo  int
	}
	type args struct {
		user_id string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser User
		wantErr  bool
	}{
		{
			name: "Test with true user_id",
			args: args{
				user_id: "firstUser",
			},
			fields: fields{
				ID:       "firstUser",
				Password: "example",
				MaxTodo:  1,
			},
			wantUser: User{
				ID:       "firstUser",
				Password: "example",
				MaxTodo:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				Password: tt.fields.Password,
				MaxTodo:  tt.fields.MaxTodo,
			}
			gotUser, err := u.GetUserById(tt.args.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("User.GetUserById() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

// func TestLogin(t *testing.T) {
// 	type args struct {
// 		username string
// 		password string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    interface{}
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := Login(tt.args.username, tt.args.password)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Login() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
