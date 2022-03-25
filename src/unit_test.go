package main

import (
	"fmt"
	"testing"
)

func TestHashFunction(t *testing.T) {
	type fields struct {
		key string
	}
	tests := []struct {
		fields fields
		hash   uint
	}{
		{
			fields: fields{
				key: "12320220114",
			},
			hash: 1,
		},
		{
			fields: fields{
				key: "a7asu2gasd982vaa20220114",
			},
			hash: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if hash := GetMaxTasks(tt.fields.key); hash != tt.hash {
				t.Errorf("GetMaxTasks() = %v, we want hash %v", hash, tt.hash)
			}
		})
	}
}
