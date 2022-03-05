package util_test

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/khangjig/togo/util"
)

func TestIsEmail(t *testing.T) {
	t.Parallel()

	t.Run("TestIsEmailSuccess", func(t *testing.T) {
		t.Parallel()

		checked, _ := util.IsEmail("admin@test.com")
		if !checked {
			t.Errorf("IsEmail() = %v, want %v", checked, true)
		}
	})

	t.Run("TestIsEmailFail", func(t *testing.T) {
		t.Parallel()

		checked, _ := util.IsEmail("admin@testcom")
		if checked {
			t.Errorf("IsEmail() = %v, want %v", checked, false)
		}
	})
}

func TestHashPassword(t *testing.T) {
	t.Parallel()

	t.Run("EmptyString", func(t *testing.T) {
		t.Parallel()
		_, err := util.HashPassword("")
		if err == nil {
			t.Errorf("current %v, want %v, ", err, errors.New("empty password"))
		}
	})

	t.Run("NotEmptyString", func(t *testing.T) {
		t.Parallel()
		_, err := util.HashPassword("12345678")
		if err != nil {
			t.Errorf("current: %v, want <nil>", err)
		}
	})
}

func TestComparePassword(t *testing.T) {
	t.Parallel()

	t.Run("ComparePasswordSuccess", func(t *testing.T) {
		t.Parallel()
		check := util.ComparePassword("12345678", "$2a$10$5hr.PXZhYHdIPbL4u5nAtOJtBnp6cGxQXIBPOHnkc/lc.ttOCovi6")
		if !check {
			t.Errorf("current %v, want %v", check, true)
		}
	})

	t.Run("ComparePasswordFail", func(t *testing.T) {
		t.Parallel()
		check := util.ComparePassword("1234567", "$2a$10$5hr.PXZhYHdIPbL4u5nAtOJtBnp6cGxQXIBPOHnkc/lc.ttOCovi6")
		if check {
			t.Errorf("current %v, want %v", check, false)
		}
	})
}
