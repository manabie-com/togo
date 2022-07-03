package routes_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/xrexonx/togo/cmd/app/config/routes"
	responseUtils "github.com/xrexonx/togo/internal/utils/response"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingDemo Suite")
}

var _ = Describe("Togo functional test", func() {
	var (
		server *ghttp.Server
	)
	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})
	AfterEach(func() {
		server.Close()
	})

	Context("When get request is sent to health check", func() {
		BeforeEach(func() {
			server.AppendHandlers(routes.HealthCheckHandler)
		})
		It("should return API is running message", func() {
			resp, err := http.Get(server.URL() + "/healthCheck")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			var response responseUtils.Response
			resp.Body.Close()
			json.Unmarshal(body, &response)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(response.Message).To(Equal("API is running"))
		})
	})

	//Context("When get request is sent to health check", func() {
	//	BeforeEach(func() {
	//		server.AppendHandlers(routes.AddTodoHandler)
	//	})
	//	It("should return API is running message", func() {
	//		var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
	//		resp, err := http.Post(server.URL() + "/api/v1/todo",  bytes.NewBuffer(jsonStr))
	//
	//		//var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
	//		//req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
	//		//req.Header.Set("Content-Type", "application/json")
	//
	//
	//		Expect(err).ShouldNot(HaveOccurred())
	//		Expect(resp.StatusCode).Should(Equal(http.StatusOK))
	//		body, err := ioutil.ReadAll(resp.Body)
	//		var response responseUtils.Response
	//		resp.Body.Close()
	//		json.Unmarshal(body, &response)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		Expect(response.Message).To(Equal("API is running"))
	//	})
	//})

})
