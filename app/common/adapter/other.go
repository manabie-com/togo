package adapter

import "sync"

const (
	NotFoundConfigOther = "not found config other"
)

// Others info
type Others map[string]string

var (
	onceOtherMutex = sync.RWMutex{}
)

// Get other by config name
func (other Others) Get(name string) (result string) {
	onceOtherMutex.Lock()
	defer onceOtherMutex.Unlock()

	if value, ok := other[name]; ok {
		result = value
	} else {
		panic(NotFoundConfigOther + name)
	}

	return
}

// Set other by config name
func (other Others) Set(name string, value string) (result bool) {
	onceOtherMutex.Lock()
	defer onceOtherMutex.Unlock()

	other[name] = value
	return true
}
