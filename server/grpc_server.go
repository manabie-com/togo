package server

import (
	"context"
	"fmt"

	"mini_project/auth_services"

	"mini_project/rpc_services"

	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/gogo/protobuf/protoc-gen-gogo/generator"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func ApplyAuth(ctx context.Context) (context.Context, error) {
	return nil, status.Errorf(codes.Unimplemented, "Please implement auth func for this service")
}

func StartGRPC(wg *sync.WaitGroup, ctx context.Context, dbUrl map[string]string) {

	_, cancel := context.WithCancel(ctx)
	defer cancel()
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	// Register gprc server
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_auth.StreamServerInterceptor(ApplyAuth),
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_auth.UnaryServerInterceptor(ApplyAuth),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(), //should be last in grpc_options
		)),
	)

	// start api services
	apiServer, err := NewAPIServer(dbUrl)
	if err != nil {
		panic(err)
	}

	//start auth services
	authServer, err := auth_services.NewAuthServer(dbUrl)
	defer authServer.Close()
	if err != nil {
		panic(err)
	}

	rpc_services.RegisterServicesServer(grpcServer, apiServer)

	auth_services.RegisterAuthServer(grpcServer, authServer)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("shutting down grpc_server...")
		grpcServer.GracefulStop() // GracefulStop is slow
		cancel()
		wg.Done()
	}()

	// start gRPC grpc_server
	fmt.Println("starting grpc server...")
	grpcServer.Serve(listen)
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origin != "" {
			addCORSHeader(rw, origin)
		} else {
			// allow all origin
			addCORSHeader(rw, "*")
		}
		if req.Method == "OPTIONS" {
			rw.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(rw, req)
	})
}

// mimics from etcd addCORSHeader adds the correct cors headers given an origin
func addCORSHeader(w http.ResponseWriter, origin string) {
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Headers", "accept, content-type, authorization, connection, upgrade, X-Auth-Token")
	w.Header().Add("Access-Control-Expose-Headers", "Grpc-Metadata-X-Subject-Token")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}

func StartHTTP(wg *sync.WaitGroup) {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024*1024))}

	//Register gRPC server endpoint
	mux := http.NewServeMux()
	rmux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(forwardResponseOption),
	)

	err := rpc_services.RegisterServicesHandlerFromEndpoint(ctx, rmux, "localhost:8082", opts)
	if err != nil {
		fmt.Println("Could not connect to grpc_server: ", zap.Error(err))
	}

	err = auth_services.RegisterAuthHandlerFromEndpoint(ctx, rmux, "localhost:8082", opts)
	if err != nil {
		fmt.Println("Could not connect to auth_server: ", zap.Error(err))
	}

	mux.Handle("/api/", wsproxy.WebsocketProxy(rmux))
	mux.HandleFunc("/debug/profile", pprof.Profile)
	mux.HandleFunc("/debug/trace", pprof.Trace)
	mux.HandleFunc("/debug/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/cmdline", pprof.Cmdline)
	mux.Handle("/debug/heap", pprof.Handler("heap"))
	mux.Handle("/debug/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/threadcreate", pprof.Handler("threadcreate"))
	mux.Handle("/debug/block", pprof.Handler("block"))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: allowCORS(mux),
	}

	// handle signal interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Shutting down rest server...")
		srv.Shutdown(ctx)
		cancel()
		wg.Done()
	}()

	fmt.Println("REST server ready...")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("ListenAndServe: ", zap.Error(err))
	}
	return
}

func forwardResponseOption(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	disposKey := "Content-Disposition"
	if vals := md.HeaderMD.Get(disposKey); len(vals) > 0 {
		w.Header().Set(disposKey, vals[len(vals)-1])
		delete(w.Header(), "Grpc-Metadata-"+disposKey)
	}

	lengkey := "Content-Length"
	if vals := md.HeaderMD.Get(lengkey); len(vals) > 0 {
		w.Header().Set(lengkey, vals[len(vals)-1])
		delete(w.Header(), "Grpc-Metadata-"+lengkey)
	}

	return nil
}
