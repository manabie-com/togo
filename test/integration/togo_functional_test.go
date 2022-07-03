package integration_test

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Status  string
	Message string
	Code    int
	Data    interface{}
	Date    time.Time
}

var _ = Describe("Togo functional test", func() {
	Context("When get request is sent to health check", func() {
		It("should return API is running message", func() {
			req, _ := http.NewRequest("GET", testENV+"/healthCheck", nil)
			resp, err := client.Do(req)
			body, _ := ioutil.ReadAll(resp.Body)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(resp.Status).To(Equal("200 OK"))
			var response Response
			json.Unmarshal(body, &response)
			Expect(response.Message).To(Equal("API is running"))
		})
	})

	Context("Given an invalid endpoint", func() {
		It("should return 404 not found", func() {
			req, _ := http.NewRequest("GET", testENV+"/invalid", nil)
			resp, err := client.Do(req)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(resp.Status).To(Equal("404 Not Found"))
		})
	})

	Context("When user add a new todo items", func() {
		It("should validate a missing userID", func() {

			var todoStr = []byte(`{"name":"test"}`)
			req, _ := http.NewRequest("POST", testENV+"/todo", bytes.NewBuffer(todoStr))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			var response Response
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &response)
			Expect(response.Message).To(Equal("userID is required"))
		})

		It("should save todo items", func() {

			var todoStr = []byte(`{"name":"test todo", "userId":"1"}`)
			req, _ := http.NewRequest("POST", testENV+"/todo", bytes.NewBuffer(todoStr))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			var response Response
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &response)
			Expect(response.Message).To(Equal("Successfully created todo"))
		})

		It("should validate user max daily limit", func() {

			var todoStr = []byte(`{"name":"test todo", "userId":"3"}`)
			req, _ := http.NewRequest("POST", testENV+"/todo", bytes.NewBuffer(todoStr))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			var response Response
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &response)
			Expect(response.Message).To(Equal("user have reached the maximum daily limit"))
		})

	})

})
