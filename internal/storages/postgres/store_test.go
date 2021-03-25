package postgres

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStore(t *testing.T) {
	store := NewStore(testDB)

	if store == nil {
		log.Fatal("cannot create sql store")
	}

	require.NotEmpty(t, store)
}
