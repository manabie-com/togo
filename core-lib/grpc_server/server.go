package grpc_server

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	//grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	//grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	//grpc_opentelemetry "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	//"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCServer struct {
	name string
	port int
	*grpc.Server
}

func New(name string, port int) *gRPCServer {
	grpcServer := &gRPCServer{
		name: name,
		port: port,
		Server: grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_recovery.StreamServerInterceptor(),
				//grpc_opentracing.StreamServerInterceptor(),
				//grpc_opentelemetry.StreamServerInterceptor(),
				//grpc_zap.StreamServerInterceptor(logger.GetInstance()),
				//grpc_auth.StreamServerInterceptor(myAuthFunction),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				//grpc_opentracing.UnaryServerInterceptor(),
				//grpc_opentelemetry.UnaryServerInterceptor(),
				//grpc_zap.UnaryServerInterceptor(zapLogger),
				//grpc_auth.UnaryServerInterceptor(myAuthFunction),
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
			)),
		),
	}

	grpc_prometheus.Register(grpcServer.Server)
	grpc_prometheus.EnableHandlingTimeHistogram()

	reflection.Register(grpcServer.Server)
	return grpcServer
}

func (g *gRPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return err
	}
	defer g.GracefulStop()
	return g.Serve(lis)
}
