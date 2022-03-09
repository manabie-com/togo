package utils_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/triet-truong/todo/utils"
)

type test struct {
	input time.Time
}

var NUM_TEST_CASES int = 20

func randomTimestamp() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow
}

func initTests() []test {
	var tests = []test{}
	for i := 1; i <= NUM_TEST_CASES; i++ {
		tests = append(tests, test{input: randomTimestamp()})
	}
	return tests
}
func TestEndOfCurrentDate(t *testing.T) {
	tests := initTests()
	format := "2006-01-02T15:04:05"
	for _, v := range tests {
		got := utils.EndOfCurrentDate(v.input)
		fmt.Printf("Input: %v, got: %v\n", v.input.Format(format), got.Format(format))

		//Same day with input's day
		assert.Equal(t, v.input.Format("2006-01-02"), got.Format("2006-01-02"))

		//hh:mm:ss is 23:59:59
		assert.Equal(t, "23:59:59", got.Format("15:04:05"))
	}
}
