// +build integration

package rest_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/laghodessa/togo/domain/todo"
	"github.com/laghodessa/togo/test/todofixture"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/tasks", func() {
	var method string
	var target string
	var body string
	var req *http.Request
	var user = todofixture.NewUser(func(u *todo.User) { u.TaskDailyLimit = 1 })

	JustBeforeEach(func() {
		req = httptest.NewRequest(method, "http://togo.com/api/v1"+target, bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
	})

	Describe("POST /tasks", func() {
		BeforeEach(func() {
			Expect(userRepo.AddUser(context.Background(), user)).To(Succeed())
			method = "POST"
			target = "/tasks"
			body = fmt.Sprintf(`
			{
				"timeZone": "Asia/Ho_Chi_Minh",
				"task": {
					"userId": "%s",
					"message": "todo message"
				}
			}`, user.ID)
		})

		It("returns 201", func() {
			resp, err := server.Test(req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(HaveHTTPStatus(http.StatusCreated))
			Expect(resp).To(HaveHTTPBody(ContainSubstring("todo message")))
		})

		Context("user does not exist", func() {
			BeforeEach(func() {
				body = `{
					"task": { "userId": "doesnotexist" }
				}`
			})

			It("returns 404", func() {
				resp, err := server.Test(req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).To(HaveHTTPStatus(http.StatusNotFound))
				Expect(resp).To(HaveHTTPBody(MatchJSON(`
				{"code": "not_found", "message": "user not found"}
			`)))
			})
		})

		Context("user's daily limit reached", func() {
			BeforeEach(func() {
				body = fmt.Sprintf(`
				{
					"timeZone": "Asia/Ho_Chi_Minh",
					"task": {
						"userId": "%s",
						"message": "todo message"
					}
				}`, user.ID)
			})
			JustBeforeEach(func() {
				// setup hit user daily limit
				resp, err := server.Test(req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).To(HaveHTTPStatus(http.StatusCreated))
			})

			It("returns 422", func() {
				resp, err := server.Test(req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).To(HaveHTTPStatus(http.StatusUnprocessableEntity))
			})
		})
	})
})
