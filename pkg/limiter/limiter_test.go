package limiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLimiter_AllowN(t *testing.T) {
	limiter := New()

	now := time.Now()

	t.Run("should handle id limit properly", func(t *testing.T) {
		limiter.WithLimitByID("id", 3*time.Second, 3)
		require.True(t, limiter.AllowN(now, "id", 3))
		require.False(t, limiter.AllowN(now.Add(2*time.Second), "id", 3))
	})

	t.Run("should handle another id with default limit properly", func(t *testing.T) {
		require.True(t, limiter.AllowN(now, "another id", 5))
		require.False(t, limiter.AllowN(now, "another new id", 6))
	})
}
