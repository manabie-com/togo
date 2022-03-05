package main

import (
	"context"
	"flag"
	"net"
	"runtime"
	"time"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	"github.com/khangjig/togo/client/logger"
	"github.com/khangjig/togo/client/mysql"
	"github.com/khangjig/togo/client/redis"
	"github.com/khangjig/togo/config"
	serviceGRPC "github.com/khangjig/togo/delivery/grpc"
	serviceHttp "github.com/khangjig/togo/delivery/http"
	"github.com/khangjig/togo/delivery/job"
	"github.com/khangjig/togo/migration"
	"github.com/khangjig/togo/proto"
	"github.com/khangjig/togo/repository"
	"github.com/khangjig/togo/usecase"
)

func main() {
	taskPtr := flag.String("task", "server", "server")

	flag.Parse()

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			logger.GetLogger().Fatal(err.Error())
			runtime.Goexit()
		}

		time.Local = loc
	}

	var (
		client  = mysql.GetClient
		repo    = repository.New(client, redis.GetClient)
		useCase = usecase.New(repo)
	)

	migration.Up(client(context.Background()))

	switch *taskPtr {
	case "server":
		executeServer(useCase, repo)
	default:
		executeServer(useCase, repo)
	}
}

func executeServer(useCase *usecase.UseCase, repo *repository.Repository) {
	cfg := config.GetConfig()

	if len(cfg.HealthCheck.EndPoint) > 0 {
		job.New().Run()
	}

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase, repo)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	// gRPC
	{
		s := grpc.NewServer()

		proto.RegisterTogoServiceServer(s, &serviceGRPC.TogoService{UseCase: useCase})

		go func() {
			errs <- s.Serve(grpcL)
		}()
	}

	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}
}
