package grpc_gw_server

import (
	"context"
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
	name         string
	port         int
	registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	*grpc.Server
	*runtime.ServeMux
}

func New(name string, port int,
	registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error,
	opts ...grpc.ServerOption) *gRPCGWServer {
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
		name:         name,
		port:         port,
		registerFunc: registerFunc,
		ServeMux:     runtime.NewServeMux(),
		Server:       grpc.NewServer(opts...),
	}

	grpc_prometheus.Register(grpcGWServer.Server)
	grpc_prometheus.EnableHandlingTimeHistogram()

	reflection.Register(grpcGWServer.Server)
	return grpcGWServer
}

func (g *gRPCGWServer) Start() error {
	endpoint := fmt.Sprintf(":%d", g.port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return err
	}

	mux := cmux.New(lis)
	grpcLis := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpLis := mux.Match(cmux.HTTP1Fast())

	errGroup := &errgroup.Group{}

	errGroup.Go(func() error {
		return g.Server.Serve(grpcLis)
	})

	errGroup.Go(func() error {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if err := g.registerFunc(ctx, g.ServeMux, endpoint, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
			return err
		}
		return http.Serve(httpLis, g.ServeMux)
	})

	errGroup.Go(func() error {
		return mux.Serve()
	})

	defer g.Server.GracefulStop()

	return errGroup.Wait()
}
