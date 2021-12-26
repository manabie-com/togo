package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal = status.Error(codes.Internal, "Internal Server Error")
)

func toPublicError(err error) error {
	switch status.Code(err) {
	case codes.InvalidArgument, codes.FailedPrecondition:
		return err
	default:
		return ErrInternal
	}
}
