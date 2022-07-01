package id

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		{
			name: "Success case",
			want: uuid.UUID{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewID()
			require.IsType(t, tt.want, got)
			require.NotEqual(t, "", got)
		})
	}
}

func TestUUIDIsNil(t *testing.T) {
	{ //Success key
		id := &uuid.UUID{}
		require.Equal(t, true, UUIDIsNil(id))
	}
	{ //Fail key
		require.Equal(t, false, UUIDIsNil(nil))
	}
}
