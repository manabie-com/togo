package postgres

import (
	"reflect"
	"testing"

	"github.com/manabie-com/togo/pkg/model"

	"github.com/manabie-com/togo/test"

	"github.com/jinzhu/gorm"
)

func Test_repository_GetUser(t *testing.T) {

	type args struct {
		userName string
	}
	tests := []struct {
		name      string
		db        *gorm.DB
		args      args
		mock      func(db *gorm.DB)
		want      *model.User
		wantErr   bool
		cleanMock func(db *gorm.DB)
	}{
		// TODO: Add test cases.
		{
			name: "nil user",
			db:   test.GetTestDb(),
			args: args{
				userName: "testUserName",
			},
			mock: func(db *gorm.DB) {
				userModel := &model.User{}
				db.DropTableIfExists(userModel)
				db.CreateTable(userModel)
			},
			want:    nil,
			wantErr: true,
			cleanMock: func(db *gorm.DB) {
				userModel := &model.User{}
				db.DropTableIfExists(userModel)
			},
		},
		{
			name: "nil user",
			db:   test.GetTestDb(),
			args: args{
				userName: "testUser",
			},
			mock: func(db *gorm.DB) {
				userModel := &model.User{
					UserName: "testUser",
					Password: "1234",
					MaxTodo:  5,
				}
				db.DropTableIfExists(userModel)
				db.CreateTable(userModel)
				db.Save(userModel)
			},
			want: &model.User{
				ID:       1,
				UserName: "testUser",
				Password: "1234",
				MaxTodo:  5,
			},
			wantErr: false,
			cleanMock: func(db *gorm.DB) {
				userModel := &model.User{}
				db.DropTableIfExists(userModel)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.db,
			}

			tt.mock(tt.db)

			got, err := r.GetUser(tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}

			tt.mock(tt.db)
		})
	}
}

func Test_repository_SaveUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		mock      func(db *gorm.DB)
		cleanMock func(db *gorm.DB)
	}{
		// TODO: Add test cases.
		{
			name: "insert user",
			fields: fields{
				db: test.GetTestDb(),
			},
			args: args{
				user: &model.User{
					UserName: "testUserName",
					Password: "abc",
					Salt:     "1234",
					MaxTodo:  10,
				},
			},
			wantErr: false,
			mock: func(db *gorm.DB) {
				userModel := &model.User{}
				db.DropTableIfExists(userModel)
				db.CreateTable(userModel)
			},
			cleanMock: func(db *gorm.DB) {
				userModel := &model.User{}
				db.DropTableIfExists(userModel)
			},
		},
		{
			name: "update user",
			fields: fields{
				db: test.GetTestDb(),
			},
			args: args{
				user: &model.User{
					ID:       1,
					UserName: "testUserName",
					Password: "abc",
					Salt:     "1234",
					MaxTodo:  10,
				},
			},
			wantErr: false,
			mock: func(db *gorm.DB) {
				userModel := &model.User{
					ID:       1,
					UserName: "testUserName",
					Password: "abc",
					Salt:     "",
					MaxTodo:  10,
				}
				db.DropTableIfExists(userModel)
				db.CreateTable(userModel)
			},
			cleanMock: func(db *gorm.DB) {
				userModel := &model.User{}
				db.DropTableIfExists(userModel)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock(tt.fields.db)
			if err := r.SaveUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.mock(tt.fields.db)
		})
	}
}
