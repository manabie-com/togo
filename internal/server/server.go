package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/f6galaxy/kitchen/errors"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

var ll = l.New()

//ServiceServer ...
type ServiceServer interface {
	RegisterWithServer(*grpc.Server)
	RegisterWithHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
	Close(context.Context)
}

type gRPCServer struct {
	*grpc.Server
}

func (s *gRPCServer) WithServiceServer(ss ServiceServer) *gRPCServer {
	ss.RegisterWithServer(s.Server)
	return s
}

func NewGRPCServer() *gRPCServer {
	// grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.PayloadUnaryServerInterceptor(ll, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
				if fullMethodName == "/common.base/Liveness" {
					return false
				}
				if fullMethodName == "/common.base/Readiness" {
					return false
				}
				return true
			}),
		)),
		grpc.MaxRecvMsgSize(10*1024*1024),
	)
	// Register Prometheus metrics handler.
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	return &gRPCServer{Server: s}
}

type gatewayServer struct {
	http.Server
	cfg *configs.Config
}

// NewGatewayServer create new server using given arguments
func NewGatewayServer(cfg *configs.Config) *gatewayServer {
	return &gatewayServer{
		cfg: cfg,
	}
}

// RunGRPCGateway will start an GRPC Gateway
func (s *gatewayServer) RunGRPCGateway(ss ServiceServer) (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	runtime.WithErrorHandler(customHTTPError)
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
				UseEnumNumbers:  true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{},
		}),
	)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10 * 1024 * 1024)),
	}
	conn, err := grpc.Dial(fmt.Sprintf(":%d", s.cfg.GRPCAddress), opts...)
	if err != nil {
		return err
	}

	if err = ss.RegisterWithHandler(ctx, mux, conn); err != nil {
		return err
	}

	muxHttp := http.NewServeMux()
	muxHttp.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})

	muxHttp.Handle("/", mux)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.HTTPAddress), muxHttp)
}

type errorBody struct {
	Err  string `json:"error,omitempty"`
	Msg  string `json:"message,omitempty"`
	Code uint32 `json:"code,omitempty"`
}

func customHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	contentType := "application/json"
	w.Header().Set("Content-type", contentType)
	w.WriteHeader(httpStatusFromCode(errors.Code(grpc.Code(err))))
	jErr := json.NewEncoder(w).Encode(errorBody{
		Err:  grpc.ErrorDesc(err),
		Msg:  grpc.ErrorDesc(err),
		Code: uint32(status.Code(err)),
	})

	if jErr != nil {
		_, _ = w.Write([]byte(fallback))
	}
}

func httpStatusFromCode(code errors.Code) int {
	switch code {
	case errors.NoError:
		return http.StatusOK
	case errors.Canceled:
		return http.StatusRequestTimeout
	case errors.Unknown:
		return http.StatusInternalServerError
	case errors.InvalidArgument:
		return http.StatusBadRequest
	case errors.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case errors.NotFound:
		return http.StatusNotFound
	case errors.AlreadyExists:
		return http.StatusConflict
	case errors.PermissionDenied:
		return http.StatusForbidden
	case errors.Unauthenticated:
		return http.StatusUnauthorized
	case errors.ResourceExhausted:
		return http.StatusTooManyRequests
	case errors.FailedPrecondition:
		return http.StatusPreconditionFailed
	case errors.Aborted:
		return http.StatusConflict
	case errors.OutOfRange:
		return http.StatusBadRequest
	case errors.Unimplemented:
		return http.StatusNotImplemented
	case errors.Internal:
		return http.StatusInternalServerError
	case errors.Unavailable:
		return http.StatusServiceUnavailable
	case errors.DataLoss:
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}

func (s *gatewayServer) WrapCheckMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}

		handler(w, r)
	}
}
