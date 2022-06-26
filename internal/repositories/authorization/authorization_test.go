package authorization

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func Test_repository_ValidateUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				DB: tt.fields.DB,
			}
			got, err := r.ValidateUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.ValidateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.ValidateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
