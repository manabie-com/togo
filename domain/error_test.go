package domain_test

import (
	"errors"

	"github.com/laghodessa/togo/domain"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NotFound", func() {
	err := domain.NotFound("not found")

	It("returns error equivalent to ErrNotFound", func() {
		Expect(errors.Is(err, domain.ErrNotFound)).To(BeTrue())
	})
})

var _ = Describe("InvalidArg", func() {
	err := domain.InvalidArg("invalid argument")

	It("returns error equivalent to ErrInvalidArg", func() {
		Expect(errors.Is(err, domain.ErrInvalidArg)).To(BeTrue())
	})
})
