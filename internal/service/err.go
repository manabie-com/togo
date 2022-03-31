package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInternal          = status.Error(codes.Internal, "internal server error")
	errDailyQuotaReached = status.Error(codes.InvalidArgument, "you have reached daily limit")
	errDailyQuotaExceed  = status.Error(codes.InvalidArgument, "you will exceed daily limit adding this list")
)
