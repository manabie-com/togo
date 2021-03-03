package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Services", func() {
	Describe("Login entry", func() {
		Context("with right account", func() {
			It("should return 401", func() {
				Expect("foo").To(Equal("A"))
			})
		})
	})
})
