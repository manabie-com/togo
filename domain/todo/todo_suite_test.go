package todo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTodo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Todo Suite")
}
