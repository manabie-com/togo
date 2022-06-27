package redis

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	address = "localhost:6379"
)

// redis does not lock
// https://stackoverflow.com/questions/10650232/locking-and-redis/10650505#:~:text=Redis%20does%20not%20lock.,same%20key%20without%20any%20problems.

func Test_Add_Multiple(t *testing.T) {
	ctx := context.Background()
	client, err := NewRedisClient(address)
	assert.NoError(t, err)
	key := "test2"
	amount := 1000
	var wg sync.WaitGroup
	m := &sync.Map{}

	for i := 1; i <= amount; i++ {
		wg.Add(1)
		go func() {
			if cmd := client.SetNX(ctx, key, 0, time.Second); cmd.Err() != nil {
				fmt.Println(cmd.Err())
			}
			cmd := client.Incr(ctx, key)
			assert.NoError(t, cmd.Err())
			m.Store(cmd.Val(), true)
			wg.Done()
		}()
	}
	wg.Wait()
	cmd := client.Get(ctx, key)
	num, err := cmd.Int()
	assert.NoError(t, err)
	assert.Equal(t, amount, num)
	for i := 1; i <= amount; i++ {
		_, ok := m.Load(int64(i))
		assert.True(t, ok)
	}
}
