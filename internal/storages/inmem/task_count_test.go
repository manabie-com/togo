package inmem

import (
	"sync"
	"testing"
	"time"

	"github.com/perfectbuii/togo/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Add_Multiple(t *testing.T) {
	var wg sync.WaitGroup
	taskCountStore := NewTaskCountStore()
	task := 1000000
	userId := "test_id"
	date := time.Now().String()
	for i := 1; i <= task; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			taskCountStore.Inc(utils.GetKey(userId, date))
		}()
	}

	wg.Wait()
	assert.Equal(t, task, taskCountStore.Value(utils.GetKey(userId, date)))
}
