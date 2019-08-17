package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	"github.com/shijuvar/gokit-examples/services/account"
	"github.com/shijuvar/gokit-examples/services/account/cockroachdb"
	accountsvc "github.com/shijuvar/gokit-examples/services/account/implementation"
	"github.com/shijuvar/gokit-examples/services/account/transport"
	grpctransport "github.com/shijuvar/gokit-examples/services/account/transport/grpc"
	"github.com/shijuvar/gokit-examples/services/account/transport/pb"
)

const (
	port     = ":50051"
	dbsource = "postgresql://shijuvar@localhost:26257/ordersdb?sslmode=disable"
)

func main() {
	// initialize our structured logger for the service
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "account",
			"ts", log.DefaultTimestampUTC,
			"clr", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	var db *sql.DB
	{
		var err error
		// Connect to the "ordersdb" database
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	// Create Account Service
	var svc account.Service
	{
		repository, err := cockroachdb.New(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = accountsvc.NewService(repository, logger)
	}

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(svc)
	}

	// set-up grpc transport
	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		accountService  = grpctransport.NewGRPCServer(endpoints, serverOptions, logger)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	var g group.Group
	{
		/*
			Add an actor (function) to the group.
			Each actor must be pre-emptable by an interrupt function.
			That is, if interrupt is invoked, execute should return.
			Also, it must be safe to call interrupt even after execute has returned.
			The first actor (function) to return interrupts all running actors.
			The error is passed to the interrupt functions, and is returned by Run.
		*/
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", port)
			pb.RegisterAccountServer(grpcServer, accountService)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}

	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	/*
		Run all actors (functions) concurrently. When the first actor returns,
		all others are interrupted. Run only returns when all actors have exited.
		Run returns the error returned by the first exiting actor
	*/
	level.Error(logger).Log("exit", g.Run())
}
