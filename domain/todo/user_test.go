package todo_test

import (
	"errors"

	"github.com/laghodessa/togo/domain/todo"
	"github.com/laghodessa/togo/test/todofixture"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewUser", func() {
	var user todo.User
	var err error
	var taskDailyLimit int

	JustBeforeEach(func() {
		user, err = todo.NewUser(
			todo.UserTaskDailyLimit(taskDailyLimit),
		)
	})

	BeforeEach(func() {
		taskDailyLimit = 10
	})

	Context("ok", func() {
		It("succeeds", func() { Expect(err).ShouldNot(HaveOccurred()) })
		It("create new user id", func() { Expect(user.ID).ToNot(BeEmpty()) })
		It("set daily limit task", func() { Expect(user.TaskDailyLimit).To(Equal(taskDailyLimit)) })
	})
})

var _ = Describe("User", func() {
	var user todo.User

	BeforeEach(func() {
		user = todofixture.NewUser()
	})

	Describe("HitTaskDailyLimit", func() {
		var err error
		var todayTotal int

		JustBeforeEach(func() {
			err = user.HitTaskDailyLimit(todayTotal)
		})

		Context("user has a limit of 10 tasks per day", func() {
			Context("user has not reached the daily limit", func() {
				BeforeEach(func() {
					todayTotal = 9
				})

				It("returns nil", func() { Expect(err).ShouldNot(HaveOccurred()) })
			})

			Context("user reached daily limit", func() {
				BeforeEach(func() {
					todayTotal = 10
				})

				It("fails", func() { Expect(err).Should(HaveOccurred()) })
				It("returns hit daily task limit error", func() { Expect(errors.Is(err, todo.ErrUserHitTaskDailyLimit)).To(BeTrue()) })
			})
		})
	})
})
