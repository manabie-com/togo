package integration_test

import (
	"fmt"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"os"
	"testing"
)

type Httpclient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	testENV string
	client  Httpclient
	_       = BeforeSuite(func() {

		// Load env using relative path
		if err := godotenv.Load("../../.env"); err != nil {
			log.Fatal("No .env file found")
		}

		apiURL := fmt.Sprintf(
			"http://%s:%s/api/%s",
			os.Getenv("HOST"),
			os.Getenv("PORT"),
			os.Getenv("API_VERSION"),
		)
		testENV = apiURL
		client = &http.Client{}
	})
	_ = AfterSuite(func() {})
)

func TestTodo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestTodo Suite")
}
