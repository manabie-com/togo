package todo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/xrexonx/togo/internal/user"
	"testing"
)

func TestTodoService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Todo Suite")
}

var _ = Describe("Todo", func() {

	BeforeEach(func() {})
	AfterEach(func() {})

	Context("initially user has no todo items", func() {
		user := user.User{MaxDailyLimit: 0}
		It("has 0 items", func() {
			Expect(user.MaxDailyLimit).Should(BeZero())
		})
	})

	// TODO: Add more testing
	Context("when todo is added", func() {})

})
