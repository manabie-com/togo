package utils

import (
	"reflect"
	"testing"
)

func TestSafeString(t *testing.T) {
	type args struct {
		s *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeString(tt.args.s); got != tt.want {
				t.Errorf("SafeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicate(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Success case 1",
			args: args{
				[]string{"1", "2", "3", "1"},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicate(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
