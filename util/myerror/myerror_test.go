package myerror_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/pkg/errors"

	"github.com/khangjig/togo/util/myerror"
)

func TestMyError_Error(t *testing.T) {
	t.Parallel()

	t.Run("NilRaw", func(t *testing.T) {
		t.Parallel()

		err := myerror.NewError(nil, http.StatusBadRequest, "000", "Failed.")

		if err.Error() != err.Message {
			t.Errorf("current = '%v' want '%v'", err.Error(), err.Message)
		}
	})

	t.Run("NotNilRaw", func(t *testing.T) {
		t.Parallel()

		err := myerror.NewError(context.Canceled, http.StatusBadRequest, "000", "Failed.")

		if err.Error() != errors.Wrap(err.Raw, err.Message).Error() {
			t.Errorf("current = '%v' want '%v'", err.Error(), errors.Wrap(err.Raw, err.Message).Error())
		}
	})
}

func TestMyError_Is(t *testing.T) {
	t.Parallel()

	t.Run("NilRaw", func(t *testing.T) {
		t.Parallel()

		err := myerror.NewError(nil, http.StatusBadRequest, "000", "Failed.")

		if !err.Is(myerror.MyError{}) {
			t.Errorf("current = '%v' want '%v'", err, nil)
		}
	})

	t.Run("NotNilRaw", func(t *testing.T) {
		t.Parallel()

		err := myerror.NewError(context.Canceled, http.StatusBadRequest, "000", "Failed.")

		if !err.Is(context.Canceled) {
			t.Errorf("current = '%v' want '%v'", err, context.Canceled)
		}
	})
}
