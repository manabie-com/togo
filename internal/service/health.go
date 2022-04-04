package service

import (
	"context"

	"github.com/vchitai/togo/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverImpl) Liveness(ctx context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	return &pb.LivenessResponse{
		Message: "ok",
	}, nil
}

func (s *serverImpl) ToggleReadiness(ctx context.Context, req *pb.ToggleReadinessRequest) (*pb.ToggleReadinessResponse, error) {
	s.isReady = req.GetIsReady()
	return &pb.ToggleReadinessResponse{
		Message: "ok",
	}, nil
}

func (s *serverImpl) Readiness(ctx context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	if !s.isReady {
		return nil, status.Error(codes.Unavailable, "server is not ready")
	}
	return &pb.ReadinessResponse{
		Message: "ok",
	}, nil
}
