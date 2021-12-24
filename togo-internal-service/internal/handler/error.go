package handler

import (
	"togo-internal-service/internal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknow = status.Error(codes.Unknown, codes.Unknown.String())
	errMap    = map[error]error{
		storage.ErrInternal:              status.Error(codes.Internal, storage.ErrInternal.Error()),
		storage.ErrExceedLimitTaskPerDay: status.Error(codes.FailedPrecondition, storage.ErrExceedLimitTaskPerDay.Error()),
		storage.ErrTaskNotFound:          status.Error(codes.NotFound, storage.ErrTaskNotFound.Error()),
	}
)

func toGRPCError(err error) error {
	if err == nil {
		return nil
	}

	if err, ok := errMap[err]; ok {
		return err
	}

	return ErrUnknow
}
