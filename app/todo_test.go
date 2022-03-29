package app_test

import (
	"context"
	"errors"

	mock "github.com/golang/mock/gomock"
	"github.com/laghodessa/togo/app"
	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
	"github.com/laghodessa/togo/test/todofixture"
	"github.com/laghodessa/togo/test/todomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TodoUsecase", func() {
	var ctrl *mock.Controller
	var userRepo *todomock.MockUserRepo
	var taskRepo *todomock.MockTaskRepo
	var uc *app.TodoUsecase

	BeforeEach(func() {
		ctrl = mock.NewController(GinkgoT())
		userRepo = todomock.NewMockUserRepo(ctrl)
		taskRepo = todomock.NewMockTaskRepo(ctrl)
	})
	AfterEach(func() { ctrl.Finish() })

	Describe("AddTask", func() {
		var userID string
		var req app.AddTask
		var task todo.Task
		var err error

		JustBeforeEach(func() {
			req = app.AddTask{
				Task: app.Task{
					UserID:  userID,
					Message: "stop coding",
				},
				TimeZone: "Asia/Ho_Chi_Minh",
			}

			uc = &app.TodoUsecase{
				TaskRepo: taskRepo,
				UserRepo: userRepo,
			}
			task, err = uc.AddTask(context.Background(), req)
		})

		BeforeEach(func() {
			userID = domain.NewID()
			userRepo.EXPECT().
				GetUser(mock.Any(), mock.Any()).
				Return(todofixture.NewUser(), nil).AnyTimes()

			taskRepo.EXPECT().
				AddTask(mock.Any(), mock.Any(), mock.Any(), mock.Any()).
				Return(nil).AnyTimes()
			taskRepo.EXPECT().
				CountInTimeRangeByUserID(mock.Any(), mock.Any(), mock.Any(), mock.Any()).
				Return(0, nil).AnyTimes()
		})

		It("succeeds", func() { Expect(err).ShouldNot(HaveOccurred()) })
		It("returns valid task message", func() { Expect(task.Message).To(Equal("stop coding")) })
		It("returns valid task user id", func() { Expect(task.UserID).To(Equal(userID)) })

		Context("request is missing user id", func() {
			BeforeEach(func() { userID = "" })
			It("fails", func() { Expect(err).Should(HaveOccurred()) })
			It("returns invalid argument error", func() { Expect(errors.Is(err, domain.ErrInvalidArg)).To(BeTrue()) })
		})

		Context("user hit daily limit", func() {
			BeforeEach(func() {
				taskRepo = todomock.NewMockTaskRepo(ctrl)
				taskRepo.EXPECT().
					CountInTimeRangeByUserID(mock.Any(), mock.Any(), mock.Any(), mock.Any()).
					Return(10, nil)
			})

			It("fails", func() { Expect(err).Should(HaveOccurred()) })
			It("returns user hit daily limit error", func() { Expect(errors.Is(err, todo.ErrUserHitTaskDailyLimit)).To(BeTrue()) })
			It("returns zero value task", func() { Expect(task).To(BeZero()) })
		})
	})
})
