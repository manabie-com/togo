// +build integration

package rest_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/tasks", func() {
	var method string
	var target string
	var body string
	var req *http.Request

	JustBeforeEach(func() {
		req = httptest.NewRequest(method, "http://togo.com/api/v1"+target, bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
	})

	Describe("POST /tasks", func() {
		BeforeEach(func() {
			method = "POST"
			target = "/tasks"
		})

		It("returns 201", func() {
			// resp, err := server.Test(req)
			// Expect(err).To(Succeed())
			// Expect(resp).To(HaveHTTPStatus(http.StatusCreated))
		})

		Context("user does not exist", func() {
			BeforeEach(func() {
				body = `{"userId": "doesnotexist"}`
			})

			It("returns 404", func() {
				resp, err := server.Test(req)
				Expect(err).To(Succeed())
				Expect(resp).To(HaveHTTPStatus(http.StatusNotFound))
				Expect(resp).To(HaveHTTPBody(MatchJSON(`
				{"code": "not_found", "message": "user not found"}
			`)))
			})
		})

		Context("user's daily limit reahed", func() {
			It("returns 422", func() {
				// resp, err := server.Test(req)
				// Expect(err).To(Succeed())
				// Expect(resp).To(HaveHTTPStatus(http.StatusUnprocessableEntity))
			})
		})
	})
})
