package todo_test

import (
	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewTask", func() {
	var userID string
	var msg string
	var task todo.Task
	var err error

	JustBeforeEach(func() {
		task, err = todo.NewTask(
			todo.TaskUserID(userID),
			todo.TaskMessage(msg),
		)
	})

	BeforeEach(func() {
		userID = domain.NewID()
		msg = "task message"
	})

	Context("ok", func() {
		It("succeeds", func() { Expect(err).ShouldNot(HaveOccurred()) })
		It("has valid message", func() { Expect(task.Message).To(Equal(msg)) })
		It("has valid user id", func() { Expect(task.UserID).To(Equal(userID)) })
	})

	Context("user id is empty", func() {
		BeforeEach(func() { userID = "" })
		It("fails", func() { Expect(err).Should(HaveOccurred()) })
		It("returns zero task", func() { Expect(task).To(BeZero()) })
	})

	Context("message is empty", func() {
		BeforeEach(func() { msg = "" })
		It("fails", func() { Expect(err).Should(HaveOccurred()) })
		It("returns zero task", func() { Expect(task).To(BeZero()) })
	})
})
