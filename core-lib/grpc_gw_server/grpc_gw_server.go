package grpc_gw_server

import (
	"fmt"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCGWServer struct {
	name string
	port int
	*grpc.Server
	*runtime.ServeMux
}

func New(name string, port int, opts ...grpc.ServerOption) *gRPCGWServer {
	opts = append(opts,
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
	)

	grpcGWServer := &gRPCGWServer{
		name:     name,
		port:     port,
		ServeMux: runtime.NewServeMux(),
		Server:   grpc.NewServer(opts...),
	}

	grpc_prometheus.Register(grpcGWServer.Server)
	grpc_prometheus.EnableHandlingTimeHistogram()

	reflection.Register(grpcGWServer.Server)
	return grpcGWServer
}

func (g *gRPCGWServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return err
	}

	mux := cmux.New(lis)
	grpcLis := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpLis := mux.Match(cmux.HTTP1Fast())

	errGroup := &errgroup.Group{}

	errGroup.Go(func() error {
		return g.Server.Serve(grpcLis)
	})

	errGroup.Go(func() error {
		return http.Serve(httpLis, g.ServeMux)
	})

	errGroup.Go(func() error {
		return mux.Serve()
	})

	defer g.Server.GracefulStop()

	return errGroup.Wait()
}
