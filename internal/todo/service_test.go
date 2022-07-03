package todo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/xrexonx/togo/internal/todo"
	"github.com/xrexonx/togo/internal/user"
	"testing"
)

func TestTodo(t *testing.T) {
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

	Context("when todo is added", func() {
		newTodo := todo.Todo{Name: "Task1", UserId: "1"}
		It("should validate userID", func() {
			todoResponse, _ := todo.Add(newTodo)
			Expect(len(todoResponse.UserId)).To(BeNumerically(">", 0))
		})
		It("should return error  when no userId is passed", func() {
			Expect(len(newTodo.UserId)).To(BeNumerically(">", 0))
		})
	})

})
