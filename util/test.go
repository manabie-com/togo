package util

import (
	"time"

	"bou.ke/monkey"
)

var (
	pseudoIndex int
	pseudoUUID  = []string{
		"af1c772f-9abd-4e3c-94af-80d57d262028",
		"31e578f1-500a-492b-b721-1747947877c9"}
)

// MockRuntimeFunc - for handling those runtime function in testing like time.Now, newID, ...
func MockRuntimeFunc() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 8, 20, 20, 34, 58, 651387237, time.UTC)
	})
	monkey.Patch(NewUUID, func() string {
		if pseudoIndex >= len(pseudoUUID) {
			pseudoIndex = 0
		}
		defer func() {
			pseudoIndex++
		}()
		return pseudoUUID[pseudoIndex]
	})
}
