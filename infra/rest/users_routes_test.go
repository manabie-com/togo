// +build integration

package rest_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/users", func() {
	var method string
	var target string
	var body string
	var req *http.Request

	JustBeforeEach(func() {
		req = httptest.NewRequest(method, "http://togo.com/api/v1"+target, bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
	})

	Describe("POST /users", func() {
		BeforeEach(func() {
			method = "POST"
			target = "/users"
			body = `{"taskDailyLimit": 1}`
		})

		It("returns 201", func() {
			resp, err := server.Test(req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(HaveHTTPStatus(http.StatusCreated))
			Expect(resp).To(HaveHTTPBody(ContainSubstring(`"taskDailyLimit":1`)))
		})
	})
})
