package graceful

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

var _ http.Handler = new(handler)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`succeeded`))
}

func TestGraceful(t *testing.T) {
	t.Run("should process request as normal before receiving signal", func(t *testing.T) {
		s := httptest.NewUnstartedServer(new(handler))

		go func() {
			require.NoError(t, Graceful(func() error {
				s.Start()
				return nil
			}, func(ctx context.Context) error {
				s.Close()
				return nil
			}))
		}()

		res, err := http.Get(fmt.Sprintf("http://%s", s.Listener.Addr().String()))

		require.NoError(t, syscall.Kill(syscall.Getpid(), syscall.SIGINT))
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}
