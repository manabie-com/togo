package grpc_client

import (
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func NewConnection(address string, opts ...*grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts)
	conn, err := grpc.Dial(address,
		grpc.WithChainUnaryInterceptor(
			grpc_prometheus.UnaryClientInterceptor,
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithBackoff(grpc_retry.BackoffExponential(50*time.Millisecond)),
				grpc_retry.WithCodes(codes.Unavailable),
				grpc_retry.WithPerRetryTimeout(10*time.Second),
			),
		),
		grpc.WithChainStreamInterceptor(
			grpc_prometheus.StreamClientInterceptor,
			grpc_retry.StreamClientInterceptor(
				grpc_retry.WithBackoff(grpc_retry.BackoffExponential(50*time.Millisecond)),
				grpc_retry.WithCodes(codes.Unavailable),
				grpc_retry.WithPerRetryTimeout(10*time.Second),
			),
		),
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithTimeout(10*time.Second),
	)
	return conn, err
}
