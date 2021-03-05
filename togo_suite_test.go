package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTogo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Togo Suite")
}
